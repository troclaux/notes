# UML

> Standardized visual modeling language that provides a general-purpose and intuitive way to visualize the design of a system.

> UML helps in specifying, visualizing, constructing, and documenting the artifacts of software systems.

- UML Diagrams:
  - Behaviour Diagram
    - Activity Diagram
    - Use Case Diagram
    - State Machine Diagram
    - Interaction Diagram
      - Sequence Diagram
      - Communication Diagram
      - Interaction Overview Diagram
      - Timing Diagram
  - Structure Diagram
    - Class Diagram
    - Object Diagram
    - Component Diagram
    - Profile Diagram
    - Deployment Diagram
    - Composite Structure Diagram

## Class diagram

[UML Class Diagram Reference 1](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/uml-class-diagram-tutorial/)

[UML Class Diagram Reference 2](https://learn.microsoft.com/en-us/previous-versions/visualstudio/visual-studio-2015/modeling/uml-class-diagrams-reference?view=vs-2015)

Main components
- Entity
- Relationship
- attribute

- cardinality symbols
  - read diagram as:
    - Entity1 has can have 0..* Entity2
    - the multiplicity is closer to the entity it applies to
  - multiplicity
    - 1: one
    - *: many
    - 0..1: zero or one
    - 0..*: zero or many
    - 1..*: one or many
    - n..m: n to m

- attributes
  - public attributes are identified by +
  - private attributes are identified by -
  - protected attributes are identified by #

- arrow
- diamond: association or aggregation
  - one entity contains another entity
  - the entity that has a diamond is the container
  - the entity that is further from the diamond is the contained
  - types:
    - empty diamond: aggregation
      - the entity that is contained in another entity can exist separately from the entity that contains it
    - filled diamond: composition
      - the entity that is contained in another entity cannot exist separately from the entity that contains it
