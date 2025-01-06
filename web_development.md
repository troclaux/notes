
# web development

> process of building/maintaining websites, web and mobile applications using programming languages/frameworks/tools

[restful api](/golang.md#http)
[database](/databases.md)
[postgres](/postgresql.md)
[authentication](/auth.md)

## basic concepts of web application

- user interface
- routing: how users navigate between different parts of your application
- data fetching: where your data lives and how to get it
- rendering: when and where you render static or dynamic content
- infrastructure: where you deploy, store and run your application code
- scalability: how your application adapts as your team, data and traffic grow

- client-side vs server-side: describes where web application code runs
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
    - endpoint url: https://api.example.com/books
    - app.get('/books', ...): This line defines an endpoint that listens for GET requests at the /books URL.
- handler: processes http requests and returns responses

when a user visits a webpage:

1. client (browser) makes a request to the server
1. server returns an HTML file to the browser
1. the browser reads HTML and then constructs the DOM
  - DOM (Document Object Model): representation of the HTML elements of a webpage as a tree of nodes and objects
    - each part of the page becomes a node
    - each node can be accessed and manipulated by programming languages like Javascript

[http](/networking.md#http-(hypertext-transfer-protocol))

- HTML represents the **initial** page content
- DOM represents the **updated** page content

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

- composed by a single html page
- client-side rendering
- asynchronous data fetching
  - allows the app to update without refreshing the entire page
- slow initial load time
- examples
  - gmail, facebook, trello

### Multi-Page Application

- each interaction/request loads a new HTML page from the server
- server-side rendering
  - each page fully reloads when the user navigates
- examples
  - e-commerce sites, news sites

### Progressive Web Application (PWA)

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

- client-side/front-end:
  - user interface and user experience
  - technologies: HTML, CSS, JavaScript, frameworks (React, Vue, Angular)
  - makes HTTP requests to the back-end
  - handles user interactions and data display

- server-side/back-end:
  - business logic, database operations
  - authentication, authorization
  - API endpoints and services
  - technologies: Node.js, Python, Go, databases (SQL, NoSQL)
  - receives HTTP requests from the front-end and sends responses

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

3. cd into `schema` directory and run the following command to create a new SQL file

```bash
goose -s create create_user sql
```

- `-s`: use sequential numbering for new migrations (random numbers if flag is not used)
- `create`: creates new migration file with the current timestamp
- `create_user`: name of the sql query
- this command will create an sql file with a name in the following format `querynumber_name.sql` (e.g. `001_create_users.sql`)

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

## authentication

> process of verifying the identity of a user

- methods of authentication
  - passwords
    - don't store passwords as plain text
    - use hashing
    - don't allow weak passwords

## authorization

> process of determining what actions a user can perform within a system
