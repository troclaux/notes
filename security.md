
- DDOS attack (Distributed Denial of Service attack)
  - a cyberattack where multiple systems flood a target with traffic to make it unavailable
  - overwhelms network/server resources so legitimate users can't access services
  - uses botnets to generate massive traffic
    - botnets: infected computers
  - defenses:
    - rate limiting
    - blackholing

- slowloris attack: DoS attack where malicious clients hold connections open by slowly sending HTTP headers or data
  - sends partial HTTP requests that never complete
  - ties up server's connection pool so legitimate traffic can't be handled
  - particularly effective against Apache and similar thread-based servers
  - mitigations:
    - rate limiting connections per IP
    - increasing max connections
    - using timeout settings
    - using reverse proxies (nginx, haproxy)
    - implementing connection limits

## public and private keys

plain text -> encryption (with public key) -> cipher text -> decryption (with private key) -> plain text

- public key
  - can be shared openly with anyone
  - used to encrypt data
  - if lost, can be regenerated from private key
- private key
  - must be kept secret
  - used to decrypt data
  - file becomes useless if lost

## Public Key Infrastructure (PKI) is a comprehensive framework that enables secure electronic communication through the management of digital certificates and encryption keys. It encompasses a combination of hardware, software, policies, and procedures necessary for the creation, distribution, management, storage, and revocation of digital certificates.

## Key Components of PKI

1. Digital Certificates: These are electronic credentials that bind a public key to an entity (such as a user or device), ensuring that the entity is who it claims to be. Digital certificates function similarly to a passport in verifying identity

2. Certificate Authority (CA): This trusted entity issues and manages digital certificates. The CA verifies the identities of entities requesting certificates and signs them to confirm their authenticity

3. Registration Authority (RA): This component verifies the identity of users or devices before they can obtain a digital certificate from the CA. The RA acts as a mediator between the user and the CA

4. Public and Private Keys: PKI relies on asymmetric encryption, which uses a pair of keys:
   - Public Key: Available to anyone; used for encrypting data
   - Private Key: Kept secret by the owner; used for decrypting data encrypted with the corresponding public key

## How PKI Works

PKI facilitates secure communications by enabling:
- Encryption: Data can be encrypted using the recipient's public key, ensuring that only the recipient can decrypt it with their private key
- Authentication: Digital certificates confirm the identities of users and devices, preventing impersonation and ensuring that data is sent to the correct party
- Integrity: Digital signatures can be used to verify that data has not been altered during transmission

## Applications of PKI

PKI is widely used in various domains, including:
- E-commerce: Securing online transactions
- Internet Banking: Protecting sensitive financial information
- Secure Email: Ensuring confidentiality and authenticity in email communications
- IoT Devices: Managing secure connections between connected devices

In summary, PKI is essential for establishing trust in digital communications, providing mechanisms for encryption, authentication, and data integrity across various applications and services
