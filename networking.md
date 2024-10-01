
[web development](./web_development.md)

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
1. data link => 2 words
1. network => n3twork
1. transport => tr4nsport
1. session => 5e55ion
1. presentation
1. application

## TCP/IP model

> [!IMPORTANT]
> TCP/IP is the most widely used model
> tcp/ip model != tcp and ip protocols

- application layer
  - application layer
  - presentation layer
  - session layer
- transport layer
  - transport layer
- network layer
  - network layer
- network access layer
  - data link layer
  - physical layer

### application layer

#### DNS (Domain Name System)

> resolves domain names to IP addresses

- URL: https://youtube.com
  - domain name/hostname: youtube.com

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
  - 1xx: informational
  - 2xx: success
    - 200 OK: The request was successful, and the server is returning the requested data
  - 3xx: redirection
    - 301 Moved Permanently: The requested resource has been moved to a new URL
  - 4xx: client error
    - 400 Bad Request: The server could not understand the request due to invalid syntax
    - 401 Unauthorized: The request requires authentication, and the client has not provided valid credentials
    - 404 Not Found: The requested resource could not be found on the server
  - 5xx: server error
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

#### SSL (Secure Sockets Layer)

- provides secure communication over a computer network
- deprecated and insecure

#### TLS (Transport Layer Security)

- successor to SSL
- addresses flaws in SSL
- provides privacy and data integrity between two communicating applications
- TLS handles the encryption of HTTP requests and responses
  - all data transferred via HTTPS is secured by TLS
- TLS handshake
  - when a browser connects to an HTTPS website, it starts with a TLS handshake
  - defines which:
    - version of TLS will be used
    - encryption algorithm will be applied
    - exchange of digital certificates to verify server's identity
  - once the handshake is complete, HTTP data is transmitted over the secured TLS connection

HTTPS vs TLS
- HTTPS is the protocol for secure web communication
- TLS is the security layer that HTTPS uses to encrypt and secure that communication

### transport layer

#### TCP (Transmission Control Protocol)

> enables application programs and computers to exchange messages over a network

- SYN-ACK process:
1. client sends SYN (Synchronize) packet to server
1. server sends SYN-ACK (Synchronize-Acknowledgment)
1. clients sends ACK packet to server (Acknowledgment)
1. connection established

### network layer

#### IP (Internet Protocol)

> google.com (domain name) => 142.251.135.142 (ip address)

- IPv4
  - 32 bits = 4 bytes
  - `byte.byte.byte.byte` <=> `142.251.135.142`
- IPv6
  - 128 bits = 16 bytes
  - `2bytes:2bytes:2bytes:2bytes:2bytes:2bytes:2bytes:2bytes` <=> `2800:3f0:4004:809::200e`
  - `16bits:16bits:16bits:16bits:16bits:16bits:16bits:16bits` <=> `2800:3f0:4004:809::200e`
    - double colon `:` represents 16 zero bits

> [!TIP]
> IPv4 32 bits => 432 => 4 bytes and 32 bits
> IPv4 bytes * 2² bytes => IPv6
> 32 bits * 2² = 128 bits

### network access layer

#### MAC (Media Access Control)

- unique identifier attributed to each network interface of a device connected to the network
- MAC address uses 6 bytes = 48 bits
- `byte:byte:byte:byte:byte:byte`
- example: `47:3E:2A:B2:11:24`

---

## default ports

- FTP 20/21
- SSH 22
- TELNET 23
- SMTP 25
- DNS 53
- DHCP 67/68
- HTTP 80
- HTTPS 443
