
## properties

- all code in java must be written inside a class
- java uses a lot of memory because of java's virtual machine
- object oriented

- statically typed
- is considered compiled AND interpreted
- has garbage collection
- functions are values (i.e. first-class citizens)
- doesn't have hoisted declarations
- type casting can be done explicitly or implicitly
- strings aren't mutable
- multiplatform: can run on windows or linux or mac without changing the code

- JVM: allows java to run in any OS (Linux, Windows, Mac)
- JRE: contains JVM and all the extra tools and libraries needed to run a Java program


## data types

TODO

## basics

- `main()` method is required and you can only have one in a program
  - `public static void main(String[] args){}`
  - entry point of a Java program

```java

// you can import outside classes
import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        int x = 5;
        if (x > 10) {
            System.out.println("x is greater than 10");
        } else {
            System.out.println("x is less than or equal to 10");
        }
    }
}

```

```java
import java.util.ArrayList;
import java.util.Collections;

public class Main {
    public static void main(String[] args) {

        ArrayList<Integer> dynamicArray = new ArrayList<>();

        dynamicArray.add(30);
        dynamicArray.add(10);
        dynamicArray.add(20);

        Collections.sort(dynamicArray);  // Sorts the array in ascending order
        System.out.println(dynamicArray);  // Outputs: [10, 20, 30]

        int element = dynamicArray.get(0);  // Gets the first element (10)
        dynamicArray.remove(1);  // Removes the element at index 1 (20)

        int[] myArray = {1, 2, 3, 4, 5};
        String[] cars = {"Volvo", "BMW", "Ford", "Mazda"};

        System.out.println (cars.length);

        for (int i = 0; i < cars.length; i++) {
          System.out.println(cars[i]);
        }
    }
}

```

### type casting

2 types of casting:
- widening casting (implicit): converting a smaller type to a larger type size
  - smaller: means less memory, e.g. `int` is smaller than `double`
  - no data loss, because the smaller type fits into the larger type
- narrowing casting (explicit): converting a larger type to a smaller size type
  - larger: means more memory, e.g. `double` is larger than `int`
  - data loss, because the larger type does not fit into the smaller type

- casting vs parsing
  - casting: converting one data type to another
  - parsing: converting a string to a primitive data type

```java
public class Main {
    public static void main(String[] args) {

        // casting (widening and narrowing)

        int myInt = 9;
        double myDouble = myInt; // widening casting: int to double
        System.out.println(myInt); // Outputs 9
        System.out.println(myDouble); // Outputs 9.0

        double myDouble = 3.14;
        int myInt = (int) myDouble; // Narrowing casting: double to int

        // parsing

        String str = "2";
        int myInt = Integer.parseInt(str); // Converts the string '2' to the integer 2
        System.out.println("The integer value is: " + myInt);

        String str = "3.14";
        float myFloat = Float.parseFloat(str); // Converts the string '3.14' to the float 3.14
        System.out.println("The float value is: " + myFloat);

    }
}
```

## if else

```java
public class IfElseExample {
    public static void main(String[] args) {

        int age = 25;
        if (age >= 18) {
            System.out.println("You are an adult.");
        } else {
            System.out.println("You are a minor.");
        }
    }
}
```

## loops

```java
public class WhileLoopExample {
    public static void main(String[] args) {

        int i = 0;
        while (i < 5) {
            System.out.println("Hello, World!");
            i++;
        }
    }
}
```

```java
public class ForLoopExample {
    public static void main(String[] args) {
        for (int i = 0; i < 5; i++) {
            System.out.println("Hello, World!");
        }
    }
}
```

## class

example of Java Class:

```java
public class Student {
    private int id;
    private String name;

    // Constructor
    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    // Method to get the name
    public String getName() {
        return name;
    }

    // Method to set the name
    public void setName(String name) {
        this.name = name;
    }

    // Main method to test the class
    public static void main(String[] args) {
        Person person = new Person("John", 25);
    }
}
```

## data types

### primitive

> Primitive data types: allow storing only one value at a particular location. They are predefined in the Java language and occupy a fixed amount of memory

- There are only 8 primitive data types in Java:
  - `int`, `double`, `boolean`, `char`, `float`, `long`, `short`, `byte`

### object or non-primitive

> unlike primitive data types, non-primitive ones are created by the users in java

> object data types can be used to store more complex data than primitive data types

