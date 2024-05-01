
# Best resources:

- Oauth
  - [https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth](https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth)
- JWT
  - [https://jwt.io/introduction](https://jwt.io/introduction)

# OAuth

- Widely used authorization framework
- Enables applications to secure designated access to user accounts on other applications and server
- Decouples authentication from authorization

## Step-by-step

1. App requests authorization from User
2. User authorizes App and delivers proof
3. App presents proof of authorization to server to get a Token
4. Token is restricted to only access what the User authorized for the specific App

## OAuth components

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

### scopes and consent

- Scopes are bundles of permissions asked by the client when requesting a token
- Consent is the permission that the user gives to the application to access data

### actors

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

### tokens

- Usually are JSON Web Tokens (JWT)
- Access token
  - Short lived (can last minutes or hours)
  - Can scale out
  - Can't be revoked
- Refresh token
  - refreshes expired access tokens
  - Long lived (can last days, months, or years)
  - Can be revoked

### authorize endpoint

- To get consent and authorization from the user
- Returns an authorization grant that says the user has consented to it
- Then authorization is passed to the token endpoint

### token endpoint

- Processes the grant
- Gives the refresh token and access token

# JWT (JSON Web Tokens)

> Handles authorization and information exchange

## JWT Structure

In its compact form, JSON Web Tokens consist of three parts separated by dots (.), which are:

- Header
- Payload
- Signature

Therefore, a JWT typically looks like the following

![jwt](./images/encoded-jwt3.png)

The output is three Base64-URL strings separated by dots that can be easily passed in HTML and HTTP environments

Let's break down the different parts

### Header

- Consists of two parts:
  1. The type of the token, which is JWT
  2. The signing algorithm being used (e.g. SHA256, RSA, etc)

- Example of a header:
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```
### Payload

- Contains the claims
  - Claims are statements about an entity (typically, the user) and additional data

- There are three types of claims:
  1. registered (not mandatory but recommended)
    - iss (issuer)
    - exp (expiration time)
    - sub (subject)
    - aud (audience)
  2. public
    - can be defined at will
    - to avoid collisions, they should be defined in the IANA JSON Web Token Registry or be defined as a URI that contains a collision-resistant namespace
  3. private
    - custom claims created to share information between parties that agree on using them

- Example of a payload:
  ```json
  {
    "sub": "1234567890",
    "name": "John Doe",
    "admin": true
  }
  ```

- sub (subject of the token): Identifies the user
- name: Name of the user
- iat (issued at): Timestamp when the token was generated

  - Contains claims
    - claims are statements about an entity (typically, the user) and additional data
  - 3 types of claims:
    - Registered claims
      - iss (issuer)
      - exp (expiration time)
      - sub (subject)
      - aud (audience)
    - Public claims
      - Can be defined at will by those using JWTs
    - Private claims
      - Custom claims created to share information between parties that agree on using them

### Signature

- Checks if the message wasn't changed along the way
- To create a signature:
  - Encode payload + a secret + the algorithm specified in the header
- If the token is signed with a private key, it can also verify that the sender of the JWT is who it says it is
- If you send JWT through HTTP header try to prevent them from getting too big
  - Some servers don't accept more than 8 KB in headers

- OAuth steps:
  1. User inserts credentials
  2. Successful login
  3. JWT will be returned
    - Security best practices:
      - Dont keep tokens longer than necessary
      - Don't store sensitive session data in browser storage
  4. whenever the user wants to access a protected resource, the user-agent should sent the JWT
    - the JWT typically is sent in the Authorization header using the Bearer schema

---

- session ID
  - stateful
  - random string that acts as pointer to the actual user record in the database
- JWT
  - stateless
