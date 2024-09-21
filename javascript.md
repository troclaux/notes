# JavaScript

## properties

- dynamically typed
- interpreted programming language
- has garbage collection
- first class functions
- hoisted declarations: variables and functions can be used before they are declared
- almost everything is an object
- supports multiple programming paradigms:
  - object oriented
  - functional
  - imperative
- immutable strings
- weakly typed
- multiplatform

### data types

> [!IMPORTANT]
> in JavaScript, almost everything is an object

- Number: represents both integers and floating-point numbers
- String
- Boolean: `true` or `false`
- Null
- Undefined: variable that has been declared but not assigned a value
- Object: collections of related data and/or functionality
  - arrays are Object
- Symbol
- BigInt

- `typeof`: returns the type of the variable
- `instanceof`: returns boolean value indicating weather the object is an instance of the specified constructor or class

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
- `const`: variable that cannot be reassigned
  - block-scoped
- `var`: variable that can be reassigned and redeclared
  - avoid using `var`
  - function-scoped

#### arrays

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

array methods that don't mutate the original array:
- `map()`: creates a new array by applying a function to each element of an existing array
  - `const doubledNumbers = numbers.map((number) => number * 2);`
- `filter()`: creates a new array with elements that pass a test specified by a function
- `reduce()`: applies a function to each element of an array to reduce the array to a single value
- `forEach()`: calls a function for each element in an array
- `some()`: checks if at least one element in an array passes a test specified by a function
- `every()`: checks if all elements in an array pass a test specified by a function

- `find()`: returns the first element in an array that passes a test specified by a function
- `findIndexOf()`: returns the index of the first element in an array that passes a test specified by a function

examples

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


#### template literals

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

#### type casting

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

## loops

```javascript
let i = 0;

while (i < 5) {
  console.log(i);
  i++;
}
```

```javascript

// for loop
for (let i = 0; i < 5; i++) {
  console.log(i);
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

## ternary operator

> `condition ? return_this_value_if_true : return_this_value_if_false`

```javascript
let x = 5;
let y = (x < 10) ? "x is less than 10" : "x is greater than or equal to 10";
console.log(y); // Output: "x is less than or equal to 10"
```

## switch case

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

const myPerson1 = new Person("John", 30);
const myPerson2 = new Person("Alice", 25);

```

initializing an object
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

#### JSON (JavaScript Object Notation)

- lightweight data interchange format
- contains name/value pairs
- values can be strings, numbers, objects, arrays, booleans, or null (not undefined)
  - OBS: JSON doesn't support `underfined` as a value
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

## modules

- to import a function from another file, it must be exported from the file where it’s defined
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

callback < promise < async/await

- async/await
  - `async`: makes a function asynchronous
  - `await`: pauses the execution of the function until the promise is resolved
    - frequently used with `fetch('https://api.example.com/data')` to make HTTP requests
  - recommended
  - more readable
  - prevents callback hell
- callback: function passed as argument to another function that gets executed once the asynchronous operation is done
  - not recommended
  - hard error handling
  - hard to read because of callback hell
    - callback hell: nested callbacks makes code hard to read
- promise: is an object that represents a value that might available now, in the future, or never
  - not recommended
  - hard to read
  - hard error handling
  - prevents callback hell
  - object can be in 3 states:
    - pending: operation in progress
    - resolved (fulfilled): operation completed successfully
    - rejected: operation failed

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
