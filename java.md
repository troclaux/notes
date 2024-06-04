
# Class

example of Java Class
```java
// Java Program for class example

class Student {
	// data member (also instance variable)
	int id;
	// data member (also instance variable)
	String name;

	public static void main(String args[])
	{
		// creating an object of
		// Student
		Student s1 = new Student();
		System.out.println(s1.id);
		System.out.println(s1.name);
	}
}
```

## Java Data Types

### Primitive

> Primitive data types allow storing only one value at a particular location. They are predefined in the Java language and occupy a fixed amount of memory

- There are only 8 primitive data types in Java:
  - `int`, `double`, `boolean`, `char`, `float`, `long`, `short`, `byte`

### Object or Non-Primitive

> Unlike primitive data types, non-primitive ones are created by the users in Java

> Object data types can be used to store more complex data than primitive data types

- e.g. `String`, `Array`, `List`, `Set`, `Map` and `Class`

## Interface

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

## Keywords

- `private`: prevent value/method from being accessed outside the class
  - creates the need for public methods inside the class that provide controlled access to private fields
    - e.g.: getAccountNumeber(), getBalance(), deposit(), withdraw()
- `final`: function/value cannot be reassigned after it has been initialized
- `static`: variable/method/block belongs to the class itself, instead of instance
- `abstract`: abstract classes cannot be instantiated and is intended to be subclassed
  - can contain abstract methods (methods without body) that must be implemented by subclasses
