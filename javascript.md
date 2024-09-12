
# basics

- multiplatform
- dynamically typed
- weakly typed
- garbage collection
- almost everything is an object
- first class functions
- supports multiple programming paradigms:
  - object oriented
  - functional
  - imperative
- interpreted, not compiled
- immutable strings

## data types

- String
- Number
- Boolean
- Null
- Undefined: variable that has been declared but not assigned a value
- Object
- Array

```javascript
let name = "Alice";     // String
let age = 25;           // Number
const isStudent = true; // Boolean
let job = null;         // Null
let address;            // Undefined
let person = { name: "Bob", age: 40 }; // Object
let fruits = ["apple", "banana", "cherry"]; // Array

console.log(fruits[0]); // Output: apple
fruits.push("orange");  // Adds "orange" to the end of the array
console.log(fruits);    // Output: ["apple", "banana", "cherry", "orange"]

```

## if else

```javascript

if (time < 10) {
  greeting = "Good morning";
} else if (time < 20) {
  greeting = "Good day";
} else {
  greeting = "Good evening";
}

```

## loops


```javascript

// For loop
for (let i = 0; i < 5; i++) {
  console.log(i);
}

// While loop
let i = 0;
while (i < 5) {
  console.log(i);
  i++;
}

```
## functions

```javascript

function greet(name) {
  return "Hello, " + name + "!";
}

const add = function(a, b) {
  return a + b;
};

const multiply = (a, b) => a * b;

console.log(greet("Alice")); // Output: Hello, Alice!
console.log(add(2, 3)); // Output: 5
console.log(multiply(4, 5)); // Output: 20

```

## objects

```javascript

let person = {
  firstName: "John",
  lastName: "Doe",
  age: 25,
  greet: function() {
    return "Hello, " + this.firstName;
  }
};

console.log(person.firstName); // Output: John
console.log(person.greet());   // Output: Hello, John

```


## array methods

- `push()`: adds an element to the end of an array
- `pop()`: removes the last element of an array
- `shift()`: removes the first element of an array and shifts all other elements down by one
- `unshift()`: adds an element to the beginning of an array and shifts all other elements up by one
- `slice()`: returns a copy of a portion of an array
  - doesn't mutate original array
- `splice()`: adds or removes elements from an array
  - mutates original array
- `sort()`: sorts the elements of an array
- `reverse()`: reverses the order of the elements in an array

```javascript
let arr = [1, 2, 3, 4, 5];
let sliced = arr.slice(1, 3);
console.log(sliced); // Outputs: [2, 3]
console.log(arr); // Outputs: [1, 2, 3, 4, 5]

let arr = [1, 2, 3, 4, 5];
let spliced = arr.splice(1, 2, 'a', 'b');
console.log(spliced); // Outputs: [2, 3]
console.log(arr); // Outputs: [1, 'a', 'b', 4, 5]
```

array methods that don't mutate the original array:
- `map()`: creates a new array by applying a function to each element of an existing array
- `filter()`: creates a new array with elements that pass a test specified by a function
- `find()`: returns the first element in an array that passes a test specified by a function
- `findIndexOf()`: returns the index of the first element in an array that passes a test specified by a function
- `reduce()`: applies a function to each element of an array to reduce the array to a single value
- `forEach()`: calls a function for each element in an array
- `some()`: checks if at least one element in an array passes a test specified by a function
- `every()`: checks if all elements in an array pass a test specified by a function

