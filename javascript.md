
# JavaScript

## properties

- almost everything is an object

- dynamically typed
- weakly typed language
- interpreted programming language
- has garbage collection
- functions are values (i.e. first-class citizens)
- function-scoped: once a variable is declared inside a function, it is accessible anywhere inside that function
- hoisted declarations: variables and functions can be used before they are declared
- immutable strings
- single-threaded: can only do one thing at the time
  - v8 engine + browser api is multi-threaded
  - node is single-threaded

## data types

- primitive types
  - `number`: represents both integers and floating-point numbers
  - `string`
  - `boolean`: `true` or `false`
  - `null`
  - `undefined`: variable that has been declared but not assigned a value
    - arrays are Object
  - `symbol`
  - `bigint`
- reference types
  - `object`: collections of related data and/or functionality
    - arrays and functions are objects in javascript

```javascript
let name = "Alice";     // String
let age = 25;           // Number
const isStudent = true; // Boolean
let job = null;         // Null
let address;            // Undefined
let person = { name: "Bob", age: 40 }; // Object
let fruits = ["apple", "banana", "cherry"]; // Array

console.log(typeof fruits);            // Output: Object
console.log(person instanceof Object); // Output: true

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

const letters = new Set(["a","b","c"]);
letters.add("d");
letters.add("e");
letters.add("f");

let map = new Map();
map.set("name", "Alice");
map.set("age", 25);
console.log(map.get("name")); // "Alice"
```

- `let`: variable that can be reassigned
  - block-scoped
  - cannot be redeclared within the same scope
- `const`: variable that cannot be reassigned
  - block-scoped
  - cannot be redeclared within the same scope
  - if it holds an object or array, it can still be mutated
- `var`: variable that can be reassigned and redeclared
  - function-scoped
- `typeof`: returns the type of the variable
- `instanceof`: returns boolean value indicating weather the object is an instance of the specified constructor or class

- `==`: compares value
  - performs type coercion: converts the operands to the same type before comparing
- `===`: compares value and type
  - doesn't do type coercion

### arrays

- `arr1.length`: property that returns the number of elements in an array
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

- array methods that don't mutate the original array:
  - `map()`: creates a new array by applying a function to each element of an existing array
    - `const doubledNumbers = numbers.map((number) => number * 2);`
  - `filter()`: creates a new array with elements that pass a test specified by a function
  - `reduce()`: applies a function to each element of an array to reduce the array to a single value
  - `forEach()`: calls a function for each element in an array
  - `some()`: checks if at least one element in an array passes a test specified by a function
  - `every()`: checks if all elements in an array pass a test specified by a function
  - `find()`: returns the first element in an array that passes a test specified by a function
  - `findIndexOf()`: returns the index of the first element in an array that passes a test specified by a function

```javascript
let numbers = [1, 2, 3, 4, 5];
let doubledNumbers = numbers.map((number) => number * 2);
console.log(doubledNumbers); // Outputs: [2, 4, 6, 8, 10]

let evenNumbers = numbers.filter((number) => number % 2 === 0);
console.log(evenNumbers); // Outputs: [2, 4]

let firstEvenNumber = numbers.find((number) => number % 2 === 0);
console.log(firstEvenNumber); // Outputs: 2

let index = numbers.findIndex((number) => number % 2 === 0);
console.log(index); // Outputs: 1

const colors = ['red', 'green', 'blue', 'yellow', 'purple'];

colors.forEach((color) => {
  console.log(`The color is: ${color}`);
});

const num = [1, 2, 3, 4, 5];
const sum = nums.reduce((accumulator, currentValue) => accumulator + currentValue, 0);
```

### template literals

> used to embed expressions into strings

- are created using backticks instead of quotes
- similar to f-strings in Python
- interpolation is done using `${}`

```javascript
const name = 'John';
const age = 30;

const greeting = `Hello, my name is ${name} and I am ${age} years old.`;
console.log(greeting); // outputs "Hello, my name is John and I am 30 years old."
```

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

### sets

> collection of unique values (no duplicates)

```javascript
let mySet = new Set([1, 2, 3, 4, 4, 5]);
mySet.add(6)
mySet.has(4)    // returns true
mySet.delete(4)
mySet.size      // returns the number of unique items
```

### maps

> collection of key-value pairs where keys can be of any type

```javascript
let myMap = new Map();
myMap.set('key', 'value');
myMap.set(1, 'one');
myMap.set('key', 'value');
myMap.get('key');
myMap.has('key');
myMap.delete('key');

let capitalCities = new Map([
    ['France', 'Paris'],
    ['Spain', 'Madrid'],
    ['Italy', 'Rome']
]);
console.log(capitalCities.get('Spain'));  // 'Madrid'
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
// while loop
let i = 0;
while (i < 5) {
  console.log('while:', i);
  i++;
}

// for loop
for (let j = 0; j < 5; j++) {
  console.log('for:', j);
}

// do...while loop
let k = 0;
do {
  console.log('do...while:', k);
  k++;
} while (k < 5);

// for...of loop
let fruits = ['apple', 'banana', 'cherry'];
for (let fruit of fruits) {
  console.log('for...of:', fruit);
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

## object oriented programming

- `class`: used to define a new class
  - don't forget to capitalize the first letter of the class
  - defining properties in the class is optional
- object: hash map that contains a collection of string-value pairs
- `constructor()`: special method for creating and initializing an object created with a class
- `this`
  - in a method (object function): refers to the object the method is called on
  - in a regular function: refers to the global object
  - in an arrow function: `this` is lexically bound, it refers to the surrounding context where the function was defined
  - can be used to allow method chaining: [code example in builder implementation](./design_patterns.md#builder)
- `new`: call the constructor function and returns a new object
  - defining a constructor is optional in javascript
  - if not defined, a default one is used
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

const myPerson1 = new Person("John", 30);
const myPerson2 = new Person("Alice", 25);
```

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
delete person.age;             // removes age, person.age changes to undefined
```

#### JSON (JavaScript Object Notation)

> stringified representation of a JavaScript Object

> [!IMPORTANT]
> json is not javascript

- semi-structured data
- lightweight data interchange format
- contains keys/value pairs
- keys are always strings and must be surrounded by double quotes
- doesn't have comment syntax (e.g. `//` or `/* */`)
- values can be strings, numbers, object literals, arrays, booleans (`true` or `false`) or `null` (not undefined)
  - OBS: JSON doesn't support `undefined` as a value
