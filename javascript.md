
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
- hoisted declarations: variables and functions can be used before they are declared

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

let a, b, rest;
// Example of destructuring assignment with arrays (can also be used with objects)
[a, b] = [10, 20];

console.log(a); // Expected output: 10
console.log(b); // Expected output: 20

[a, b, ...rest] = [10, 20, 30, 40, 50];

console.log(rest); // Expected output: Array [30, 40, 50]

```

- `let`: variable that can be reassigned
  - block-scoped
- `const`: variable that cannot be reassigned
  - block-scoped
- `var`: variable that can be reassigned and redeclared
  - avoid using `var`
  - function-scoped

### type casting

- `Number()`: converts a value to a number
  - if the value cannot be converted, it returns `NaN`
  - `const num = Number("123");`
- `parseInt()`: converts a string to an integer
  - `const num = parseInt("123.45");`
- `parseFloat()`: converts a string to a floating-point number
  - `const num = parseFloat("123.45");`
- `String()`: converts a value to a string
  - `const text = String(123);`
- `toString()`: also converts a number to a string
  - `const text = (123).toString();`
- `Boolean()`: converts a value to a boolean
  - `const bool = Boolean(0);`

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

### ternary operator

> `condition ? return_this_value_if_true : return_this_value_if_false`

```javascript
let x = 5;
let y = (x < 10) ? "x is less than 10" : "x is greater than or equal to 10";
console.log(y); // Output: "x is less than or equal to 10"
```

### switch case

```javascript
let day = 3;

switch (day) {
  case 1:
    console.log("Monday");
    break;
  case 2:
    console.log("Tuesday");
    break;
  case 3:
    console.log("Wednesday");
    break;
  default:
    console.log("Invalid day");
}
```


## loops

```javascript

// for loop
for (let i = 0; i < 5; i++) {
  console.log(i);
}

// while loop
let i = 0;
while (i < 5) {
  console.log(i);
  i++;
}

// do while loop
let i = 0;
do {
  console.log(i);
  i++;
} while (i < 5);


// for of loop
let fruits = ['apple', 'banana', 'cherry'];
for (let fruit of fruits) {
  console.log(fruit);
}

```
## functions

types of function initialization:
```javascript

// function declaration
function greet(name) {
  return "Hello, " + name + "!";
}

// function expression
const add = function(a, b) {
  return a + b;
};

// arrow function
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

## object oriented programming

- `class`: used to define a new class
- `constructor()`: special method for creating and initializing an object created with a class
- `this`: used to call the constructor of a parent class
- `new`: creates a new instance of a class
- `extends`: used to create a child class that inherits from a parent class
- `static`: defines a static method for a class
- `super()`: calls the constructor of a parent class

- methods don't need the `function` keyword

```javascript

class Person {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }

  greet() {
    return "Hello, " + this.name;
  }
}

const myCar1 = new Car("Ford", 2014);
const myCar2 = new Car("Audi", 2019);

```

### JSON (JavaScript Object Notation)

- lightweight data interchange format
- contains name/value pairs
- values can be strings, numbers, objects, arrays, booleans, or null
- can be parsed and converted to JavaScript objects

```javascript
my_object = {
  "name": "John",
  "age": 25,
  "isStudent": true,
  "job": null,
  "address": undefined,
  "father": { "name": "Bob", "age": 40 },
  "fruits": ["apple", "banana", "cherry"]
};
```

- `JSON.stringify()`: converts a JavaScript object to a JSON string
  - `const text = JSON.stringify(obj);`
- `JSON.parse()`: converts a JSON string to a JavaScript object
  - `const obj = JSON.parse(text);`

## import/export

- `import`: used to import functions, objects, or primitive values from a file
- `export`: used to export functions, objects, or primitive values from a file


```javascript

// mathFunctions.js
export function add(x, y) {
  return x + y;
}

export function subtract(x, y) {
  return x - y;
}

```

in another file, you can import these named exports:

```javascript

// main.js
import { add, subtract } from './mathFunctions.js';

console.log(add(5, 3));  // Output: 8
console.log(subtract(5, 3));  // Output: 2

```

## error handling

- `try`: defines a block of code to try
- `catch`: defines a block of code to execute if an error occurs in the try block
- `finally`: optional block of code that runs after try and catch, regardless of the outcome

```javascript

try {
  let x = 10 / 0;
  console.log(x);
} catch (error) {
  console.log('An error occurred');
} finally {
  console.log('This runs no matter what');
}

```

## asynchronous operations

- `async`: used to declare an asynchronous function
- `await`: pauses the execution of the function until the promise is resolved

```javascript

async function fetchData() {
  const response = await fetch('https://api.example.com/data');
  const data = await response.json();
  console.log(data);
}

```

# keywords

- `default`: used to export a single value from a file
