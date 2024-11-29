# SOLID Principles

> set of five design guidelines that help developers create more maintainable, flexible, and scalable software

- [Single responsibility principle](#single-responsibility-principle)
- [Open-closed principle](#open-closed-principle)
- [Liskov substitution principle](#liskov-substitution-principle)
- [Interface segregation principle](#interface-segregation-principle)
- [Dependency inversion principle](#dependency-inversion-principle)

## Single Responsibility Principle

> a class should have only one reason to change, meaning it should only have one responsibility or job

- a class should have only one job
- if a class has more than one job, it should be split into multiple classes

bad example:

```python
class User:
    def __init__(self, name, email):
        self.name = name
        self.email = email

    def save_to_database(self):
        # Code to save user to database
        pass

    def send_email(self, message):
        # Code to send email to user
        pass
```

good example:

```python
class User:
    def __init__(self, name, email):
        self.name = name
        self.email = email

class UserRepository:
    def save_to_database(self, user):
        # Code to save user to database
        pass

class EmailService:
    def send_email(self, user, message):
        # Code to send email to user
        pass
```

## Open-Closed Principle

> software entities (classes, functions, modules) should be open for extension but closed for modification

- you should be able to add new functionality without changing existing code
- this reduces the risk of introducing bugs into existing, tested code when adding new features

bad example:

```python
class Rectangle:
    def __init__(self, width, height):
        self.width = width
        self.height = height

class AreaCalculator:
    def calculate_area(self, shapes):
        total_area = 0
        for shape in shapes:
            if isinstance(shape, Rectangle):
                total_area += shape.width * shape.height
        return total_area
```

good example:

```python
from abc import ABC, abstractmethod
import math

class Shape(ABC):
    @abstractmethod
    def area(self):
        pass

class Rectangle(Shape):
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def area(self):
        return self.width * self.height

class Circle(Shape):
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return math.pi * self.radius ** 2

class AreaCalculator:
    def calculate_area(self, shapes):
        return sum(shape.area() for shape in shapes)
```

## Liskov Substitution Principle

> objects of a superclass should be replaceable with object of a subclasse without affecting the correctness of the program

- bad example: `Ostrich` cannot fly, but it inherits from `Bird` which has a fly method
- this violates Liskov Substitution Principle because replacing a `Bird` with an `Ostrich` changes expected behavior

```python
class Bird:
    def fly(self):
        pass

class Sparrow(Bird):
    def fly(self):
        print("Sparrow flying")

class Ostrich(Bird):
    def fly(self):
        raise NotImplementedError("Ostriches can't fly")
```

- good example: by separating flying and non-flying birds, `Ostrich` no longer inherits from a class that expects it to fly, adhering to LSP

```python
from abc import ABC, abstractmethod

class Bird(ABC):
    @abstractmethod
    def move(self):
        pass

class FlyingBird(Bird):
    @abstractmethod
    def fly(self):
        pass

class Sparrow(FlyingBird):
    def move(self):
        self.fly()

    def fly(self):
        print("Sparrow flying")

class Ostrich(Bird):
    def move(self):
        print("Ostrich running")
```

## Interface Segregation Principle

> a client should never be forced to implement an interface that it doesn’t use
> clients should not be forced (?) to depend on interfaces they do not use

- a class should not be forced to implement interfaces it doesn't use

bad example: `SimplePrinter` is forced to implement `scan_document`, which it doesn't support

```python
class Printer(ABC):
    @abstractmethod
    def print_document(self, document):
        pass

    @abstractmethod
    def scan_document(self, document):
        pass

class SimplePrinter(Printer):
    def print_document(self, document):
        print("Printing document")

    def scan_document(self, document):
        raise NotImplementedError("SimplePrinter cannot scan")
```

good example: by splitting the interfaces into `PrinterInterface` and `ScannerInterface`, `SimplePrinter` only implements what it needs, adhering to ISP

```python
from abc import ABC, abstractmethod

class PrinterInterface(ABC):
    @abstractmethod
    def print_document(self, document):
        pass

class ScannerInterface(ABC):
    @abstractmethod
    def scan_document(self, document):
        pass

class SimplePrinter(PrinterInterface):
    def print_document(self, document):
        print("Printing document")

class MultiFunctionPrinter(PrinterInterface, ScannerInterface):
    def print_document(self, document):
        print("Printing document")

    def scan_document(self, document):
        print("Scanning document")
```

## Dependency Inversion Principle

> high-level modules should not depend on low-level modules. both should depend on abstractions

- abstractions should not depend on details. Details should depend on abstractions
- high-level modules should not depend on low-level modules. Both should depend on abstractions

bad example: `Switch` directly depends on the concrete `LightBulb` class. Changing `LightBulb` or using a different bulb type would require modifying `Switch`

```python
class LightBulb:
    def turn_on(self):
        print("LightBulb on")

    def turn_off(self):
        print("LightBulb off")

class Switch:
    def __init__(self, bulb: LightBulb):
        self.bulb = bulb

    def operate(self, action: str):
        if action == "on":
            self.bulb.turn_on()
        elif action == "off":
            self.bulb.turn_off()
```

good example: `Switch` now depends on the `Switchable` abstraction. You can easily swap `LightBulb` with `Fan` or any other device that implements `Switchable` without changing the `Switch` class, adhering to DIP

```python
from abc import ABC, abstractmethod

class Switchable(ABC):
    @abstractmethod
    def turn_on(self):
        pass

    @abstractmethod
    def turn_off(self):
        pass

class LightBulb(Switchable):
    def turn_on(self):
        print("LightBulb on")

    def turn_off(self):
        print("LightBulb off")

class Fan(Switchable):
    def turn_on(self):
        print("Fan on")

    def turn_off(self):
        print("Fan off")

class Switch:
    def __init__(self, device: Switchable):
        self.device = device

    def operate(self, action: str):
        if action == "on":
            self.device.turn_on()
        elif action == "off":
            self.device.turn_off()
```
