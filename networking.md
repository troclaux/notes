
[web development](./web_development.md)
[protocols](./protocols.md)

## OSI model

[osi-model explained in a simple way](https://osi-model.com)

layers:
1. physical
  - transmits raw data bits over physical medium
  - USB, RJ45
2. data link/enlace
  - ensures error-free data transfer between two connected nodes
  - defines format of data on the network
  - transfers data between network nodes to WAN (Wide Area Network)
  - MAC addresses, ethernet, switches
3. network
  - decides which physical path the data will take
  - IP
4. transport
  - transmits data using transmission protocols
  - TCP, UDP, port numbers
5. session
  - maintains connections and controls ports/sessions
  - syn/ack
6. presentation
  - ensures data is in a readable format
  - encryption happens here
  - ASCII, JPEG
7. application
  - HTTP, FTP

> [!TIP]
> layer order

1. physical
2. data link => 2 words
3. network => n3twork
4. transport => tr4nsport
5. session => 5e55ion
6. presentation
7. application

## TCP/IP model

- tcp/ip model != tcp and ip protocols

1. application layer
  - application layer
  - presentation layer
  - session layer
2. transport layer
  - transport layer
3. network layer
  - network layer
4. network access layer
  - data link layer layer
  - physical layer

### network layer 3

- IPv4
  - 32 bits = 4 bytes
- IPv6
  - 128 bits = 16 bytes

> [!TIP]
> IPv4 32 bits => 432 => 4 bytes and 32 bits
> IPv4 bytes * 2² bytes => IPv6
> 32 bits * 2² = 128 bits

### application layer 7

#### HTTP (HyperText Transfer Protocol)

> defines how requests and responses between a client and a server should be formatted and transmitted

- foundation of data communication on the World Wide Web
- stateless protocol
  - this means that each request from a client to a server is independent, with no memory of previous interactions

- uses client-server architecture
  - client: device or application that initiates an HTTP request
    - e.g. web browser, mobile app
  - server: device or application that processes the request and sends back a response
    - e.g. web server
- requests and responses
  - request: contains the following elements:
    - HTTP method (like GET or POST)
    - URL
    - HTTP version
    - headers: key-value pairs that provide additional information about the request (e.g. content type, user agent)
    - body (optional): contains data being sent from the client to the server (e.g. form data or file uploads)
  - response
    - status line: contains:
      - HTTP version
      - status code (e.g. 404, 200)
      - headers
    - body: contains main content of HTTP response
      - HTML for web pages
      - JSON or XML for API responses
      - binary data for file downloads

- HTTP methods:
  - GET: requests data from the server (e.g. retrieving a webpage)
  - POST: submits data to the server (e.g. submitting a form or uploading a file)
  - PUT: updates or replaces existing data on the server
  - DELETE: deletes data from the server
  - PATCH: partially updates data on the server
  - HEAD: similar do GET, but only retrieves headers, not the body
  - OPTIONS: used to describe the communication options available for the target resource

- HTTP status codes:
  - 200 OK: The request was successful, and the server is returning the requested data
  - 301 Moved Permanently: The requested resource has been moved to a new URL
  - 400 Bad Request: The server could not understand the request due to invalid syntax
  - 401 Unauthorized: The request requires authentication, and the client has not provided valid credentials
  - 404 Not Found: The requested resource could not be found on the server
  - 500 Internal Server Error: server encountered an unexpected condition that prevented it from fulfilling the request

example of HTTP request

```
GET /index.html HTTP/1.1
Host: www.example.com
User-Agent: Mozilla/5.0
Accept: text/html
```

example of HTTP response

```
HTTP/1.1 200 OK
Content-Type: text/html
Content-Length: 1024

<html>
  <body>
    <h1>Welcome to Example!</h1>
  </body>
</html>
```

- HTTP URL (Uniform Resource Locator)
  - `http://`: specifies that the http protocol will be used for communication

#### HTTPS

> extension of HTTP that uses encryption (via SSL/TLS) to secure the communication between the client and the server
