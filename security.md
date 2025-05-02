
[authentication](./auth.md)

# security

## concepts

- firewall: network security system that monitors and controls incoming and outgoing network traffic
  - can be hardware-based and/or software-based
  - prevents unauthorized access

- multi-factor authentication (MFA): requires two or more forms of identification to access an account
  - two-factor authentication (2FA): requires two forms of identification to access an account
    - something you know (password)
    - something you have (phone, token)
    - something you are (biometrics)
  - adaptive multi-factor authentication: uses multiple factors to determine which form of authentication to use
    - e.g. if a user logs in from a new device or a new location, they may be required to provide additional verification

- hashing: process that takes an input of arbitrary size and returns a fixed-size string of characters, known as a hash
  - hash value output is unique to the input
  - one-way function: easily generate the hash, but you can't reconstruct the original data from its hash
  - deterministic: same input, same hash output
  - fixed output size: the size of the input doesn't affect the length of the output
  - avalanche effect: even a small change in the input data creates a completely different output
  - has multiple implementation algorithms (e.g. Bcrypt, SHA-3, MD5, etc)
    - always use Bcrypt
    - [avoid MD5 and SHA-3](https://blog.boot.dev/cryptography/how-sha-2-works-step-by-step-sha-256/#can-i-use-sha-256-to-hash-passwords)

- how can i reconstruct the original github repo from its .git folder with `git cat-file`?
  - because git is not decoding the hash, it's using the hash to look up the corresponding object in `.git`

## principles

- confidentiality: only authorized users can access the data
- integrity: data is not tampered with during transit
- availability: data is available when needed
- authenticity: data is from a trusted source
- irretractability: data origin is undeniable
- non-retroactivity: data cannot be changed retroactively without detection
  - e.g. a digitally signed document cannot have its signature removed without detection
- legality: data processing must comply with the law

## threats

- data theft: unauthorized access to data
- traffic interception: unauthorized access to data in transit
- brute force attacks: trying all possible combinations to guess a password
- DDOS (Distributed Denial of Service) attack: prevent legitimate users from accessing the service
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
- SQL Injection: malicious SQL code injected into application queries to manipulate database
- XSS (Cross-Site Scripting): malicious scripts injected into trusted websites to run in users' browsers
- CSRF (Cross-Site Request Forgery): tricks users into submitting malicious requests to websites they're authenticated with

## types of authentication

### authentication methods

- id + password
  - id can be email, username, etc
- 3rd party authentication
  - e.g. sign in with google, sign in with github, etc
- magic links
  - e.g. `https://example.com/login?token=...`
- api keys

### authentication factors

1. something you know

- id + password
- pin
- security questions

2. something you have

- phone
- hardware token
- smart card

3. something you are

- biometrics:
  - fingerprint
  - face recognition
  - iris scan
  - voice recognition

## encryption

> encryption: process of converting plain text into unreadable ciphertext that can only be decoded back to plain text using a specific key

- ciphertext: encrypted text that is unreadable and can only be read with an encryption key
- if unauthorized party gains access to the encrypted text, they won't be able to read it without the decryption key

- types of encryption:
  - symmetric: user the same key for encryption and decryption
  - asymmetric: uses different keys for encryption (public key) and decryption (private key)

### asymmetric encryption

plain text -> encryption (with public key) -> cipher text -> decryption (with private key) -> plain text

- public key
  - can be shared openly with anyone
  - used to encrypt data
  - if lost, can be regenerated from private key
- private key
  - must be kept secret
  - used to decrypt data
  - file becomes useless if lost

## Public Key Infrastructure (PKI)

> comprehensive framework that enables secure electronic communication through the management of digital certificates and encryption keys. It encompasses a combination of hardware, software, policies, and procedures necessary for the creation, distribution, management, storage and revocation of digital certificates

- establishes trust in digital communications
- provides mechanisms for encryption
- enables authentication
- ensures data integrity across various applications and services

### key components of PKI

1. Digital Certificates: These are electronic credentials that bind a public key to an entity (such as a user or device), ensuring that the entity is who it claims to be. Digital certificates function similarly to a passport in verifying identity

2. Certificate Authority (CA): This trusted entity issues and manages digital certificates. The CA verifies the identities of entities requesting certificates and signs them to confirm their authenticity

3. Registration Authority (RA): This component verifies the identity of users or devices before they can obtain a digital certificate from the CA. The RA acts as a mediator between the user and the CA

4. Public and Private Keys: PKI relies on asymmetric encryption, which uses a pair of keys:
   - Public Key: Available to anyone; used for encrypting data
   - Private Key: Kept secret by the owner; used for decrypting data encrypted with the corresponding public key

### how PKI works

- PKI facilitates secure communications by enabling:
  - encryption: data can be encrypted using the recipient's public key, ensuring that only the recipient can decrypt it with their private key
  - authentication: digital certificates confirm the identities of users and devices, preventing impersonation and ensuring that data is sent to the correct party
  - integrity: digital signatures can be used to verify that data has not been altered during transmission

### applications of PKI

- PKI is widely used in various domains, including:
  - E-commerce: Securing online transactions
  - Internet Banking: Protecting sensitive financial information
  - Secure Email: Ensuring confidentiality and authenticity in email communications
  - IoT Devices: Managing secure connections between connected devices

## firewalls

> network security system that monitors and controls incoming/outgoing network traffic based on predetermined security rules

TODO

## OWASP (Open Web Application Security Project)

> non-profit organization focused on improving software security

top 10 web application security risks:

1. broken access control: users access data they're not authorized to view or modify
1. cryptographic failures: encryption and hashing aren't implemented correctly
1. injection: malicious code is not executed
1. insecure design: architecture/design of the application are not secure
1. security misconfiguration: application, server or database are not configured securely
1. vulnerable and outdated components: outdated libraries, frameworks and software with known vulnerabilities
1. identification and authentication failures
1. software and data integrity failures
1. security logging and monitoring failures: insufficient logging and monitoring for security-related events
1. server-side request forgery: when an attacker can make the server initiate requests to internal or external resources on his behalf

### OWASP SAMM (Open Web Application Security Project Software Assurance Maturity Model)

> allows organizations to evaluate security practices

- domains:
  - governance
    - strategy and metrics
    - policy and compliance
    - education and guidance
  - design
    - threat assessment
    - security requirements
    - secure architecture
  - implementation
    - secure build
    - security deployment
    - defect management
  - verification
    - architecture assessment
    - requirements-driven testing
    - security testing
  - operation
    - incident management
    - environment management
    - operational management

- each domain has 3 maturity levels:
  - level 1: define metrics with insights
  - level 2: establish unified strategic roadmap for software security within the organization
  - level 3: align security efforts with the relevant organizational indicators and asset values

## database security

- main objectives
  - confidentiality: prevent unauthorized access to sensitive data
  - integrity: ensure data is accurate and has not been tampered with
  - availability: ensure data is data is accessible when needed by authorized users

- access control models:
  - DAC (Discretionary Access Control): access is controled by the owner of the resource
  - MAC (Mandatory Access Control): access is controled by a central authority, based on strict policies
  - RBAC (Role-Based Access Control): access is based on roles assigned to users (e.g. "admin", "manager", "employee")

- create user
  - postgresql: `CREATE USER myuser WITH PASSWORD 'mypassword';`
  - mysql/mariadb: `CREATE USER 'myuser'@'localhost' IDENTIFIED BY 'mypassword';`

role management in postgresql:

```sql
-- Create a role without login
CREATE ROLE readonly;

-- Grant privileges to the role
GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;

-- Create a user and assign the role
CREATE USER analyst WITH PASSWORD 'securepass';
GRANT readonly TO analyst;
```

role management in mysql/mariadb:

```sql
-- Create a role
CREATE ROLE 'readonly';

-- Grant privileges to the role
GRANT SELECT ON mydatabase.* TO 'readonly';

-- Create a user
CREATE USER 'analyst'@'localhost' IDENTIFIED BY 'securepass';

-- Assign the role to the user
GRANT 'readonly' TO 'analyst'@'localhost';

-- Set default active role for user (optional)
SET DEFAULT ROLE 'readonly' TO 'analyst'@'localhost';
```

- `GRANT`: gives permissions to a user or role
  - `GRANT SELECT, INSERT ON Employees TO alice;`
- `REVOKE`: remove previously granted permissions
  - `REVOKE INSERT ON Employees FROM alice;`
- `DENY`: explicitly block a permission
  - `DENY SELECT ON Employees TO alice;`

## database security threats

- SQL Injection: injecting malicious SQL into queries to dump data, modify/delete records or bypass login
  - causes
    - poorly filtered strings
    - incorrect type handling
    - signature evasion
    - filter bypassing
    - blind sql injection
- credential stuffing: using leaked credentials from other breaches to access databases
- brute force attacks: repeatedly trying username/password combinations until one works
- Denial of Service (DoS): flooding the database with queries to exhaust resources
- Man-in-the-Middle (MitM): intercepting database traffic, especially over insecure connections

## tools

- [nmap](./cli_tools.md#nmap): network scanner that discovers hosts and services on a network
- wireshark: network protocol analyzer that captures and inspects data
- metasploit: penetration testing framework that helps find vulnerabilities in systems
- netcat (`nc`): networking utility that reads/writes data across network connections

---

- SDL (Security Development Lifecycle)
  - process framework that integrates security and privacy considerations into every phase of software development
  - goal: reduce vulnerabilities and development costs by embedding security early
  - lifecycle stages: requirements => design => implementation => verification => release => response

- CLASP (Comprehensive, Lightweight Application Security Process)
  - set of security best practices for existing software development processes

- BSIMM (Building Security in Maturity Model)

- NIST SSDF (Secure Software Development Framework)
