
# auth

## authentication

> process of verifying the identity of a user

- types of authentication:
  - id + password
    - don't allow weak passwords
    - don't store passwords as plain text, hash them before storing in database
      - always use bcrypt to hash passwords that will be stored in the database
      - [avoid MD5 and SHA-3](https://blog.boot.dev/cryptography/how-sha-2-works-step-by-step-sha-256/#can-i-use-sha-256-to-hash-passwords)
    - id can be email, username, etc
  - 3rd party authentication
    - e.g. sign in with google, sign in with github, etc
  - magic links
    - e.g. `https://example.com/login?token=...`
  - API keys

- token: string of characters that represents a user's identity, permissions or access rights
- access token: piece of data that represents the authorization granted to a client to access specific resources and identifies the user's permissions and identity
  - used to secure access to APIs and other resources
  - isn't stored in the database
- refresh token: allows a client to obtain a new access token when an existing access token expires
  - doesn't provide access to resources directly
  - can be revoked
  - generally last longer than access token
  - stored in the database
  - why not increase access tokens expiration time?
    - access tokens are exposed while in transit, while refresh tokens are stored securely and only used to request new access tokens
    - access tokens are bearer tokens, where "possession equals authority"
    - access tokens can be stolen in cross-site scripting (XSS) attacks

- cookie: small piece of data stored on client's browser by websites to remember stateful information
  - e.g. user preferences, items in a shopping cart, refresh and access tokens, etc
  - initially sent by the server
  - client sends it back in afterwards
  - used in browsers, not in mobile apps

> [!IMPORTANT]
> don't store secrets/JWTs/credentials in the browser's local storage
> because local storage is accessible to any javascript running on the page

## authorization

> process of determining what actions a user can perform within a system

## JWT (JSON Web Tokens)

> stateless authentication mechanism that is widely used in web development

[JWT](https://jwt.io/introduction)

JWTs lifecycle:

1. user submits login + password
1. handler makes sql query on database and tries to get the user row
1. compare user's hashed password in the database with password in the request
1. create JWT with user ID and sent it to client
1. client sends JWT in all future login requests
1. on every authenticated request, server validades JWT
1. eventually JWT expires and the process repeats

- JWTs are:
  - base64 encoded: the payload (data) is easily readable if decoded with a Base64Url decoder
  - **NOT** encrypted
  - tamper-proof because the signature will be invalid if the header or payload is modified after token creation
    - since the signature is cryptographically bound to the original content using the server's secret key
  - stateless: the server doesn't need to keep track of which users are logged in via JWT
  - irrevocable: can't be revoked once generated
  - expirable: can be set to automatically expire after a specific time period

### JWT Structure

JSON Web Tokens consist of three parts separated by dots in the following order:

1. header
1. payload
1. signature

example of a JWT (the token is the complete string):

- format: `<base64url(header)>.<base64url(payload)>.<base64url(signature)>`
- real example: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`

- base64: enables the convertion of data (images, text, audio) into an ASCII string that can be converted back
  - it's a reversible process, so you can obtain the original data from the base64 string (unlike hashing)
  - use case: when you need to send that over systems that can only handle text
  - a common variante is "base64 URL safe", which omits characters that might cause problems in URL path segments

> [!IMPORTANT]
> once the token is created by the server
> the data in the token cannot be tampered with because the signature is cryptographically bound to the content
> If anyone modifies the header or payload, the signature verification will fail since it was generated using the server's secret key
> This ensures data integrity and authenticity of the token

#### Header

- Consists of two parts:
  1. The type of the token, which is JWT
  2. The signing algorithm being used (e.g. SHA256, RSA, etc)

Example of a header:

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

#### Payload

> stores claims

- claims: statements about an entity (typically, the user) and additional data
- three types of claims:
  1. registered claims (not mandatory but recommended)
    - iss (issuer)
    - exp (expiration time)
    - sub (subject)
    - aud (audience)
    - iat (issued at)
  2. public claims
    - can be defined at will by those using JWTs
    - to avoid collisions, they should be defined in the IANA JSON Web Token Registry or be defined as a URI that contains a collision-resistant namespace
  3. private claims
    - custom claims created to share information between parties that agree on using them

example of a payload:

```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "admin": true
}
```

- `sub` (subject of the token): Identifies the user
- `name`: Name of the user
- `iat` (issued at): Timestamp when the token was generated

#### Signature

> used to verify that the header and the payload weren't modified

- obtained by hashing the `header` + the `payload` with a signing algorithm (e.g. HS256, RSA) and a secret key
  - secret key is used to:
    - generate the signature of a jwt
    - verify the signature of a jwt that was received
- if the token is signed with a private key, it can also verify that the sender of the JWT is who it says it is
- if you send JWT through HTTP header try to prevent them from getting too big
  - Some servers don't accept more than 8 KB in headers

## OAuth

> protocol that allows users to grant third-party application limited access ot their resource on another service provider's website (e.g. Google, Github)

- Widely used authorization framework
- Enables applications to secure designated access to user accounts on other applications and server
- Decouples authentication from authorization

[Oauth](https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth)

step-by-step:

1. User inserts credentials
2. Successful login
3. JWT will be returned
4. whenever the user wants to access a protected resource, the user-agent should sent the JWT
  - the JWT typically is sent in the Authorization header using the Bearer schema

- Security best practices:
  - Dont keep tokens longer than necessary
  - Don't store sensitive session data in browser storage

### Step-by-step

1. App requests authorization from User
2. User authorizes App and delivers proof
3. App presents proof of authorization to server to get a Token
4. Token is restricted to only access what the User authorized for the specific App

### OAuth components

![OAuth protocol flow](./images/oauth_flow.png)

1. Resource Owner
  - The person who is giving access to some portion of their account
  - Example: User, John Doe
2. Resource Server
  - The server that holds the data that the User wants to access
  - Example: Google Fit API
3. Authorization Server
  - The server that authenticates the User and issues the Access Token
  - Example: Google's OAuth 2.0 Authorization Server
4. Client
  - The application that wants to access the User's data
  - Example: Fitness Tracker App

#### scopes and consent

- Scopes are bundles of permissions asked by the client when requesting a token
- Consent is the permission that the user gives to the application to access data

#### actors

- Resource Owner
  - Owner of the data that the client wants to access
  - Example: Bob's cloud storage service is a resource server. It stores Bob's files securely. When a client requests access to specific files the resource server checks the validity of the access token before providing the requested files
- Resource Server
  - API that stores the information that the application wants to access
- Client
  - Application that wants to access your data
  - Public clients (client identification):
    - Browsers
    - Mobile apps
    - IoT devices
  - Confidential clients (client authentication):
    - Also IoT devices (e.g., smart TVs)
    - APIs
- Authorization Server
  - Is responsible for authenticating the resource owner
  - Issues access tokens to the client after the resource owner grants authorization

#### tokens

- Usually are JSON Web Tokens (JWT)
- Access token
  - Short lived (can last minutes or hours)
  - Can scale out
  - Can't be revoked
- Refresh token
  - refreshes expired access tokens
  - Long lived (can last days, months, or years)
  - Can be revoked

#### authorize endpoint

- To get consent and authorization from the user
- Returns an authorization grant that says the user has consented to it
- Then authorization is passed to the token endpoint

#### token endpoint

- Processes the grant
- Gives the refresh token and access token

---

## Auth Providers

What is the best auth provider?
- choices:
  - Supabase (most recommended for beginners)
    - built on top of PostgreSQL
    - open-source
    - easy to use
    - bad documentation
  - Firebase auth
    - good documentation
    - easy to use
  - Auth0
    - expensive
    - good documentation

---

- claims: stores statements about an entity (typically, the user) and additional data
- session ID: stateful random string that acts as pointer to the actual user record in the database
