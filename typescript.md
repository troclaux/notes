
# typescript

> typed programming language that is a superset of javascript with new features

- all javascript code is also valid typescript code
- typescript adds features on top of javascript
  - static typing
  - type aliases
  - interfaces
  - enums
  - generics
  - etc
- uses `tsconfig.json` for compiler configuration

- CLI commands
  - run typescript file: `tsc <file_name.ts>`
  - create a `tsconfig.json`: `tsc --init`

## basic syntax

- `boolean`
- `string`
- `number`
- `bigint`
- `null`
- `undefined`
- `symbol`

```typescript
let isDone: boolean = false;
let count: number = 10;
let name: string = "John";

// not recommended
let cpf: any = "123.456";

// array declaration
let numbers: number[] = [1, 2, 3];

// tuple declaration
let tuple: [string, number] = ["Alice", 25];

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
  private name: string;      // must assign it
  private nickname?: string; // optional

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

## interface and type

- `interface` and `type` are used to define custom types
  - `interface` is primarily used to define object types
    - recommended because it provides better error messages
  - `type` is used to define object types, union types, tuple types, etc

```typescript
interface Person {
  name: string;
  age: number;
}

interface Employee extends Person {
  role: string;
}
```

```typescript
type StringOrNumber = string | number;

type Data = [number, string];

type Person = {
  name: string;
  age: number;
};
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

## modules

```typescript
export const name: string = "Alice";
import { name } from './moduleFile';
```

```typescript
// default import
import React from 'react';
import myFunction from './utils';

// named imports
import { useState, useEffect } from 'react';
import { add, subtract } from './math';

// ranaming imports
import { add as addition } from './math';

// import all
import * as mathUtils from './math';
mathUtils.add(2, 3);

// default + named imports
import React, { useState } from 'react';
```

## cli

- create a `tsconfig.json` and initialize a project: `tsc --init`
- convert `.ts` file to `.js`: `tsc example.tsc`

