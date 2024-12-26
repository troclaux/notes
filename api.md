
# API (Application Programming Interface)

> rules and protocols that allows two software programs to communicate with each other

[what are apis](https://www.redhat.com/en/topics/api/what-are-application-programming-interfaces)
[API design - principles and best practices](https://yourbasic.org/algorithms/your-basic-api/)
[rapid APIs](https://rapidapi.com/)
[public APIs](https://github.com/public-apis/public-apis)
[golang API](/golang.md#restful-api-with-http-requests)

- defines the structure of requests and responses
- APIs allow services to communicate with other services without knowing how they're implemented
- normally they implement CRUD operations
  - Create, Read, Update, Delete
  - can use http verbs: Get, Post, Put, Delete

- types of APIs:
  - RESTful APIs
  - SOAP (Simple Object Access Protocol) APIs
  - RPC (Remote Procedure Call) APIs
  - Websocket APIs

- scope of use:
  - private APIs: used internally in enterprises to connect systems within the business
  - public APIs: open to the public
  - Composite APIs: combines multiple APIs to address complex requirements

- API endpoints: URLs that are used to access the resources
  - complete URL: `https://api.bookstore.com/books/123`
    - base URL: `https://api.bookstore.com`
    - endpoint: `/books/123`
  - must be secured with:
    - authentication tokens
    - API keys

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
- does not necessarily uses HTTP as underlying protocol for communication

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
