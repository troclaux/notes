
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

- private: prevent value/method from being accessed outside the class
  - creates the need for public methods inside the class that provide controlled access to private fields
    - e.g.: getAccountNumeber(), getBalance(), deposit(), withdraw()
- final: make function/value immutable
- static: variable/method/block belongs to the class itself, instead of instance
- abstract: abstract classes cannot be instantiated and is intended to be subclassed
  - can contain abstract methods (methods without body) that must be implemented by subclasses
