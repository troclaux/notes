
# functional programming

## prerequisite knowledge

- passing value as an argument: reference vs value
  - reference: function has access to original value and can modify it
  - value: function only has access to copy, changes to copy within function don't affect the original
- higher-order function: function that does at least one of the following:
  1. Takes one or more functions as arguments
  2. Returns a function as its result
- memoization: caching the result of a computation to avoid computing it again in the future
  - it's a tradeoff between memory and speed
- referential transparency: function can be replaced by its value without changing the program's behavior, without side effects

example of referential transparency:
```python

def add_nums(a, b):
    return a + b

```

## first-class functions

> functions that are treated as first-class citizens

- first-class citizens are entities that can be:
1. Created at runtime
2. Assigned to a variable or data structure
3. Passed as a parameter to a function
4. Returned as the result of a function


## pure functions

pure functions have the 2 following properties:
1. deterministic: always returns the same value given the same arguments
2. no side effects: the function does not alter any state or interact with the outside world
  - no modifying global variables
  - no input/output operations, etc
  - exceptions are considered side effects
    - in FP, represent error as data (e.g. the ParseError class), rather than raising exceptions

## function transformation

> modifying a function's input and/or output to create a new function

- higher-order function: a function that receives one or more functions as arguments or returns a function as a result

```python
def multiply(x, y):
    return x * y

def add(x, y):
    return x + y

# self_math is a higher order function
# input: a function that takes two arguments and returns a value
# output: a new function that takes one argument and returns a value

def self_math(math_func):
    def inner_func(x):
        return math_func(x, x)
    return inner_func

square_func = self_math(multiply)
double_func = self_math(add)

print(square_func(5)) # prints 25
print(double_func(5)) # prints 10
```

## anonymous/lambda functions

## recursion instead of loops

## closures

> functions that have access to the parent scope, even after the parent function has closed

properties of closures:
- can remember and access variables and arguments of its outer function even after that function has finished execution
- often used to keep state from function call to function call

- in python, use `nonlocal` keyword to access variables from parent scope
  - don't use `nonlocal` when you are modifying a variable that is immutable (lists, dictionaries or sets)
    - you can use `nonlocal` if you are reassigning the immutable variable

```javascript

function makeCounter() {
    let count = 0;
    return function() {
        return ++count;
    };
}

const counter = makeCounter();
console.log(counter()); // 1
console.log(counter()); // 2

```

> [!NOTE]
> in the code above, the inner function returned by makeCounter is the closure
> in the code below, the closure is `word_count`

```python

def word_count_aggregator():
    count = 0

    def word_count(doc):
        nonlocal count
        words = doc.split()
        count += len(words)
        return count

  return word_count

```

example of closure without using `nonlocal` keyword
```python

def new_collection(initial_docs):
    docs = initial_docs.copy()

    def add_doc(new_word):
        docs.append(new_word)
        return docs

    return add_doc

# Create a new collection
add_doc = new_collection(["doc1", "doc2", "doc3"])

# Test the closure
print(add_doc("doc4"))
# Output: ['doc1', 'doc2', 'doc3', 'doc4']

print(add_doc("doc5"))
# Output: ['doc1', 'doc2', 'doc3', 'doc4', 'doc5']
```

## currying

> convert single function that receives multiple arguments into multiple functions that each receives a single argument

```python

def add(x, y, z):
    return x + y + z

def curried_add(x):
    def add_y(y):
        def add_z(z):
            return x + y + z
        return add_z
    return add_y

# Using the curried function
result = curried_add(1)(2)(3)
print(result)  # Output: 6

# We can also partially apply the function
add_one = curried_add(1)
add_one_and_two = add_one(2)
final_result = add_one_and_two(3)
print(final_result)  # Output: 6

```

- why would I ever want to do this?
  - flexibility, by breaking down a multi-argument function into smaller, single-argument functions, you can:
    - Use the function in different contexts, such as with map, filter, or reduce
    - Create higher-order functions that take other functions as arguments
    - Easily compose functions together to create more complex operations


## decorators

- syntatic sugar for higher-order functions
  - higher-order function: function that does at least one of the following:
  1. Takes one or more functions as arguments
  2. Returns a function as its result

```python

def prefix(func_to_decorate):
    def wrapper():
        return "Hello " + func_to_decorate()

    return wrapper


# printer = prefix(printer)
@prefix
def printer():
    return "World"


print(printer())  # Output: Hello World

```
explanation:
- `def prefix(func):`: function that receives another function `func_to_decorate` as input
- `@prefix`: equivalent to `printer = prefix(printer)`
  - passes the `printer()` function as argument to `prefix()`
  - also renames `prefix(printer)` to `printer()`

## sum types

> type that can hold different kinds of values, but only one at a time

- also known as tagged unions or variant types

## python functional programming

- lambda functions: anonymous functions that can have any number of arguments, but only one expression
    - `lambda x: x + 1`
- nonlocal var: used in [closures](#closures) to access variables from parent scope
- map(): applies a function to each item in an iterable
  - `squared_numbers = list(map(lambda x: x**2, [1, 2, 3, 4, 5]))`
- filter(): 
- reduce(): 
- arguments that allow a function to accept a variable number of arguments
  - `*args`: collects positional arguments into a tuple
  - `**kwargs`: collects keyword/named arguments into a dictionary

```python

def print_arguments(*args, **kwargs):
    print(f"Positional arguments: {args}")
    print(f"Keyword arguments: {kwargs}")

print_arguments("hello", "world", a=1, b=2)
# Positional arguments: ('hello', 'world')
# Keyword arguments: {'a': 1, 'b': 2}

```