- can be parsed and converted to JavaScript objects with `JSON.parse()`

- json vs javascript:
  - json:
    - only double quotes allowed
    - no comments
  - javascript object
    - can use single quotes
    - has comments

```json
{
  "name": "John",
  "age": 25,
  "isStudent": true,
  "job": null,
  "father": { "name": "Bob", "age": 40 },
  "fruits": ["apple", "banana", "cherry"]
}
```

- `JSON.stringify()`: converts a JavaScript object to a JSON string
  - `const text = JSON.stringify(obj);`
- `JSON.parse()`: converts a JSON string to a JavaScript object
  - `const obj = JSON.parse(text);`

- common use cases:
  - HTTP request and response bodies
  - formats for text files
    - `.json` files are often used as configuration files
  - in NoSQL databases like MongoDB, ElasticSearch and Firestore

## modules

> file that contain code that can be imported/exported from/to other files

examples of `import` statements:

```javascript
import foo from 'module-name';               // default import
import { foo, bar } from 'module-name';      // named import
import { foo as myFoo } from 'module-name';  // named import with alias
import * as myModule from 'module-name';     // namespace import
import 'module-name';                        // import for side effects only
```

- to import a function from another file, it must be exported from the file where it’s defined
- variables/functions are not exported by default
- each module has its own scope
- 2 types of exports:
  - named exports: when you use only the `export` keyword
  - default exports: when you use the `export default` keywords

- `import`: used to import functions, objects, or primitive values from a file
- `export`: used to export functions, objects, or primitive values from a file
  - requires curly braces when importing
  - you can use export before a function, class, or variable declaration
  - OR you can use a single export statement at the end of the file for everything
    - `export { name, draw, reportArea, reportPerimeter };`
- `default`: used to export a single value from a file
  - can be imported without curly braces
  - `import greet from './module.js';`

- OBS: you can have both named and default exports in the same file
  - `import greet, { name, sayName } from './module.js';`

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
import { addTwo as add, subtract } from './modules/mathFunctions.js';

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

- promise: object that represents a value that might available now, in the future, or never
  - prevents callback hell
  - object can be in 3 states:
    - pending: operation in progress
    - resolved (fulfilled): operation completed successfully
    - rejected: operation failed
- `async`/`await` (recommended)
  - `async`: marks a function that returns a Promise
  - `await`: pauses the execution of the function until a Promise is resolved or rejected
    - frequently used with `fetch('https://api.example.com/data')` to make HTTP requests
  - more readable
  - prevents callback hell
- callback (not recommended): function passed as argument to another function that gets executed once the asynchronous operation is done
  - hard error handling
  - hard to read because of callback hell
    - callback hell: nested callbacks makes code hard to read

async/await example:

```javascript
async function fetchData() {
  const response = await fetch('https://api.example.com/data');
  const data = await response.json();
  console.log(data);
}
```

callback example (not recommended):

```javascript
function fetchData(callback) {
  setTimeout(() => {
    callback("Data received");
  }, 2000); // Simulates a 2-second delay
}

fetchData((data) => {
  console.log(data);  // Logs "Data received" after 2 seconds
});
```

example of promise:

```javascript
function fetchData() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve("Data received");
    }, 2000);
  });
}

fetchData()
  .then((data) => {
    console.log(data);  // Logs "Data received" after 2 seconds
  })
  .catch((error) => {
    console.error(error);  // Handles errors
  });
```

## conventions

- `#` indicates that a variable is private and should not be accesses directly from outside the class
  - e.g. `#age`

## javascript execution model

> javascript execution requires the cooperation of two pieces of harwdware: javascript engine and host environment

> javascript runs inside an environment, and each environment gives it new functionalities

- browser: `document`, `window`, events (e.g. `click`)
- node.js: `require()`, file system access, etc
- react native: its own API and event system
- deno: its own API and event system

## event loop

> waits for and handles events

- functions:
  - wait for user actions
  - runs event handler functions
  - timers like `setTimeout`

- is part of the host environment (e.g. browser engine or node.js)
  - implemented by the browser engine (in chrome: blink + V8)
  - implemented in node using the library libuv

## ECMAScript

> standardized especification that defines how the javascript language should work

- blueprint for javascript
- javascript is an implementation of ECMAScript, just like browsers (chrome) or runtimes (node.js) follow this blueprint
- new ECMAScript versions are released yearly, adding incremental improvements
- modern javascript incorporates features from ES6 and later versions

- ES5: ECMAScript 5 (2009)
- **ES6** (major update): ECMAScript 6 (2015)
  - also known as ES2015
  - introduced:
    - `let` and `const` for variable declarations
    - arrow functions
    - multi-line strings
    - default parameters
    - template literals
    - ternary operator
    - classes
    - modules
    - spread and rest operators (`...`)
    - promises
- ES7 (2016), ES8 (2017), ES9 (2018), ..., ES15 (2024)

