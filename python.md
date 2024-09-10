
# python

## basics

- easy to read and understand
- multiplatform
- blocks are defined by indentation
- garbage collection
- object oriented
  - everything is an object
- can also be used for functional programming
- huge variety of libraries

- parenthesis in `if` statements are optional
- immutable strings

## basic methods

- `len(obj1)`: returns the number of elements in an object (string, list, set, etc)
- `min(iter1)`: returns the minimum value of an iterable
- `max(iter1)`: returns the maximum value of an iterable
- `sum(iter1)`: returns the sum of all values in an iterable

- `str1.upper()`: return str1 in upper case
- `str1.lower()`: return str1 in lower case
- `str1.strip()`: remove leading/trailing spaces and line breaks from input string
- `str1.replace(str2)`: if str2 is in str1, replace it
- `str1.split()`: split string into a list of strings
  - by default, the delimiter is space
  - provide string/char as input to change default delimiter
- `str1.count(str2)`: return the number of times str2 occurs in str1
- `str1.isalnum()`: returns True if all chars in string are alphanumeric, False otherwise
- `str1.isdigit()`: returns True if all chars in string are numbers, False otherwise
- `list1.append(elem1)`: appends an element to the end of a list
- `list1.copy()`: returns a copy of a list
- `list1.insert(idx, elem1)`: inserts elem1 at index idx of the list
- `list1.push(elem1)`: inserts elem1 at the end of list1
- `list1.pop()`: removes and returns last element of list1
  - if you provide a number as input, `pop()` will try to remove the number at that index
- `list1.remove(elem1)`: first instance of elem1, if there's no instance of elem1 raises `ValueError`
  - removes element by value
- `del list1[1:3]`: removes element by index
- `list1.clear()`: removes all elements of list1
- `list1.reverse()`: reverse the order of elements in list1
- `list1.sort()`: sort elements in list1
- `dict1.clear()`: removes all elements in dict1
- `dict1.copy()`: returns copy of dict1
- `dict1.keys()`: returns a list with all keys in dict1
- `dict1.values()`: returns a list with all values in dict1
- `dict1.items()`: returns a list containing a tuple for each key-value pair

```python
fruits = ["apple", "banana", "cherry", "date"]
del fruits[1:3]  # Removes "banana" and "cherry" (index 1 to 2)
print(fruits)    # Output: ["apple", "date"]
```

```python
s = "Hello, World!"

# Convert to uppercase
upper_s = s.upper()
print(upper_s)  # Outputs: "HELLO, WORLD!"

# Convert to lowercase
lower_s = s.lower()
print(lower_s)  # Outputs: "hello, world!"
```

```python
s = "Hello, World!"
s = s.replace("World", "Python")
print(s)  # Outputs: "Hello, Python!"
```

```python
s1 = "Hello123"
s2 = "Hello, World!"

print(s1.isalnum())  # Outputs: True
print(s2.isalnum())  # Outputs: False
```

## functions

```python
def my_function(input1, input2):
  print(input1, input2)

my_function('hello', 'world')
```

## enums

> variable that can only take a set of fixed values

```python

from enum import Enum

Color = Enum('Color', ['RED', 'GREEN', 'BLUE'])
print(Color.RED)  # this works, prints 'Color.RED'
print(Color.TEAL) # this raises an exception

```

## lambda functions

> basically small functions that are declared with the `lambda` keyword

lambda input1: operation which result will be returned

```python
add_nums = lambda x, y: x + y
print(add_nums(2, 3)) # 5
```

```python
numbers = [1, 2, 3, 4, 5]
squared = map(lambda x: x ** 2, numbers)
print(list(squared))  # Outputs: [1, 4, 9, 16, 25]
```


## data types

- text
  - str
- numeric
  - int
  - float
  - complex: for complex numbers
- sequence
  - list
  - tuple
  - range
- map
  - dict
- set
  - set
  - frozenset
    - immutable set
- boolean
  - bool
    - True
    - False
- binary
  - bytes, bytearray, memoryview
- None

### initialization of data types

```python
age = 23
name = 'Arthur'
tuple1 = ("apple", "banana", "orange", "apple")
tuple2 = (1, 5, 7, 9, 3)
tuple3 = (True, False, False)
tuple4 = ("abc", 34, True, 40, "male")
my_set = {1, 2, 3}
my_set.add(4)  # Now my_set is {1, 2, 3, 4}
my_frozenset = frozenset([1, 2, 3])
```

