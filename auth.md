
# auth

## authentication

> process of verifying the identity of a user

> who are you?

- types of authentication:
  - id + password
    - don't allow weak passwords
    - don't store passwords as plain text, hash them before storing in database
      - always use bcrypt to hash passwords that will be stored in the database
      - [avoid MD5 and SHA-3](https://blog.boot.dev/cryptography/how-sha-2-works-step-by-step-sha-256/#can-i-use-sha-256-to-hash-passwords)
    - id can be email, username, etc
  - OAuth
    - e.g. sign in with google, sign in with github, etc
  - magic links
    - e.g. `https://example.com/login?token=...`
  - passkeys

- token: string of characters that represents a user's identity, permissions or access rights
- access token: piece of data that represents the authorization granted to a client to access specific resources and identifies the user's permissions and identity
  - used to secure access to APIs and other resources
  - isn't stored in the database
- refresh token: allows a client to obtain a new access token when an existing access token expires
  - doesn't provide access to resources directly
  - can be revoked
  - generally last longer than access token
  - stored in the database (back end) and in http-only cookies (front end)
  - why not increase access tokens expiration time?
    - access tokens are exposed while in transit, while refresh tokens are stored securely and only used to request new access tokens
    - access tokens are bearer tokens, where "possession equals authority"
    - access tokens can be stolen in cross-site scripting (XSS) attacks
  - bearer token: a type of access token that is sent with the request to access protected resources
    - included in the Authorization header of the request, prefixed by the word "Bearer"
  - Sender-Constrained Tokens (SCT): tokens that are bound cryptographic to the client through a process that uses certificates or private/public keys
    - can use JWTs

- cookie: small piece of data stored on client's browser by websites to remember stateful information
  - e.g. user preferences, items in a shopping cart, refresh and access tokens, etc
  - initially sent by the server
  - client sends it back in afterwards
  - used in browsers, not in mobile apps

> [!IMPORTANT]
> don't store secrets/JWTs/credentials in the browser's local storage
> because local storage is accessible to any javascript running on the page (including browser extensions)
> instead, store access tokens in variables and refresh tokens in http-only cookies (not accessible by javascript)
> if you store the access token in a variable, it will be lost when the component is refreshed

```jsx
const [accessToken, setAccessToken] = useState(null);
```

- benefits of storing access token with `useState`:
  - persists after re-renders
  - triggers ui updates
  - works with `useEffect` to fetch data

## authorization

> process of determining what actions a user can perform within a system

> what can the user do?

- most common auth operations:
  - login component
  - signup component
  - logout
  - password recovery/reset
  - get tokens
  - refresh tokens
  - get user data
  - update user data
  - block page access if unauthenticated

## JWT (JSON Web Tokens)

> stateless authentication mechanism that is widely used in web development

[JWT](https://jwt.io/introduction)

JWTs lifecycle:

1. user submits login + password
1. server's handler queries the database to find the user
1. compare user's hashed password in the database with password in the request
1. the server generated an access token and a refresh token
1. the server responds with the access token while storing the refresh token in an http-only cookie
1. client stores the access token in state or context
1. when access token expires, client sends refresh token to server to get a new access token (`api/refresh`)
1. client sends the access token in the Authorization header in all future requests to the server
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

- between multiple login requests with jwt using the same user and password:
  - the header doesn't change
  - the payload changes even if it's the same user and password, because there's also `iat` (issued at) and `exp` (expiration time)
  - the signature changes because it's generated based on the header and payload

login implementation summary:

> [!TIP]
> login is basically a POST request to the server with the user's credentials that receives a JWT in return

- what should the application's login api route do when a user logs in?
  - validate credentials
  - generate an access token (short-lived, e.g., 15 minutes)
  - generate a refresh token (long-lived, e.g., 7 days)
  - store the refresh token in the database
  - send both tokens to the frontend

- when the access token expires:
  - the front end sends the refresh token to get a new access token
  - if the refresh token is valid:
    - issue a new access token
    - optionally issue a new refresh token (rotate refresh tokens)
  - if the refresh token is invalid or expired:
    - force the user to log in again

implementation steps:

- install packages: `npm install jsonwebtoken bcrypt pg`
- create refresh token table in the database (back end)
- create `JWT_SECRET` in `.env`
  - generate a random string to be the value of `JWT_SECRET`  with the command `openssl rand -base64 64 | tr -d '\n'`
- test access and refresh token generation (back end)
- implement login api (back end)
- create table for refresh tokens in database (back end)
- implement endpoint to refresh refresh tokens (back end)
- implement login page (front end)
  - login page generates access and refresh tokens
  - store access token with `useState`
  - store refresh token in http-only cookie

example of a refresh token table:

```
-- +goose Up
-- +goose StatementBegin
CREATE TABLE refresh_tokens (
  token TEXT PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  expires_at TIMESTAMP NOT NULL,
  revoked_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;
-- +goose StatementEnd
```

#### Header

Consists of two parts:

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
  - registered claims (not mandatory but recommended)
    - iss (issuer)
    - exp (expiration time)
    - sub (subject)
    - aud (audience)
    - iat (issued at)
  - public claims
    - can be defined at will by those using JWTs
    - to avoid collisions, they should be defined in the IANA JSON Web Token Registry or be defined as a URI that contains a collision-resistant namespace
  - private claims
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

### good practices

- use safe algorithms
  - HS256, RS256
- keep the secret key safe
  - don't expose it in the client side
  - don't store it in the database
- keep short-lived tokens
  - access tokens: minutes
  - refresh tokens: days, weeks, months
- use HTTPS
- allow token revocation
- validate tokens on the server, checking signature and claims
- don't add too many claims
  - keep the payload small to minimize token size
- verify the `iss` claim, to check if the token was issued by a trusted party
- use default claims
  - `exp` (expiration time)
  - `iat` (issued at)
  - `nbf` (not before)
  - `sub` (subject)
  - `aud` (audience)

## OAuth

> protocol that allows users to grant third-party application limited access ot their resource on another service provider's website (e.g. Google, Github)

- widely used authorization framework
- enables applications to secure designated access to user accounts on other applications and server
- decouples authentication from authorization

[Oauth](https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth)

- OAuth 2.0 vs OpenID Connect vs SSO vs SAML:
  - OAuth: authorization framework
  - OpenID Connect (OIDC): adds authentication layer on top of OAuth 2.0
    - authorization + authentication
    - allow users to authenticate using a third-party identity provider (e.g. google, facebook)
  - SSO (Single Sign-On): allows users to authenticate once and access multiple applications
    - e.g. sign in with google and access multiple apps
  - SAML (Security Assertion Markup Language): XML-based open standard for exchanging authentication and authorization data between parties
    - enables SSO implementation
    - supports OIDC and OAuth 2.0

### step-by-step

1. user inserts credentials
1. User authorizes App and delivers proof
1. app sends proof of authorization to server, that sends a token in response
1. whenever the user wants to access a protected resource, the user-agent should sent the JWT
  - the JWT typically is sent in the Authorization header using the Bearer schema

- security best practices:
  - don't keep tokens longer than necessary (try to minimize the time that the token is valid)
  - don't store sensitive session data in browser storage (local storage, session storage, etc)
    - because it's accessible by any javascript running on the page (including browser extensions)

### OAuth 2.0 Authentication Flow (Google Sign-In)

1. user clicks "Sign in with Google"
   - frontend initiates the OAuth login flow
2. browser redirects to Google's authorization URL
   - this includes query parameters like `client_id`, `redirect_uri`, `response_type=code`, and `scope`
   - traffic: AWS VPC (frontend) => Google (all egress traffic is allowed by default)
3. user sees Google's consent screen and approves access
   - if the user agrees, Google redirects them back to your app's `redirect_uri` with an authorization code in the URL
4. backend server exchanges the authorization code for an access token
   - backend sends a `POST` request to `https://oauth2.googleapis.com/token` with:
     - `client_id`, `client_secret`, `code`, `redirect_uri`, and `grant_type=authorization_code`
5. Google responds with an access token
   - The response includes:
     - `access_token` (to access Google APIs)
     - `id_token` (contains user profile in JWT format)
     - `refresh_token` (if applicable)
     - `expires_in`
6. backend retrieves user info from Google
   - makes `GET` request to `https://www.googleapis.com/oauth2/v2/userinfo` using the `access_token`
   - Google returns user details (name, email, profile picture, etc)

- configure OAuth 2.0 client IDs [here](https://console.cloud.google.com/auth/clients)

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

- [better auth (recommended)](https://www.better-auth.com/)
  - keep user data in your database
  - OAuth connections
  - good documentation
  - quick and easy to set up
  - free
- [clerk](https://clerk.com/)
  - user data is kept in clerk's database instead of yours, but you can still access it
  - OAuth connections
  - good documentation
  - quick and easy to set up
  - pay only after a certain number of users
- [auth.js](https://authjs.dev/)
  - keep user data in your database
  - harder to set up
- [auth0](https://next-auth.js.org/)
- [okta](https://www.okta.com/)

### auth.js

1. install auth.js package
1. follow auth.js [auth.js](https://authjs.dev/) documentation to setup login and signup components
1. setup database
1. setup .env.locale with the required environment variables listed below

> [!WARNING]
> if you get the error `edge runtime does not support Node.js 'crypto' module`
> delete the `middleware.ts` folder and restart the server

- use `.env.local` to store your environment variables:
  - `GOOSE_DBSTRING`
  - `GOOSE_DRIVER`
  - `AUTH_GOOGLE_ID`
  - `AUTH_GOOGLE_SECRET`
  - `AUTH_SECRET`
  - `DATABASE_HOST`
  - `DATABASE_NAME`
  - `DATABASE_USER`
  - `DATABASE_PASSWORD`

### clerk

- clerk hooks
- clerk helpers
  - client-side
    - `useUser()`
    - `useClerk()`
    - `useAuth()`
    - `useSignIn()`
    - `useSignUp()`
    - `useSession()`
    - `useSessionList()`
    - `useOrganization()`
    - `useOrganizationList()`
  - server-side
    - `auth()`: returns `Auth` object that contains:
      - `sessionId`
      - `userId`
      - `orgId`
      - `sessionClaims`
      - `getToken()`
      - `has()`: checks if the user has authorization to access a component
    - `currentUser`: holds info about a user
    - route handlers
    - server actions

### better auth

#### add Google SSO OAuth to your Next.js app

To obtain the values for `AUTH_GOOGLE_ID` (your Google Client ID) and `AUTH_GOOGLE_SECRET` (your Google Client Secret), you need to create an OAuth 2.0 credential in the Google Developer Console. Here’s how you can generate them:

1. Go to Google Cloud Console:

- Navigate to the [Google Cloud Console](https://console.cloud.google.com/)
- If you don’t have a project, create one by clicking on the `Select a project` dropdown at the top, then `New Project`

2. Enable the Google API:

- Once your project is created, you need to enable the Google API that you want to use (for example, Google OAuth, Google Sign-In, or another API).
- Go to `APIs & Services > Library` in the Google Cloud Console.
- Search for the API you want to enable (e.g., `Google+ API` or `Google Identity Services` for login)
- Click on the API and enable it by clicking the `Enable` button

3. Create OAuth 2.0 Credentials:

- In the Google Cloud Console, go to `APIs & Services > Credentials`
- Click on the `Create Credentials` button and select OAuth 2.0 Client IDs

4. configure the OAuth consent screen:

- if you haven’t set up the OAuth consent screen yet, you’ll be prompted to configure it
- this screen is what users will see when they are asked to grant your app permission to access their data
- fill in the required fields (app name, email, etc.)
- leave the other options as default if you're just setting up for development

5. Generate OAuth Client ID:

- under `application type`, select `web application` (if you’re building a web app)
- In the Authorized JavaScript origins field, add the URLs that will be using this OAuth credential (e.g., `http://localhost:3000` for local development).
- In the Authorized redirect URIs field, add the URL that users will be redirected to after logging in (e.g., `http://localhost:3000/api/auth/callback/google` for Next.js).
- Click Create.

6. Get the Client ID and Client Secret:

- Once your OAuth 2.0 credentials are created, you’ll see a popup containing the **Client ID** and **Client Secret**.
- You can also find them later in `APIs & Services > Credentials`

7. Copy the Values:

- Copy the Client ID and Client Secret and use them in your application. These values are the ones you'll put in your environment variables (e.g., `AUTH_GOOGLE_ID={CLIENT_ID}` and `AUTH_GOOGLE_SECRET={CLIENT_SECRET}`).

Example:

```env
AUTH_GOOGLE_ID=your-client-id.apps.googleusercontent.com
AUTH_GOOGLE_SECRET=your-client-secret
```

Additional Notes:

- Client ID: It uniquely identifies your app to Google's OAuth servers.
- Client Secret: A secret key used to authenticate your app when requesting tokens or making API requests on behalf of users.

- Now, you can use these credentials in your app to integrate Google authentication
- If you're using an authentication library like `next-auth` or another OAuth client, you can add these values to your environment variables and pass them to the OAuth configuration.

---

- claims: stores statements about an entity (typically, the user) and additional data
- session ID: stateful random string that acts as pointer to the actual user record in the database
