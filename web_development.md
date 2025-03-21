[restful api](/golang.md#http)
[database](/databases.md)
[postgres](/postgresql.md)
[authentication](/auth.md)

# web development

> process of building/maintaining websites, web and mobile applications using programming languages/frameworks/tools

## basic concepts of web application

- static site: content doesn't change dynamically
  - no server-side processing
- dynamic site: pages changes based on data/user

- user interface
- routing: how users navigate between different parts of your application
- DOM (Document Object Model): representation of the HTML elements of a webpage as a tree of nodes and objects
- data fetching: where your data lives and how to get it
- rendering: when and where you render static or dynamic content
- infrastructure: where you deploy, store and run your application code
- scalability: how your application adapts as your team, data and traffic grow
- CORS (Cross-Origin Resource Sharing): security feature enforced by web browsers to prevent malicious websites from making unauthorized requests to another domain
  - a cross-origin request happens when the frontend and backend have different domains, protocols or ports
  - if your frontend and backend are on different origins, your backend must allow cors explicitly
  - a domain/origin consists of: `protocol://domain:port`
    - example of different domains: `http://localhost:3000` and `http://localhost:5000`
      - so if this is the frontend (next.js) of an application making a request to the backend (api), this is considered a cross-origin request
    - example of same-origin request: `https://myapp.com` and `https://myapp.com/api`
      - because the protocol (`https`), domain (`myapp.com`) and port (default 443) are the same
  - by default, browsers block cross-origin requests unless the server explicitly allows them using cors headers
  - when does it matter?
    - when frontend (Next.js, React, etc.) is hosted on a different domain than your backend API
    - when calling an external API from your frontend
    - when using a client-side request (e.g., `fetch()` or `Axios`) to access an API on another origin
  - common cors errors and fixes:
    - Error: "CORS policy: No 'Access-Control-Allow-Origin' header"
      - Fix: Ensure your API includes Access-Control-Allow-Origin in its response headers
    - Error: "CORS policy: Method not allowed"
      - Fix: Make sure your Access-Control-Allow-Methods includes the request method
    - Error: "CORS policy: Request header not allowed"
      - Fix: If sending custom headers (like Authorization), add them to Access-Control-Allow-Headers

- client-side vs server-side: describes where web application code runs
  - client-side: code that runs in the client's computer
  - server-side: code that runs in the server
  - client: hardware or software that requests and consumes resources/services from a server
    - resource: digital asset or data that a client can request from a server
    - initiates communication and requests resources
    - e.g. web browsers, mobile apps
    - responsible for:
      - displaying information for users
      - capturing user input
      - making http requests to the server to fetch or send data
  - server: hardware or software that provides resources/services to clients over a network (usually the internet)
    - waits for requests and provides responses
    - e.g. web servers
    - responsible for:
      - processing client requests
      - reading/writing data to a database or external API
      - sending responses back to the client
    - static web server: web server that serves files that are stored on disk
      - e.g. HTML, CSS, JS files
    - dynamic web server: web server that generates content on the fly
      - e.g. based on user input, user authentication, etc

- service: functionality or resource that a service provides to an client upon request

- endpoint: specific URL where the server provides a resource/service
  - example:
    - endpoint url: `https://api.example.com/books`
    - `app.get('/books', ...)`: defines an endpoint that listens for GET requests at the /books URL
- handler: processes http requests and returns responses

when a user visits a webpage:

1. client (browser) makes a request to the server
1. server returns an HTML file to the browser
1. the browser reads HTML and then constructs the DOM

- DOM (Document Object Model): representation of the HTML elements of a webpage as a tree of nodes and objects
  - each part of the page becomes a node
  - each node can be accessed and manipulated by programming languages like Javascript
- DOM nodes
  - element node: represents html elements (e.g. `<div>`, `<p>`)
  - attribute node: represents attributes of elements (e.g. `class`, `id`)
  - text nodes: represent the text content inside elements
    - `<p>This is a text node.</p>` is an HTML paragraph element where:
      - `<p>` and `</p>` are the element nodes (the HTML tags)
      - `This is a text node.` is the text node (the actual content between the tags)
  - comment nodes: comments in the document
    - e.g.: `<!-- This is a comment -->`

[http](/networking.md#http-(hypertext-transfer-protocol))

- HTML represents the **initial** page content
- DOM represents the **updated** page content

- cookie: small piece of data stored on client's browser by websites to remember stateful information
  - e.g. user preferences, items in a shopping cart, refresh and access tokens, etc
  - initially sent by the server
  - client sends it back in afterwards
  - used in browsers, not in mobile apps
- webhook: event that is sent to the server by an external service when something happens
  - webhook handlers must be idempotent, because the external service will problably retry requests multiple times
    - idempotent: operation that has the same result when repeated
  - 2xx http codes will signal the external service that you processed the request successfully, and they'll stop retrying it
  - http request sent by external service to notify about events
  - one-way communication
- websockets: protocol that enables two-way communication between client and server over a single, long-lived connection

- pull vs push based system

- blob: Binary Large OBject
  - large binary/non-text data like: images, videos and audio

## steps

- add simple CI tasks from the beggining (linting, testing)
- add tests as you progress
- build back end first, then go to front end

1. define features to implement
1. sketch user interface
1. choose stack (front end, back end)
1. choose where project will be hosted (recommended is to start with localhost then change to public cloud)
1. initialize project (e.g. `npm init -y` or `npx create-next-app@latest .`)
1. install dependencies
1. define database schema
1. create database with migration tool (I recommend [goose](https://github.com/pressly/goose))
1. add `.env` with necessary secrets
1. set up and test database connection
1. set up and test api endpoints
1. implement and test user authentication (optional)
1. handle requests for each endpoint
1. set up basic front end
1. integrate front end with back end
1. deploy
1. monitor and maintain

## routing

### URL (Uniform Resource Locator)

> web address that specifies the location of a resource on the internet

- URI contains URL and also URN

- URI (Uniform Resource Identifier): identifies a resource
  - URN (Uniform Resource Name): identifies a resource by name in a particular namespace
  - URL

#### URL structure

- `https://username:passwd@www.example.com:80/blog/?search=test&sort_by=created_at#header`
  - protocol (scheme): `http://`, `https://`, `ftp://`, `mailto:`
    - all URLs have `:`
    - `//` is only included for schemes that have an [authority component](https://www.rfc-editor.org/rfc/rfc3986#section-3.2)
  - username: `username`
  - password: `passwd`
  - subdomain: `www`
  - domain name: `example.com`
  - port: `80`
    - the port component of a URL is often not visible when browsing the internet, because 99% of the time you're using default ports
  - path: `/blog`
  - query parameters: `?search=test&sort_by=created_at`
    - key-value pairs that are separated by `&`
      - `search=test`
      - `sort_by=created_at`
  - fragment/anchor: `#header`

## types of web applications

### Single Page Application (SPA)

- dynamically updates the content on a single page without requiring full page reload
  - only necessary data is fetched from the server, and Javascript updates the view on the client side
- composed by a single html page
- client-side rendering
- asynchronous data fetching
  - allows the app to update without refreshing the entire page
- slow initial load time
- examples
  - gmail, facebook, trello

### Multi-Page Application

- each interaction/request loads a new HTML page from the server
  - each time the user navigates to a new page, the browser requests the html, css, and Javascript for that page
- server-side rendering
  - each page fully reloads when the user navigates
- examples
  - e-commerce sites, news sites

### Progressive Web Application (PWA)

> aims to provide UX similar to a native mobile app in the web browser

- AKA native web app
- can work offline
- can be installed on user's device
- examples
  - starbucks

### Server-side Rendered Application (SSR)

- server generates de HTML for each page request
- examples
  - next.js apps

### Static Site Generator (SSG)

- examples
  - blogs, documentation sites

## debugging a server

- write code
- build and run
- send request to server using a browser or a tool like curl
- see if response is what you expected
- if not, add logging or change the code and repeat

## full-stack web development

> development of both client-side and server-side software components of a web application

- client-side/front end:
  - user interface and user experience
  - technologies: HTML, CSS, JavaScript, frameworks (React, Vue, Angular)
  - makes HTTP requests to the back end
  - handles user interactions and data display

- server-side/back end:
  - business logic, database operations
  - authentication, authorization
  - API endpoints and services
  - technologies: Node.js, Python, Go, databases (SQL, NoSQL)
  - receives HTTP requests from the front end and sends responses

### API integration

> process of connecting different software applications and services to allow them to communicate with each other

- APIs allow services to communicate with other services without knowing how they're implemented
  - service: software component that provides specific functionality
- normally they implement CRUD operations

- create database schema
- program api endpoints
- handle CRUD operations
- find a tool to execute SQL migrations (e.g. goose)
  - add credentials (e.g. connection string)
- convert database results into api responses

#### database migration with goose and api integration with sqlc

- goose: runs database migrations from a set of SQL files using golang
  - reads migration files and applies them to the database
- sqlc: generates go code from SQL files that enables api integration with a database
  - parses sql statements
  - generates type-safe go code that interacts with the database for applications to use

[example of golang api implementation](./code/golang/example5)
[sqlc tutorial](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html#setting-up)
[using sqlc and goose](https://pressly.github.io/goose/blog/2024/goose-sqlc/)
[goose basic operations](https://github.com/pressly/goose)

1. create the following directories from the project's root
  - `./sql/schema`: contains SQL files to build schema
  - `./sql/queries`: contains SQL files to run queries

2. write `.yaml` configuration file for sqlc

sqlc will use this file's configuration to generate go code

[example](./code/golang/example5/sqlc.yaml)

3. `cd` into `./sql/schema` directory and create the sql file that will define the current migration

```bash
goose -s create create_user sql
```

- `-s`: use sequential numbering for new migrations (random numbers if flag is not used)
- `create`: creates new migration file with the current timestamp
- `create_user`: name of the sql query
- this command will create an sql file with a name in the following format `querynumber_name.sql` (e.g. `00001_create_users.sql`)

4. write the desired sql query in the new file:

```sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  email TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
```

- `-- +goose Up`: operation to change state of database
- `-- +goose Down`: operation to revert changes
- `-- +goose StatementBegin`: marks the beggining of a statement
- `-- +goose StatementEnd`: marks the end of a statement

5. run the following command with the [conection string](/postgresql.md#connection-string) to create the tables:

```bash
goose postgres "connection_string" up
```

> [!IMPORTANT]
> run the command on the same directory as the sql query for the migration

to revert changes:

```bash
goose postgres "connection_string" down
```

- use `goose postgres "connection_string" up` to apply the changes to the database
  - `postgres`: name of the database driver (e.g. `mysql`, `turso`)

6. write sql query file in `sql/queries` (e.g. `users.sql`)

```sql
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email)
VALUES (
    ...
)
RETURNING *;
```

- query annotations: sqlc requires each query to have a small comment indicating the name (`-- name:`) and command (`:one` or `:many`)
  - `:one`: the query will return 1 record
    - specifies how many users the query will return
  - `:many`: the query will return multiple records

7. `cd` into the root and run the command below to generate go code that interacts with the database:

```bash
sqlc generate
```

8. use query in the endpoint handler

> [!IMPORTANT]
> this code uses the `uuid` package, so add it to your module
> go to definition of `CreateUser()` to see the function signature if you don't know what is the input
> `Context()` is used to handle timeouts, so it is always part of the input
> always convert the obtained resource from `DbResource{}` to the `Resource{}` struct
> if you change the query, you need to run `sqlc generate` again at the root of the project
> endpoint with id on url path: `mux.HandleFunc("GET /api/users/{userID}", apiCfg.handleUserGet)`
> get the id from the path: `http.Request.PathValue("userID")`
> after getting id from url path, convert it to a `uuid.UUID` with `userID, err := uuid.Parse(strID)`

```go
newUser, err := cfg.databaseQueries.CreateUser(r.Context(), post.Email)
if err != nil {
  log.Printf("error creating user: %v", err)
  w.WriteHeader(http.StatusInternalServerError)
  return
}
```

- the .env file stores credentials to access the database and connection strigns
- why do i need to turn off ssl?
- how do i use the context package?
  - you dont need to use it directly, sqlc will generate code that uses it
- how do i tell golang to run the sql query with the data sent by the user?

#### integrate next.js with postgresql

1. install `pg`

```bash
npm install pg
```

2. set environment variables in the root of your project (`.env.developemnt`)

```
DB_USER=your_user
DB_HOST=your_host
DB_NAME=your_database
DB_PASSWORD=your_password
DB_PORT=5432
```

3. configure database connection in `src/lib/db.ts`

```typescript
import { Pool } from 'pg';

dotenv.config();

const pool = new Pool({
  user: process.env.DB_USER,
  host: process.env.DB_HOST,
  database: process.env.DB_NAME,
  password: process.env.DB_PASSWORD,
  port: Number(process.env.DB_PORT),
});

console.log(`Server will run on port: ${process.env.DB_PORT}`);

export const query = (text: string, params?: any[]) => pool.query(text, params);

// Function to test the database connection
export async function testDBConnection() {
  try {
    const res = await pool.query('SELECT 1+1 AS result');
    console.log('✅ Database connected successfully!', res.rows);
  } catch (error) {
    console.error('❌ Database connection failed:', error);
  }
}
```

to test if database connection is working, add this code in `src/app/page.tsx`:

```typescript
import { testDBConnection } from '@/lib/db';

testDBConnection();
```

4. create api route in `src/app/api/users/route.ts`

consult the users table on postgresql before programming the sql query

[example](https://github.com/troclaux/peso/blob/main/src/app/api/users/route.ts)

5. create database table with goose as a migration tool

6. test the api with curl

#### ORM (Object-Relational Mapping)

> programming technique that allows developers to interact with a database using an object-oriented paradigm

- prisma (recommended)
  - typescript
  - supports postgresql, mysql, mongodb, sqlite
  - simpler
- TypeORM
  - typescript
- sequelize
  - javascript
  - supports postgresql, mysql, sqlite, mongodb

## types of web applications

#### Single-Page Applications (SPAs)

- loads a single HTML page and dynamically updates content
- smooth user experience with no page refreshes
- examples: Gmail, Facebook

#### Multi-Page Applications (MPAs)

- traditional approach where each page is a separate HTML document
- full page reload on navigation
- better for content-heavy sites
- examples: Wikipedia, news sites

#### Progressive Web Apps (PWAs)

- web apps that work like native mobile apps
- can work offline
- installable on devices
- examples: Twitter, Pinterest

## Client-Side Rendering vs Server-Side Rendering

### Client-Side Rendering (CSR)

- client-side rendering is the most common way to render a web page
- the browser downloads a minimal HTML page, then downloads and executes JavaScript that renders the page
- the JavaScript can make requests to the server to fetch data and update the page

### Server-Side Rendering (SSR)

- server generates the HTML for the page and sends it to the browser
- the browser receives a fully rendered page, which can be faster to load and better for SEO
- server-side rendering can be more complex to set up and maintain

- server generates complete HTML pages
- faster initial page load
- better SEO performance
- examples: Next.js applications

## deploy

> process of making web application available and accessible to user over the internet

- requires transfering the application's code, assets and configurations from development environment to a production server

- steps in deployment
  - development
  - building
  - testing
  - provisioning: setting up servers/databases/infrastructure
  - deployment
  - configuration
  - monitoring and maintenance

### step-by-step

summary:

1. Register domain that points to the server's public IP
1. Implement dockerfile for each app service (e.g. next.js, nginx)
1. Implement docker-compose or kubernetes yaml config file to setup the app
1. Authenticate to be able to pull from container image registry (docker hub or ECR)
1. Start Nginx in HTTP mode (without SSL)
1. Run Certbot to generate SSL certificates
1. Verify certificates exist (ls ~/peso/certbot/conf/live/pesodevops.com/)
1. Modify nginx.conf to enable HTTPS
1. Restart Nginx (docker compose down && docker compose up -d)
1. Verify HTTPS works (curl -I https://pesodevops.com)
1. Automate renewal using a cron job

1. Register domain that points to the server's public IP

- make sure the hosted zone's records are correct
  - `A` type record points to the server's public IP
  - `NS` type record points to the registered domain's name servers
    - defines which records (e.g. `ns-1234.awsdns.net`) are authoritative for a registered domain (e.g. `www.example.com`)
      - authoritative: it means the server holds the official, up-to-date records for that domain and is responsible for answering DNS queries about it

1. Implement dockerfile for each app service (e.g. next.js, nginx)

[example](./code/dockerfiles/nginx.Dockerfile)

1. Implement docker-compose or kubernetes yaml config file to setup the app

1. Authenticate to be able to pull from container image registry (docker hub or ECR)

`aws ecr get-login-password --region sa-east-1 | docker login --username AWS --password-stdin [AWS_ACCOUNT_ID].dkr.ecr.sa-east-1.amazonaws.com`

1. set up domain name with a domain registar (e.g. route 53, cloudfare)

- go to domain provider (e.g. router 53, cloudfare, namecheap)
- create an `A` record
  - host: `@`
  - value: the server public IP
- wait for dns to propagate


1. set up app with containers

recommended directory structure in server:

```
- app_name
  - certbot
    - www
    - conf
```

1. start nginx in http-only mode (without ssl certificates and https)

- nginx must be running on port 80 without ssl
  - this allows Let's Encrypt to perform domain validation
  - use this [template](./code/nginx/nginx1.conf)

- why?
  - Let's Encrypt validates your domain by checking if it can access files served at:
    - `http://pesodevops.com/.well-known/acme-challenge/`
    - this means port 80 must be open and functional

1. start docker compose services

- start containers: `docker compose up -d`
- verify nginx is running and serving http requests: `curl -I http://pesodevops.com`

expected response:

```bash
HTTP/1.1 200 OK
Server: nginx
...
```

1. generate SSL certificates using certbot

once nginx is running in http mode, manually run certbot to request ssl certificates

```bash
docker compose run --rm certbot certonly \
  --webroot -w /var/www/certbot \
  -d mydomain.com -d www.mydomain.com \
  --email my_email@hotmail.com \
  --agree-tos \
  --no-eff-email
```

- what happens here?
  - certbot places a temporary challenge file inside `/var/www/certbot/.well-known/acme-challenge/`
  - Let's Encrypt tries to access `http://pesodevops.com/.well-known/acme-challenge/{token}`
  - if the request succeeds, Let's Encrypt issues the SSL certificate

1. confirm certificate generation

```bash
ls -l ~/peso/certbot/conf/live/mydomain.com/
```

- expected files: `cert.pem  chain.pem  fullchain.pem  privkey.pem`

1. update nginx.conf to enable SSL

[updated nginx.conf](./code/nginx/nginx2.conf)

1. restart nginx with ssl enabled

```bash
docker compose down
docker compose up -d
```

check that nginx is running properly:

```bash
docker compose logs nginx
```

1. verify https is working

```bash
curl -I https://pesodevops.com
```

expected response:

```
HTTP/2 200
server: nginx
...
```

insert URL in browser to see if it's working

1. automate SSL certificate renewal

ssh into server and run:

```bash
crontab -e
```

- add this cron job: `0 3 * * * cd ~/peso && docker compose run --rm certbot renew && docker compose restart nginx`

- test automatic renewal: `sudo certbot renew --dry-run`
  - let’s encrypt certificates expire every 90 days

### questions

- what is the purpuse of `certbot/www` and `certbot/conf` directory?
  - `certbot/www`: temporary challenge files for domain verification
    - what are challenge files?
      - temporary verification files that certbot creates to prove that you own a domain before issuing an SSL certificate
    - when certbot needs to prove domain ownership, it places challenge files inside this directory
  - `certbot/conf`: stores certbot’s configuration
    - stores SSL certificates and renewal settings
    - after a certificate is issued, the private key, full chain and other related files are stored here
    - when the certificate needs renewal, certbot updates this directory
- do i need to add a cron job to renew certificate? yes
- how frequently should i renew certificate?
  - Let's Encrypt certificates are valid for 90 days
  - it's recommended to renew them each 60 days
- whats the limit for certificate renewal?
  - `certbot renew` only renews if the certificate expires in ≤ 30 days, so don't worry
