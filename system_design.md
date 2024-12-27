
# system design

## pub/sub architecture

> pattern that software systems can use to communicate

- makes architectures
  - scalable
  - reliable
  - maintainable
- pub/sub = publish/subscribe
- examples
  - RabbitMQ
  - kafka
  - google cloud pub/sub

1. client sends message to a message broker
1. broker sends message to all subscribers

- point-to-point vs pub/sub
  - point-to-point: sender -> receiver
    - similar to how email works
  - pub/sub: publishers => broker => subscribers
    - similar to how instagram works

## monolithic vs microservices

- architecture: how software is structured
- coupling: between data and presentation logic
  - has that same name as [coupling](/software_engineering.md#software-architecture-patterns), but it's a different concept

- monolithic: single, large program that contains all functionality for both front-end and back-end of an application
  - advantages
    - easier to develop
    - easier to deploy
- decoupled: front-end and back-end are separated into different codebases
  - e.g. front-end is hosted by a static file server on one domain and back-end is hosted on a subdomain by a different server
  - advantages
    - scale according to demand for each part
    - better separation of concerns
    - hosted in separate servers with different technologies
