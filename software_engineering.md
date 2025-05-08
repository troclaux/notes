
# software engineering

> application of engineering principles to design, develop, test and maintain software over a long time

- [sequential models](#sequential-models)
  - [waterfall](#waterfall)
  - [V](#v-model)
- [iterative models](#iterative-models)
  - [incremental models](#incremental-models)
    - [RAD](#rad-rapid-application-development)
    - [agile methodology](#agile-methodology)
      - [scrum](#scrum)
      - [XP](#xp-extreme-programming)
      - [kanban](#kanban)
      - TDD
    - [evolutionary models](#evolutionary-models)
      - [spiral](#spiral-model)
      - prototyping

## sequential models

### waterfall

- documentation for each step
- sequential
- disadvantages:
  - projects rarely follow sequential order
  - clients normally can't define all requirements at the beginning
  - delivering Minimum Viable Product (MVP) takes too long
  - high cost to fix errors

1. requirements gathering and analysis
1. system design
1. implementation
1. integration and testing
1. deployment
1. maintenance

### V model

- variation of waterfall model

steps: left side of "V" shape => right side of "V" shape

1. requirements analysis => acceptance testing
1. system design => system testing
1. architectural design => integration testing
1. module design => unit testing
1. code implementation

## iterative models

### incremental models

- the complete system is known in advance, but developed in increments
- multiple teams developing separate finished pieces of the software that will be joined in the end

#### RAD (Rapid Application Development)

- iterative and incremental model
- short development cycle (60 - 90 days)
- advantages
  - fast development
  - fast prototyping
  - high flexibility for development process
- disadvantages
  - conflict between developers and clients caused by incorrect requirement definition

### agile methodology

> practices and values that aim to optimize software development

- adaptive, iterative and incremental development
- promotes flexibility, customer collaboration and responsiveness to change

12 key principles:

1. highest priority is satisfying the customer with early and continuous delivery of valuable software
1. high flexibility for changing requirements
1. deliver working software frequently
1. daily communication between business people and developers
1. provide support for motivated developers
1. best way to communicate is face-to-face
1. working software is primary measure of progress
1. maintain a sustainable and constant development
1. focus on technical excellence and good design to improve agility
1. maximize simplicity
1. self-organize teams for best architectures, requirements and designs
1. team must reflect and adjust to improve effectiveness

- artifacts
  - product backlog: tasks required to be completed to achieve product vision
  - sprint backlog: subset of product backlog
  - increment: all tasks completed at the end of each sprint
  - kanban board: board showing the status of tickets/issues (e.g. to do, in progress, done)
  - burndown graph: line chart that shows the remaining work versus time in a sprint
  - user stories

#### scrum

> framework that organizes work into short time cycles called sprints

- roles:
  - developers: responsible for incrementally delivering the product through sprints
  - Product Owner (PO): responsible for product vision
    - responsible for defining items in product backlog
  - scrum master: responsible for the adoption of scrum and developer enabler
    - guides Product Owner when creating product backlog
    - TIP: Scrum Master => SMooth process

- events:
  - sprint planning: entire team discusses priorities and the sprint's goals
    - happens once per sprint
    - duration: 2 hours
  - daily scrum: short meeting to keep track of progress toward the sprint's goals
    - happens daily
    - duration: 15 min
  - sprint review: development team presents sprint progress to stakeholders to receive feedback and plan next sprint
    - happens at the end of each sprint
    - duration: less than 4 hours
  - sprint retrospective: scrum team discuss how to improve the process of development
    - happens at the end of each sprint
    - duration: 45 min

sprint cycle:

1. sprint planning
1. implementation (contains daily scrum)
1. sprint review
1. sprint retrospective

- artifacts: tools that provide transparency and traceability
  - product backlog: list of everything that needs to be done to improve the product
    - managed by the PO
  - sprint backlog: tasks that need to be completed in the current sprint
    - subset of product backlog
  - product increment: sum of all completed tasks from the product backlog

#### XP (eXtreme Programming)

steps:

1. planning
1. designing
1. coding
1. testing

five core values:

1. courage
1. simplicity
1. communication
1. feedback
1. respect

- standard practices:
  - incremental planning
  - small releases
  - continuous integration
  - pair programming
  - test-driven development (TDD)
    - write unit tests before coding
  - refactoring code to improve quality
  - sustainable development (no crunch)

#### kanban

> team continuously work on tasks and use a Kanban board to keep track of progress

- tasks move from "to do" to "in progress" to "done"

#### TDD

> software development practice of writing tests before writing the actual code

- test-driven development

##### BDD

> software development approach that writes test cases in natural language that non-technical stakeholders can understand

- Behavior-Driven Development
- focuses around the behavior of the software from a user perspective
- extends tdd

## evolutionary models

- iterative and incremental models with a strong focus on risk assessment and mitigation

### spiral model

1. communication
1. planning
1. modelling
1. implementation
1. deployment

## lean methodology

> approach that focuses on maximizing value for customers while minimizing waste

1. identify the value desired by the customer
1. map the value stream
1. create flow
1. establish pull
1. seek perfection

- map the value stream: analyze the entire workflow (value stream) to identify steps that add value vs those that don't (waste)
- create flow: ensure work flows smoothly and continuously without interruption or wait times
- establish pull: production is driven by client demand, instead of predictions

lean methodology for software development changes the formula:

1. remove waste
1. increase knowledge
1. strengthen the team and promote respect
1. fast deliveries
1. build
1. optimize the process
1. postpone decisions

## requirements

- functional requirement: describes what software should do
- non-functional requirement: describes how software should accomplish a task

- normal requirements: explicitly stated by the customer
- system requirements: detailed functional and non-functional specifications
- user requirements: high-level functionalities described from the user’s point of view
- expected requirements: functionalities users assume will be present even if not explicitly requested
- domain requirements: specific to the application domain
  - examples
    - healthcare: The software must comply with HIPAA regulations for handling patient data
    - finance: The system should adhere to GAAP standards for financial reporting
    - e-commerce: The software should support various payment gateways like PayPal, Stripe, and credit cards

## requirement engineering

- interviews: team elaborates questions for the stakeholders
  - formal
    - pyramid: detailed questions => generic questions
    - funnel: generic questions => detailed questions
    - diamond: detailed questions => generic questions => detailed questions
  - informal
- forms
- brainstorming: group problem-solving method for spontaneous contribution of creative ideas and solutions
- requirements workshop
- user stories: artifact that's a short description of a functionality from the user's perspective

example of user story:

```
User Story for Login Functionality:

As a registered user,
I want to log into the system using my email and password,
So that I can access my personal dashboard and manage my account.

Acceptance Criteria:
- The login form must validate email format and password length.
- The system should display an error message on invalid credentials.
- Successful login redirects to the user's dashboard.
```

example of user story:

```
User Story for Shopping Cart:

As a customer,
I want to add products to my shopping cart,
So that I can review and purchase them later.

Acceptance Criteria:
- The cart should show product name, quantity, price, and subtotal.
- The user can update quantities or remove items.
- The cart persists between sessions (e.g., via cookies or database).
```

## software architecture patterns

- cohesion: how much elements of the module are functionally related
- coupling: how dependent parts of a system are on each other

> [!TIP]
> ideal software architecture:
> High coHesion
> Low coupLing

### MVC (Model-View-Controller) architecture

> pattern in software design used to implement user interfaces, data and controlling logic

1. Model: manages data and business logic (e.g. postgresql)
1. View: handles layout and display (e.g. react.js)
1. Controller: acts as an intermediary between the model and view (e.g. express.js)

- model and view can communicate directly
- controller doesn't update view directly

### MVP (Model-View-Presenter) architecture

1. Model: manages data and business logic
1. View: handles layout and display
1. Presenter: connects Model and View

- view and model can't communicate directly
  - presenter always intermediates communication between model and view
- presenter updates view directly

> [!TIP]
> PRESENTer is always PRESENT between the communication of view and model

### MVVM (Model-View-ViewModel) architecture

1. Model: manages data and business logic
1. View: handles layout and display
1. ViewModel: acts as an intermediary between the view and the model

- data binding: automatic synchronization between view and viewmodel without adding code
  - this means that any change in viewmodel data also happens in view
- angular framework uses this architecture

- MVC x MVVM
  - controller x viewmodel as intermediaries
    - viewmodel uses data binding
    - MVC manually updates data between view and model

## design principles

- service contract: all services should have a contract
  - contract: document that describes what the service does
- service loose coupling: minimize coupling
  - ESB (Enterprise Service Bus) services: software architecture pattern that enables communication across applications
    - there can be ESB without SOA and SOA without ESB
- service abstraction: services should focus on important functionality and remove complexity
- service reusability: maximize service versatility and compatibility with different technologies
- service autonomy: each service controls its own logic and resources
- statelessness: services do not retain session state between requests, improving scalability
- service discoverability: make accessible and comprehensible services
- service composability: ability to combine services together to create new service or system

- service orchestration: centralized control system that coordinates interactions between multiple services
- service choreography: decentralized system
  - each service is autonomous and flexible

## SOA (Service-Oriented Architecture)

> principles and practices that makes software components interoperable and reusable

- service: software that provides access to resources through an interface
  - doesn't require knowledge of internal logic of the service to access resources
    - similar to services in real life

- each service provides a business capability
- services can also communicate with each other across platforms and languages
- use SOA to reuse services in different systems or combine several independent services to perform complex tasks
- uses the protocols: SOAP, restful http, etc

- advantages:
  - reusability
  - flexibility
  - maintainability
  - interoperability
  - loose coupling
  - governance
  - abstraction
- disadvantages:
  - performance depends of the server that hosts the service
  - limited scalability
  - single point of failure

### SOAP (Simple Object Access Protocol)

> protocol for exchanging structured information when implementing web services in XML format

- XML-based messaging protocol for exchanging information between computers
- components:
  - envelope(obligatory): defines start/end of message
  - header(optional): contains optional attributes
  - body(obligatory): contains actual message
    - fault(optional): carries error messages (contained in body)

- advantages:
  - platform and language independent
  - can work with different transport protocols (HTTP, SMTP, etc)
  - built-in error handling
  - language/platform/transport independent
  - works well with firewalls

- disadvantages:
  - verbose XML format
  - slower processing
  - more bandwidth required

### WSDL

> language that describes web services (written in XML)

- WSDL splits the description of a service in two perspectives:
  - abstract: describes service's interface and what it does
  - concrete: describes service's implementation

- why split service description in two types (abstract/concrete)?
  - if implementation (concrete) changes, interface (abstract) can still be available

- operation types: describes the behavior of the system making the request (like a client) and if it expects a response
  - one-way: operation receives request, but doesn't return an answer
  - request-response: operation receives request and will wait for an answer
  - solicit-response: operation sends request and will wait for a response
  - notification: operation sends request, but doesn't wait for a response

---

- governance: framework that defines processes, accountability, transparency and decision-making within an organization/system
- Minimum Viable Product (MVP): simplest version of a product that can be released to market
  - purpose:
    - validate assumptions about the product or market
    - test core functionalities
    - learn from real users before investing in full development
