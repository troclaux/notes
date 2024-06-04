# UML

[UML Class Diagram Reference 1](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/uml-class-diagram-tutorial/)

[UML Class Diagram Reference 2](https://learn.microsoft.com/en-us/previous-versions/visualstudio/visual-studio-2015/modeling/uml-class-diagrams-reference?view=vs-2015)

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
