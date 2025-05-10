# UML

> standardized visual modeling language that provides a general-purpose and intuitive way to visualize the design of a system

> helps in specifying, visualizing, constructing and documenting the artifacts of software systems

> helps communicate system architecture and behavior to stakeholders

- types
  - structure diagram
    - class diagram
    - object diagram
    - component diagram
    - profile diagram
    - composite structure diagram
    - package diagram
    - deployment diagram
  - behaviour diagram
    - activity diagram
    - use case diagram
    - state machine diagram
    - interaction diagram
      - sequence diagram
      - communication diagram
      - interaction overview diagram
      - timing diagram

## structure diagrams

### class diagram

[UML Class Diagram Reference 1](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/uml-class-diagram-tutorial/)

[UML Class Diagram Reference 2](https://learn.microsoft.com/previous-versions/visualstudio/visual-studio-2015/modeling/uml-class-diagrams-reference)

- main components:
  - entity
  - relationship
  - attribute

- cardinality symbols
  - how to read class diagrams:
    - Entity1 ----------------------> 0..* Entity2
    - Entity1 can have 0..* relations with Entity2
      - the multiplicity is closer to the entity it applies to
  - multiplicity
    - 1: one to one
    - *: many
    - 0..1: zero or one
    - 0..*: zero or many
    - 1..*: one or many
    - n..m: n to m

- attributes
  - public attributes `+`: accessible everywhere
  - private attributes `-`: accessible only by the same class
  - protected attributes `#`: accessible by the same class and its child classes
  - package/default attributes `~`: accessible by the same class and inside the package

- diamond: composition or aggregation
  - one entity contains another entity
  - the entity that is near the diamond is the container
  - the entity that is further from the diamond is the contained
  - types:
    - aggregation (empty diamond): one class contains another, but both exist independently
    - composition: (filled diamond): one class contains another and the contained object cannot exist independently

types of relationships:

| Relationship                     | Symbol    | Description                                                           |
| -------------------------------- | --------- | --------------------------------------------------------------------- |
| **Association**                  | ───────── | A general connection between classes                                  |
| **Reflexive Association**        | ───────── | A class has a relationship with itself                                |
| **Aggregation**                  | ────────◇ | A "has-a" relationship; the part can exist independently              |
| **Composition**                  | ────────◆ | A strong "part-of" relationship; the part is destroyed with the whole |
| **Inheritance** (Generalization) | ────────▶ | "Is-a" relationship; subclass inherits from superclass                |
| **Realization**                  | ╌╌╌╌╌╌╌╌▶ | Interface implementation                                              |
| **Dependency**                   | ╌╌╌╌╌╌╌╌> | One class uses another temporarily (e.g. as a method parameter)       |

## behavior diagrams

### use case diagram

- use cases: horizontally shaped ovals that represent the different actions that a user might do
- actors: stick figures that represent the people that will adopt the use cases
- associations
- system boundary boxes
- packages

- `<<includes>>`: indicates that a use case always uses the behavior of another use case
- `<<extends>>`: indicates that a use case optionally adds behavior to another use case
- Generalization (e.g. search book ────────▶ search): child use case inherits the behavior and meaning of the parent use case

### activity diagram

| Symbol                      | Name                            | Description                                                               |
| --------------------------- | ------------------------------- | ------------------------------------------------------------------------- |
| `Black Circle`           | **Initial Node**                | Entry point of the workflow. There is usually one per diagram.            |
| `Black Circle with Ring` | **Final Node (Activity Final)** | Represents the end of the entire activity flow.                           |
| `Rounded Rectangle`      | **Activity (Action) Node**      | Represents a task or action performed.                                    |
| `Arrow`                  | **Control Flow**                | Shows the direction of flow from one activity to the next.                |
| `Diamond`                | **Decision Node**               | Represents a branching point, usually with a condition (`yes`/`no`).      |
| `Diamond`                | **Merge Node**                  | Combines multiple incoming flows into one.                                |
| `Bar (Thick Line)`       | **Fork Node**                   | Splits a single flow into multiple concurrent flows (parallel execution). |
| `Bar (Thick Line)`        | **Join Node**                   | Synchronizes multiple flows into one (end of parallel execution).         |
| `Rectangle`              | **Swimlane (optional)**         | Divides responsibilities among actors or components.                      |

<!--| `Encircled X`             | **Flow Final Node**             | Ends one specific flow, but not the whole activity.                       |-->

