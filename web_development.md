
# web development

## basic concepts of web application

- user interface
- routing: how users navigate between different parts of your application
- data fetching: where your data lives and how to get it
- rendering: when and where you render static or dynamic content
- infrastructure: where you deploy, store and run your application code
- scalability: how your application adapts as your team, data and traffic grow

- client: software that requests and receives data/services from a server over a network (usually the internet)
  - e.g., web browsers, mobile apps
- server: computer program/device that provides data/services to clients over a network (usually the internet)
  - e.g., web servers, database servers, file servers
  - web server: handles mainly http requests
    - static file server: web server that serves files from the host machine
      - serves html, css, js, images, etc.
  - dynamic web server: web server that generates different content for different users
    - e.g., based on user input, user authentication, etc

- service: functionality or resource that a service provides to an client upon request

- endpoint: specific URL where the server provides a resource/service
  - example:
    - endpoint url: https://api.example.com/books
    - app.get('/books', ...): This line defines an endpoint that listens for GET requests at the /books URL.
- handler: processes http requests and returns responses

when a user visits a webpage:

1. client (browser) makes a request to the server
2. server returns an HTML file to the browser
3. the browser reads HTML and then constructs the DOM
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
