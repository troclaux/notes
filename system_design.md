
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
