- pass by value: actual value of the argument is passed
- pass by reference: address to the actual argument is passed

- namespace: declarative region that provides a scope to the identifiers (names of types, functions, variables )inside it

# introduction

- object oriented
- multiplatform: can run on windows or linux or mac without changing the code
- concurrent
- distributed
- portable
- safe
- extensible
- vast set of libraries and APIs
- garbage collector
- statically typed

- JVM: allows java to run in any OS (Linux, Windows, Mac)
- JRE: contains JVM and all the extra tools and libraries needed to run a Java program


## basic operations


```java
import java.util.ArrayList;

ArrayList<Integer> dynamicArray = new ArrayList<>();

dynamicArray.add(10);
dynamicArray.add(20);
dynamicArray.add(30);

int element = dynamicArray.get(0);  // Gets the first element (10)
dynamicArray.remove(1);  // Removes the element at index 1 (20)

int[] myArray = {1, 2, 3, 4, 5};
String[] cars = {"Volvo", "BMW", "Ford", "Mazda"};

System.out.println (cars.length);

for (int i = 0; i < cars.length; i++) {
  System.out.println(cars[i]);
}
```

type casting:

```java
public class Main {
  public static void main(String[] args) {

    int myInt = 9;
    double myDouble = myInt; // Automatic casting: int to double
    System.out.println(myInt); // Outputs 9
    System.out.println(myDouble); // Outputs 9.0

    String str = "2";
    int myInt = Integer.parseInt(str); // Converts the string '2' to the integer 2
    System.out.println("The integer value is: " + myInt);

    String str = "3.14";
    float myFloat = Float.parseFloat(str); // Converts the string '3.14' to the float 3.14
    System.out.println("The float value is: " + myFloat);

  }
}
```

# Class

- compilation: process of translating source code into machine code or intermediary form

| .java | => compilation => | .class (bytecode) |

java program development:
1. write java program
2. compilation
3. load program to primary memory
4. bytecode verification
5. execute program

example of Java Class

```java
class Student {
  int id;
  String name;

  public static void main(String args[]) {
    // creating an object of Student
    Student s1 = new Student();
    System.out.println(s1.id);
    System.out.println(s1.name);
  }
}
```

## Java Data Types

### Primitive

> Primitive data types: allow storing only one value at a particular location. They are predefined in the Java language and occupy a fixed amount of memory

- There are only 8 primitive data types in Java:
  - `int`, `double`, `boolean`, `char`, `float`, `long`, `short`, `byte`

### Object or Non-Primitive

> Unlike primitive data types, non-primitive ones are created by the users in Java

> Object data types can be used to store more complex data than primitive data types

- e.g. `String`, `Array`, `List`, `Set`, `Map` and `Class`

# Interface

- a contract that defines a set of methods that a class must implement

example:

```java
interface Gadget {
  void turnOn();
  void turnOff();
  void charge();
}
```

implementing the interface:

```java
class Smartphone implements Gadget {
  @Override
  public void turnOn() {
    System.out.println("Smartphone is turning on.");
  }

  @Override
  public void turnOff() {
    System.out.println("Smartphone is turning off.");
  }

  @Override
  public void charge() {
    System.out.println("Smartphone is charging.");
  }
}
```
# inheritance and encapsulation

- inheritance: allows one class to inherit the properties/methods of another class
- encapsulation: the concept of bundling data and methods that operate on data within a single class or object

# polymorphism

> allows one method to do different things based on the object it is acting upon

example of polymorphism with method overloading:

```java
class Example {
  void display(int a) {
    System.out.println("Argument: " + a);
  }
  void display(String a) {
    System.out.println("Argument: " + a);
  }
}

public class Main {
  public static void main(String[] args) {
    Example obj = new Example();
    obj.display(10);          // Calls display(int a)
    obj.display("Hello");     // Calls display(String a)
  }
}
```

# exceptions

- [ ] TODO: learn java exceptions

# synchronism and multithreading

- [ ] TODO: learn java synchronism and multithreading

# packages

> mechanism to organize and group together files/directories/classes/interfaces in a logical and reusable manner

to build a package, you must provide:
- package objects (files and directories)
- two information files (pkginfo and prototype files)
- optional information files
- optional installation scripts

# Keywords

- `abstract`: abstract classes cannot be instantiated and are intended to be subclassed
  - can contain abstract methods (methods without body) that must be implemented by subclasses
  - can be accessed without creating an object
- `extends`: indicates a child class inherits properties of a parent class
- `final`: variable/method/class cannot be changed after it has been initialized
- `implements`: implements an interface
  - all the methods of the interface are implicitly public abstract
- `interface`: class with only abstract methods
- `static`: variable/method/block belongs to the class itself, instead of instance

- `public`: class contents (methods, attributes, constructor) can be accessed anywhere
- `protected`: can be accessed from within the same package and subclasses
- `private`: can only by accessed from same class
  - creates the need for public methods inside the class that provide controlled access to private fields
    - e.g.: getAccountNumber(), getBalance(), deposit(), withdraw()

---

# applet

> small application that is designed to be transmitted over the internet and executed by a web browser

5 methods that define the life cycle of an applet:
- init(): initializes the applet
- start(): called after init, starts the applet
- stop(): stops playing animation/audio/video
- paint(): draws something on the applet
- destroy(): free resources, called when the browser is closed
