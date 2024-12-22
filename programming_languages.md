
- abstraction: simplification of complex operations by focusing on essential features while hiding unnecessary details
  - examples:
    - interfaces that define what an object can do without specifying how
    - functions that hide complex implementations behind simple names
    - classes that represent real-world concepts in simplified form
- encapsulation: bundling data and methods that operate on that data within a single unit, restricting direct access to some of an object's components
  - hides internal details and provides an interface to interact with the object
  - examples:
    - classes that bundle data (fields) and behavior (methods)
    - modules that group related functions and variables
    - objects that protect their internal state with getters/setters

- primitives: fundamental data types that are not composed of other data types
  - e.g. integers, floating-point numbers, characters, booleans, strings, null or undefined values

- variable: the name given to a memory location
- field: data encapsulated within a class or object

- function: set of instructions that perform a task
- method: set of instructions associated with an object

- function signature: defines input and output of a function

- pass by value: actual value of the argument is passed
- pass by reference: address to the actual argument is passed

- namespace: declarative region that provides a scope to the identifiers (names of types, functions, variables) inside it
- anonymous function: function that is defined without name

- high-order functions: function that takes another function as an argument
- first-class citizen: objects that can be:
  - created at runtime
  - passed as arguments to other functions
  - returned as values from other functions
  - stored in data structures
  - used as values in expressions

- concurrency: ability of a program to manage multiple tasks at the same time
  - can happen on a single-core system with task switching
- parallelism: ability to execute multiple tasks at the same time
  - requires a multi-core system where multiple tasks can be executed simultaneously
- asynchronous: starting a task that runs on the background without blocking the main program flow
  - means that the program performs non-blocking tasks, usually involving I/O operations

- prototype-based approach to object-oriented programming:
  - style where objects are created by cloning existing objects (prototypes) rather than instantiating classes
    - objects can directly inherit from other objects
    - there are no classes; instead, objects serve as prototypes for other objects
    - new objects can be created by copying an existing object and then modifying it
    - this allows for more flexibility and dynamic behavior compared to classical class-based oop
