
# RabbitMQ

> message broker that helps applications communicate by sending messages through queues

- effects:
  - reduce coupling: senders and receivers don't need to run at the same time
  - improve scalability: multiple consumers can process messages in parallel
  - reliability: supports message durability, acknowledgments, retries
  - flexibility: complex routing via different exchange types

- use cases:
  - job queues (e.g. background processing)
  - event-drive systems
  - logging and monitoring pipelines
  - decoupling microservices

- based on AMQP (Advanced Message Queuing Protocol)
- producer: application that sends messages
  - e.g. web app sending a job to be processed (like sending a welcome email)
- consumer: application that receives and processes messages form the queue
  - e.g. background worker that reads messages and sends out emails
- queue: buffer where messages are stored until they're consumed
  - similar to a mailbox: producers put letters (messages) in, consumers take them out
- exchange: decides how messages are routed to queues
  - type:
    - direct: routes based on exact match (routing key)
      - routing key: string used by an exchange to determine which queue(s) a message should go to
    - fanout: broadcasts to all queues
    - topic: routes based on pattern matching (`logs.*`, `*.error`)
    - headers: routes based on message headers
- binding: link between an exchange and a queue, with optional rules (like routing keys)

