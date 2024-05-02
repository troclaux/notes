## HTTP/HTTPS
- error codes:
  - 1xx: informational
  - 2xx: success
  - 3xx: redirection
  - 4xx: client error
  - 5xx: server error

## SSL (Secure Sockets Layer)

- provides secure communication over a computer network
- deprecated and insecure

## TLS (Transport Layer Security)

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
