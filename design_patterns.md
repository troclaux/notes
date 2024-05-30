# Design Patterns

[catalog of design patterns](https://refactoring.guru/design-patterns/catalog)

> reusable solutions to common problems in software design

- provide proven solutions
- improve code readability and maintainability
- facilitate communication
- encourage best practices
- speed up development process

## Creational patterns

> Object creation mechanisms

### Abstract Factory

> Interface to create related or dependent objects without specifying their concrete classes

### Singleton

> garantees that a class has only 1 instance

Implementation:

- Make default constructor private, to prevent new instances from being created
- create a static method that returns the same instance of its own class
  - getInstance should be the only way of getting Singleton object

### Builder

> allows you to produce different types and representations of an object using the same construction code

Implementation:

- Abstract the object constructions code out of its own class

### Prototype

> reuse existing objects without making code dependent on their classes

- an object that supports cloning is called a prototype

Implementation:
- Add a clone() method that carries over all of the field values of the old object into the new one

## Structural patterns

### Adapter

> allows objects with incompatible interfaces to collaborate

- e.g. 

### Bridge

### Composite

### Decorator

### Facade

### Flyweight

### Proxy

## Behavioral patterns

> Focus on algorithms and assignment of responsibilities between objects

### Chain of Responsibility

### Command

### Iterator

### Mediator

### Memento

### Observer

### State

### Strategy

### Visitor
