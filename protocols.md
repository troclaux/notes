## layer 1


## layer 2

### MAC (Media Access Control)

- unique identifier attributed to each network interface of a device connected to the network
- MAC address uses 6 bytes = 48 bits
- example: `47:3E:2A:B2:11:24`

## layer 3

### IP (Internet Protocol)

> unique address that is used to send data across the internet

### HTTP/HTTPS
- error codes:
  - 1xx: informational
  - 2xx: success
  - 3xx: redirection
  - 4xx: client error
  - 5xx: server error

## layer 4

### TCP (Transmission Control Protocol)

> enables application programs and computers to exchange messages over a network

### SSL (Secure Sockets Layer)

- provides secure communication over a computer network
- deprecated and insecure

### TLS (Transport Layer Security)

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

## layer 5
## layer 6
## layer 7

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
