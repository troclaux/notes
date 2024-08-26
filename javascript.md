
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
