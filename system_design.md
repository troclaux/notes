
# system design

> process of defining the architecture, components, and data flow of a system to meet specific requirements, ensuring scalability, reliability and maintainability

- service: self-contained unit of functionality that performs a specific task and communicates with other parts of the system over a network
  - runs independently, often in its own server or container
  - talks with other services via APIs
  - can scale independently
  - e.g. payment service, notification service, authentication service, etc
- vertical scaling: increasing/decreasing the power of a single server (e.g. more CPU, RAM)
- horizontal scaling: adding more servers to distribute the load
- load balancing: distributes traffic across multiple servers
  - e.g. nginx, cloudfare load balancer
- api gateway: intermediates data traffic between clients and servers
- database design
  - sharding: splitting data across multiple databases for scalability
  - replication: duplicating data across multiple servers for fault tolerance
  - denormalization
- caching: stores frequently accessed data in-memory to reduce database load and improve performance
  - common strategies:
    - write-through cache: updates cache and database simultaneously
    - lazy-loading cache: updates cache only when requested
  - examples: [Redis](./redis.md), Memcached
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

## steps in system design

- understand requirements (functional and non-functional
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
  - peer-to-peer (P2P): decentralized model where wach device (peer) acts **both** as a client and a server

### CAP theorem

- a distributed system can only guarantee 2 out of following 3 properties at the same time:
  - Consistency: every read gets the most recent write or an error
  - Availability: every request gets a non-error response, even if it's not the latest data
  - Partition tolerance: the system keeps working even if network problems split it into disconnected parts

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