- e.g. `String`, `Array`, `List`, `Set`, `Map` and `Class`


## data structures

types of collections:
- list: ordered collection of elements
- set: collection of unique elements
- map: key-value pairs
- tree: sorted elements

- arrays
  - fixed size
  - fast access
- `ArrayList`
  - dynamic resizing
  - fast access
- `LinkedList`
  - fast insertion/deletion
  - slower access
- `Stack`: avoid using, use `ArrayDeque` instead
- `HashMap`
  - implements Map interface
- `HashSet`: stores unique elements with fast retrieval time
  - implements Set interface
- `TreeMap`: Red-Black tree based implementation of the Map interface
  - stores sorted key-value pairs
  - maintains the order of the keys
  - implements Map interface
- `TreeSet`: stores unique elements in sorted order
  - implements Set interface
- `PriorityQueue`: retrieves elements based on their priority
  - ONLY IN JAVA: min-heap based implementation (not a priority queue property)
  - elements are ordered based on their priority
  - highest priority element is removed first
  - time complexity (?)


examples
```java

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.TreeMap;
import java.util.TreeSet;
import java.util.PriorityQueue;
import java.util.Stack;

public class Main {
    public static void main(String[] args) {

        int[] arr = new int[5];  // Array of 5 integers
        arr[0] = 10;
        arr[1] = 20;
        System.out.println(arr[1]);  // Outputs: 20


        ArrayList<String> list = new ArrayList<>();
        list.add("Alice");
        list.add("Bob");
        list.remove("Alice");
        System.out.println(list.get(0));  // Outputs: Bob


        LinkedList<String> list = new LinkedList<>();
        list.add("Alice");
        list.addFirst("Bob");
        list.removeLast();
        System.out.println(list.getFirst());  // Outputs: Bob

        HashMap<String, Integer> fruit_counter = new HashMap<>();
        fruit_counter.put("Apple", 50);
        fruit_counter.put("Banana", 20);
        System.out.println(fruit_counter.get("Apple"));  // Outputs: 50

        HashSet<String> set = new HashSet<>();
        set.add("Alice");
        set.add("Bob");
        set.add("Alice");  // Duplicate, will not be added
        System.out.println(set.size());  // Outputs: 2

        TreeMap<String, Integer> map = new TreeMap<>();
        map.put("Charlie", 35);
        map.put("Alice", 25);
        map.put("Bob", 30);
        System.out.println(map.firstKey());  // Outputs: Alice

        TreeSet<String> set = new TreeSet<>();
        set.add("Charlie");
        set.add("Alice");
        set.add("Bob");
        System.out.println(set.first());  // Outputs: Alice

        PriorityQueue<Integer> pq = new PriorityQueue<>();
        pq.add(20);
        pq.add(15);
        pq.add(30);
        System.out.println(pq.poll());  // Outputs: 15 (smallest element)

        Stack<Integer> stack = new Stack<>();
        stack.push(10);
        stack.push(20);
        System.out.println(stack.pop());  // Outputs: 20
    }
}

```

## interface

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

another example:

```java
interface Observer {
    void onTemperatureChanged(float temperature);
    void onHumidityChanged(float humidity);
    void onWeatherUpdate(float temperature, float humidity);
    boolean isActive();
    String getObserverId();
}
```

## inheritance and encapsulation

- inheritance: allows one class to inherit the properties/methods of another class
- encapsulation: the concept of bundling data and methods that operate on data within a single class or object

## polymorphism

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

## exceptions

- [ ] TODO: learn java exceptions

## synchronism and multithreading

- [ ] TODO: learn java synchronism and multithreading

## packages

> mechanism to organize and group together files/directories/classes/interfaces in a logical and reusable manner

to build a package, you must provide:
- package objects (files and directories)
- two information files (pkginfo and prototype files)
- optional information files
- optional installation scripts

## Keywords

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

## compilation

| .java | => compilation => | .class (bytecode) |

java program development:
1. write java program
2. compilation
3. load program to primary memory
4. bytecode verification
5. execute program

---

## applet

> small application that is designed to be transmitted over the internet and executed by a web browser

5 methods that define the life cycle of an applet:
- `init()`: initializes the applet
- `start()`: called after init, starts the applet
- `stop()`: stops playing animation/audio/video
- `paint()`: draws something on the applet
- `destroy()`: free resources, called when the browser is closed
