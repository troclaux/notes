
# API (Application Programming Interface)

> rules and protocols that allow different software programs to communicate with each other

[what are apis](https://www.redhat.com/en/topics/api/what-are-application-programming-interfaces)
[API design - principles and best practices](https://yourbasic.org/algorithms/your-basic-api/)
[rapid APIs](https://rapidapi.com/)
[public APIs](https://github.com/public-apis/public-apis)
[golang API](/golang.md#restful-api-with-http-requests)

- defines the structure of requests and responses
- APIs allow services to communicate with other services without knowing how they're implemented
  - service: software component that provides specific functionality
- normally they implement [CRUD](/databases.md#crud) operations
  - Create, Read, Update, Delete
  - can use http verbs: Get, Post, Put, Delete

- types of APIs:
  - RESTful APIs
  - SOAP (Simple Object Access Protocol) APIs
  - RPC (Remote Procedure Call) APIs
  - WebSocket APIs

- scope of use:
  - private APIs: used internally in enterprises to connect systems within the business
  - public APIs: open to the public
  - Composite APIs: combine multiple APIs to address complex requirements

## endpoints

- API endpoint: A specific URL path where a client can interact with a specific function/resource of an API
  - complete URL: `https://api.bookstore.com/books/123`
    - base URL: `https://api.bookstore.com`
    - endpoint: identifies the specific resource/action (e.g. `/books/123`)
      - resource: `books`
      - resource id: `123`
  - must be secured with:
    - authentication tokens
    - API keys

- collection: collection of resources
  - endpoint to get a collection: `GET /api/users`
    - returns all users
- singleton: single instance of a resource
  - endpoint to get a singleton: `GET /api/users/{user_id}`
    - returns a user

> [!TIP]
> use versioning in the API to avoid breaking changes
> e.g. `https://api.bookstore.com/v1/books/123`

## RESTful API (REpresentational State Transfer)

> architectural style for an application program interface (API) that uses HTTP requests to access and use data

- representational: a representation of the resource (e.g. JSON, XML, HTML, text, images) is used to interact with the resource
- stateless: each request from a client to a server must contain all the information necessary to understand the request
  - server does not store any client-specific session data between requests
  - each request is independent and can be understood in isolation
- implementation of the client and server can be created independently of one another
- while REST is most commonly implemented over HTTP, it is not strictly bound to it

## testing

- recommended tools:
  - `curl`
  - jest + supertest

## 5 commandments

1. tell me what this thing is
1. tell me what it does
1. don't tell me how it works
1. grant me the right to use it
1. don't change it

- keep it simple
  - don't use complicated constructs where simple ones will do
  - one package, one idea
  - just say no
  - math is simple
- give it time
- show, don't tell
  - create tutorials
  - use examples
- tools of the trade
  - keep it consistent
  - write functions that need little and give much
  - discover a well-fitting interface
  - make it generic
  - names, keep them short and sweet
