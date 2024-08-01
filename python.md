
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

- immutable strings

## basic methods

- len(obj1): returns the number of elements in an object (string, list, set, etc)
- min(iter1): returns the minimum value of an iterable
- max(iter1): returns the maximum value of an iterable
- sum(iter1): returns the sum of all values in an iterable
- str1.upper(): return str1 in upper case
- str1.lower(): return str1 in lower case
- str1.strip(): remove leading/trailing spaces and line breaks from input string
- str1.replace(str2): if str2 is in str1, replace it
- str1.split(): split string into a list of string
  - by default, the delimiter is space
  - provide string/char as input to change default delimiter
- str1.count(str2): return the number of times str2 occurs in str1
- str1.isalnum(): returns True if all chars in string are alphanumeric, False otherwise
- str1.isdigit(): returns True if all chars in string are numbers, False otherwise
- list1.append(elem1): appends an element to the end of a list
- list1.copy(): returns a copy of a list
- list1.insert(idx, elem1): inserts elem1 at index idx of the list
- list1.push(elem1): inserts elem1 at the end of list1
- list1.pop(): removes and returns last element of list1
- list1.clear(): removes all elements of list1
- list1.reverse(): reverse the order of elements in list1
- list1.sort(): sort elements in list1


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

type casting:

```python
# converts float number into int
int_num = int(4.0)

# converts int number into float
float_num = float(4)

# converts number into string
str_num = (5)
```
