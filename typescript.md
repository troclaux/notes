
## basics

example of types, functions and classes:
```typescript

let isDone: boolean = false;
let count: number = 10;
let name: string = "John";

// union types
let value: string | number;
let value = "Hello";  // Valid
let value = 42;       // Valid
// let value = true;  // Error: Type 'boolean' is not assignable to type 'string | number'

// type alias
type StringOrNumber = string | number;

let id: StringOrNumber;
id = 101;    // Valid
id = "abc";  // Valid

function add(x: number, y: number): number {
    return x + y;
}

let sum: number = add(5, 10);  // sum is inferred as number

class Animal {
    private name: string;

    constructor(name: string) {
        this.name = name;
    }

    public speak(): void {
        console.log(`${this.name} makes a noise.`);
    }
}

let dog = new Animal("Dog");
dog.speak(); // Output: Dog makes a noise.

```

## generics

> allow you to create reusable components that can work with any data type


```typescript

function identity<T>(arg: T): T {
    return arg;
}

let output1 = identity<string>("Hello");
let output2 = identity<number>(42);

```

## enums

```typescript

enum Direction {
    Up,
    Down,
    Left,
    Right
}

let dir: Direction = Direction.Up;

```

## sum types

> type that can hold different kinds of values, but only one at a time

```typescript

// Define the sum type
type Shape =
  | { kind: "circle", radius: number }
  | { kind: "rectangle", width: number, height: number }
  | { kind: "triangle", sideA: number, sideB: number, sideC: number };

// Example usage
function getArea(shape: Shape): number {
    switch (shape.kind) {
        case "circle":
            return Math.PI * shape.radius * shape.radius;
        case "rectangle":
            return shape.width * shape.height;
        case "triangle":
            // Heron's formula for the area of a triangle
            const s = (shape.sideA + shape.sideB + shape.sideC) / 2;
            return Math.sqrt(s * (s - shape.sideA) * (s - shape.sideB) * (s - shape.sideC));
    }
}

// Example shapes
const circle: Shape = { kind: "circle", radius: 10 };
const rectangle: Shape = { kind: "rectangle", width: 5, height: 10 };
const triangle: Shape = { kind: "triangle", sideA: 3, sideB: 4, sideC: 5 };

console.log(getArea(circle));     // Output: 314.159...
console.log(getArea(rectangle));  // Output: 50
console.log(getArea(triangle));   // Output: 6

```