```python
# Initialize a dictionary
person = {
    "name": "John",
    "age": 30,
    "city": "New York"
}

# Accessing elements
print(person["name"])  # Outputs: "John"

# Adding a new key-value pair
person["profession"] = "Engineer"

# Updating a value
person["city"] = "San Francisco"

# Removing a key-value pair
del person["age"]

# Iterating over a dictionary
for key, value in person.items():
    print(f"{key}: {value}")

# Checking if a key exists
if "name" in person:
    print("Name is present in the dictionary")

# Get the number of key-value pairs
print(len(person))

# Removing all items
person.clear()
```

### type checking
```python

elem = 5.1

# checks if variable is a number
isinstance(elem, (int, float)) # returns True
isinstance(5, list) # returns False
isinstance([5, 6], list) # returns True

```

### type casting

```python
# converts float number into int
int_num = int(4.0)

# converts int number into float
float_num = float(4)

# converts number into string
str_num = str(5)
```
### dates

```python
import datetime
x = datetime.datetime.now()
```

## classes/objects

```python

# Define a class
class Dog:
    def __init__(self, name):
        self.name = name

    def bark(self):
        return "Woof!"

# Create an object of the class Dog
my_dog = Dog("Fido")

# Access the object's attribute
print(my_dog.name)  # Outputs: "Fido"

# Call the object's method
print(my_dog.bark())  # Outputs: "Woof!"
```

## inheritance

```python

class ParentClass:
    def __init__(self, name):
        self.name = name

    def greet(self):
        print(f"Hello, {self.name}!")

# ChildClass inherits from ParentClass
class ChildClass(ParentClass):
    def __init__(self, name, age):
        # Call the constructor of ParentClass
        super().__init__(name)
        self.age = age

    def display_age(self):
        print(f"{self.name} is {self.age} years old.")

# Creating an object of the ChildClass
child = ChildClass("Alice", 20)
child.greet()         # Method inherited from ParentClass
child.display_age()   # Method of ChildClass

```

## iterators

> object that contains a countable amount of values


## modules

> modules are individual python files (.py) that can be imported and reused in other python files.

```python
import my_module as mm
import pandas as pd

x = dir(pd) # list the names of all functions and variables in a module
```

## packages

> groups related modules into a directory structure

- to make a directory a package, it must contain a special file called __init__.py
- the file can be empty, but its presence makes the directory importable
- you can also include initialization code in __init__.py, which will run when the package is imported

```
my_project/
├── my_package/
│   ├── __init__.py
│   ├── module_a.py
│   ├── module_b.py
│
├── tests/
│   ├── __init__.py       # Makes tests a package
│   ├── test_module_a.py  # Test file for module_a
│   └── test_module_b.py  # Test file for module_b
│
├── requirements.txt      # Project dependencies
└── main.py
```

to import a module from a package, use the following syntax:
```python
import math as m        # creates alias for module
from module import *
from my_package.module_a import *
from my_package.module_b import Node
```

## libraries

> collection of packages and modules that provide specific functionality

- can be built-in (come with Python) or third-party (need to be installed, normally using pip)
  - standard library
    - included with python
    - no need to install
    - e.g.: `os`, `sys`, `math`, `random`, `datetime`, etc
  - third-party libraries
    - not included with python
    - need to be installed
    - e.g.: `numpy`, `pandas`, `matplotlib`, `requests`, etc

standard library example
```python
import math

print(math.sqrt(16))  # Outputs: 4.0
```

## automation with scripts

- PyAutoGUI: allows scripts to control mouse and keyboard to automate interaction with apps
- pywinauto: automates GUI testing for windows applications
- selenium: automates web browser interaction for testing web applications
- splinter: simplifies web application testing
- scrapy: web crawling framework for extracting data from websites
- windmill: tests web applications with automation and debugging tools
  - OBS: Windmill => Web
- pytest: testing framework
- ReportLab: creates pdf reports programmatically
- PDFMiner: extracts texts, images, and metadata from pdfs
- Borb: library for creating and manipulating pdfs
- OpenPyXL: reads and writes excel files
- PyXLL: allows you to create excel functions with python
- XlsxWriter: creates excel files
- Tagui: 
- Robot Framework: 
