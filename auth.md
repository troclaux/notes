
# Best resources:

- Oauth
  - [https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth](https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth)
- JWT
  - [https://jwt.io/introduction](https://jwt.io/introduction)

# OAuth

Decouples authentication from authorization

## Step-by-step

1. App requests authorization from User
2. User authorizes App and delivers proof
3. App presents proof of authorization to server to get a Token
4. Token is restricted to only access what the User authorized for the specific App

## OAuth components

- scopes and consent
  - Scopes are bundles of permissions asked by the client when requesting a token
  - Consent is the permission that the user gives to the application to access data
- actors
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
- tokens
  - Usually are Json Web Tokens (JWT)
  - Access token
    - Short lived (can last minutes or hours)
    - Can scale out
    - Can't be revoked
  - Refresh token
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

Handles authorization and information exchange

## A JWT looks like

`xxxxx.yyyyy.zzzzz`

## Structure

- Header
  - Example of a header:
    ```json
    {
      "alg": "HS256",
      "typ": "JWT"
    }
    ```
- Payload
  - Example of a payload:
    ```json
    {
      "sub": "1234567890",
      "name": "John Doe",
      "admin": true
    }
    ```
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
- Signature
  - To create a signature:
    - Encode payload + a secret + the algorithm specified in the header
  - The signature is used to check if the message wasn't changed along the way
  - In the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is
  - If you send JWT through HTTP header try to prevent them from getting too big. Some servers don't accept more than 8 KB in headers

  - Example:
    ```javascript
    const crypto = require('crypto');

    const header = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'; // Base64Url encoded header
    const payload = 'eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIn0'; // Base64Url encoded payload
    const secret = 'your-secret-key';
    const algorithm = 'HS256';

    const signature = crypto.createHmac('sha256', secret)
      .update(header + '.' + payload)
      .digest('base64');

    const jwt = header + '.' + payload + '.' + signature;
    console.log(jwt);
    ```

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

