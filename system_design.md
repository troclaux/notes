
# system design

> process of defining the architecture, components and data flow of a system to meet specific requirements, ensuring scalability, reliability and maintainability

- service: self-contained unit of functionality that performs a specific task and communicates with other parts of the system over a network
  - runs independently, often in its own server or container
  - talks with other services via APIs
  - can scale independently
  - e.g. payment service, notification service, authentication service, etc
- vertical scaling: increasing/decreasing the power of a single server (e.g. more CPU, RAM)
- horizontal scaling: adding more servers to distribute the load
- api gateway: intermediates data traffic between clients and servers
  - single entry point that manages client requests, routing them to appropriate backend services
  - can also handle authentication, rate limiting and logging
- database design
  - sharding: splitting data across multiple databases for scalability
  - replication: duplicating data across multiple servers for fault tolerance
  - denormalization: intentionally introducing redundancy in the database to optimize read performance, at the cost of write complexity
- message queues and event-driven architecture
  - used for handling background tasks and asynchronous processing
  - examples: Kafka, RabbitMQ, AWS SQS
- API Design
  - REST APIs: Stateless, using HTTP methods (GET, POST, etc)
  - GraphQL: Allows clients to request specific data
  - gRPC: High-performance API communication using Protocol Buffers
- security considerations
  - [authentication and authorization](./auth.md): OAuth, JWT, API keys
  - rate limiting: prevents API abuse (e.g., using Cloudflare, Nginx)
  - encryption: HTTPS (TLS), encrypting sensitive data

- resilience: capacity of a system to handle and recover from failures gracefully
  - handles problems like errors, crashes, or even cyber-attacks
- caching: store frequently accessed data in faster storage layers, allowing systems to:
  - avoid repeated expensive operations (like db queries with indexing or api calls)
  - continue serving data even when the original source is temporarily unavailable (fault tolerance)
  - there are local or distributed caching strategies (e.g. redis, CDN)
    - CDN (Content Delivery Network): caches static content (images, css, js) at edge locations globally
      - reduces latency by serving users from the nearest data center
      - examples: cloudflare, aws cloudfront, akamai
- fallback: backup mechanism or secondary option that is used when the primary system/service/component fails or becomes unavailable
  - examples:
    - graceful degradation: system gradually loses functionality rather than failing completely when components become unavailable
    - static fallback pages: serving static HTML when dynamic backend is down
    - alternate services: switch to a backup API or microservice
- circuit breaker: pattern to prevent cascading failures in microservices or distributed systems
  - it monitors for repeated failures
  - if thresholds are exceeded, it "opens the circuit" to block further calls
  - after a time period (timeout), retries a single request to check if the service is available again
- disaster recovery: strategies for recovering from major outages (e.g. snapshots, offsite backups)
  - objectives:
    - minimize Recovery Time Objective (RTO)
    - preserve data consistency and accuracy
    - ensure business continuity
  - contingency planning: broader term for any backup or alternative action plans, same as disaster recovery
- GSLB (Global Server Load Balancing): distributes user traffic across multiple geographic regions using DNS or application-level logic
- failover: auto-switching to backup systems on failure (active-passive, active-active, etc)
  - active-active site: redundant systems operating simultaneously in different regions/data centers
- load balancing: distributes traffic across multiple servers
  - e.g. nginx, cloudfare load balancer
  - prevents overloading any single node
  - improve scalability and fault tolerance
  - High Availability (HA): ensures the system remains accessible even if some components fail
  - can be implemented as active-passive or active-active

## steps in system design

- understand requirements (functional and non-functional)
- define high-level architecture
  - microservices vs monolithic
  - define databases, caching, load balancing and api gateways
- design database and storage
  - sql vs nosql
  - plan indexing, sharding and replication
- plan scaling strategies
  - auto-scaling servers
  - caching mechanisms
  - load balancing
- design APIs and data flow
  - REST vs GraphQL vs gRPC
  - api gateway to manage traffic

## architectural styles

- architecture: how software is structured
- coupling: between data and presentation logic
  - has that same name as [coupling](/software_engineering.md#software-architecture-patterns), but it's a different concept

### monolithic architecture

> single, large program that contains all functionality for both front end and back end of an application

- advantages
  - simpler to develop
  - simpler to deploy
- disadvantages
  - harder to scale

### microservices architecture

- often implemented as a distributed system

- application split into small independent services
- each service handles specific business function
- easier to scale and maintain individual components
- more complex to develop and deploy initially
- horizontal and vertical scaling
- decoupled: front end and back end are separated into different codebases
  - e.g. front end is hosted by a static file server on one domain and back end is hosted on a subdomain by a different server
  - advantages
    - scale according to demand for each part
    - better separation of concerns
    - hosted in separate servers with different technologies

- decentralized data management
- smaller services than SOA

### Service-Oriented Architecture (SOA)

> software is organized as a collections of loosely coupled services that communicate over a network

- break functionality into independent services
  - each service does one specific job and talks to other services through well-defined interfaces (usually APIs)

- horizontal scaling
- loose coupling
- interoperability: services can use different technologies (java, python, .net, etc)
- centralized data management

## distributed systems

> group of independent computers/nodes that work together to appear as a single system to the user

- objectives
  - scalability: ability to grow smoothly when demand increases
  - fault tolerance: if one machine fails, the system keeps running
  - resource sharing: machines share data, storage
  - transparency: hide complexity from users

- nodes: individual computers or servers in the system
- consistency: nodes must agree on the system's state, even if network issues occur

- basic models:
  - client-server: centralized model where clients request resources or services from servers
  - peer-to-peer (P2P): decentralized model where each device (peer) acts **both** as a client and a server

### CAP theorem

- a distributed system can only guarantee 2 out of following 3 properties at the same time:
  - Consistency: every read gets the most recent write or an error
  - Availability: every request gets a non-error response, even if it's not the latest data
  - Partition tolerance: the system keeps working even if network problems split it into disconnected parts

> [!IMPORTANT]
> Partition tolerance is non-negotiable in real distributed systems
> most systems choose either CP or AP

### pub/sub architecture

> pattern that software systems can use to communicate

- pub/sub = publish/subscribe

- scalable
- reliable
- maintainable

- examples
  - RabbitMQ
  - kafka
  - google cloud pub/sub

1. client sends message to a message broker
1. broker sends message to all subscribers

- point-to-point vs pub/sub:
  - point-to-point: sender => receiver
    - similar to how email works
  - pub/sub: publishers => broker => subscribers
    - similar to how instagram works

## types of services

- web server: serves web pages
  - e.g. apache, nginx
- application server: runs application code
  - e.g. tomcat, jboss
- database server: stores data
  - e.g. mysql, postgresql
- cache server: stores data in memory
  - e.g. redis, memcached
- message broker: routes messages between services
  - e.g. rabbitmq, kafka
- load balancer: distributes traffic between servers
  - e.g. nginx, haproxy
- reverse proxy: forwards requests to servers
  - e.g. nginx, haproxy
