
# golang

## properties

- syntax is similar to C
- has strong standard library
- has strong concurrency support
- is not an object-oriented language
  - uses structs instead of classes
- doesn't have `while`, uses `for` instead
  - `for sum < 10 { sum += 1 }`
- does not require semicolons at the end of every statement

- statically typed
- compiled language
- has garbage collection
- functions are values (i.e. first-class citizens)
- hoisted declarations: variables and functions can be used before they are declared
- no implicit type casting
- mutable strings
- multiplatform
- block-scoped: once a variable is declared inside a block, it is accessible anywhere inside that block

- value types: Integers, Floats, Complex numbers, Strings, Bytes, Arrays, Structs
- reference types: Pointers, Slices, Maps, Channels, Interfaces, Functions

## cli commands

- `go mod init github.com/troclaux/<repo-name>`: creates a new module
  - module: collection of packages
  - `package`: directory of related .go files
  - writes a new go.mod file in the current directory
  - use the repo's URL to create a module when the project is hosted in a github repository
- `go mod tidy`: clean up go.mod file
  - removes unused dependencies
  - add missing dependencies
- `go run main.go`: compiles and runs go code without creating binary
- `go build main.go`: compiles go code and produces a binary executable
- `go fmt .`: formats go code according to the standard go style guidelines
- `go install github.com/charmbracelet/glow@latest`: builds and install go package in `$GOPATH/bin` or `$HOME/go/bin`
- `go version`: prints the installed go version
- `go test`: runs unit tests
  - searches files with names matching `*_test.go`
- `go get`: downloads and installs packages and dependencies

## directory structure

initialize module by running one of these commands:
- `go mod init`
- `go mod init module_name`
- `go mod init github.com/user/project`

```
my-project/
├── go.mod
├── go.sum (generated later)
├── main.go
└── (other directories and files)
```

- `go.mod`
  - stores module's name (e.g. `github.com/user/project`)
  - tracks module's dependencies
  - stores the go version used to compile the project
- `go.sum`: ensures integrity of dependencies
  - saves the versions of modules being used and their checksums
- `main.go`: entrypoint of your go program
  - typically located on the root of your project
  - execution of program begins here

- module: collection of go packages
  - stores packages in a file tree with a go.mod file at its root
  - packages in a module should have distinct names
- package: directory of related .go files
  - `package main`: contains exactly 1 `main.go` file, which is the entry point of the program
  - packages by other names are "library packages"
    - library package: only export functionality that can be used by other packages
      - have no entry point
  - multiple packages on the same directory are **NOT** allowed
  - all .go files in a single directory must belong to the same package
    - if they don't, the compiler will throw an error
- `import`: imports packages
  - `import "fmt"`: import package with basic formatting functions (includes Println)
  - when importing multiple packages, don't use comma between each package

## basics

```go

package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println("Hello, World!")
  fmt.Println(math.Pi)
}

```

- variables declared without explicit initial value are given their zero value
  - 0 for numeric types
  - false for the boolean type
  - "" for strings (empty string)
- `// single-line comment`
- `/* multi-line comment */`

- `var a int = 10`: declares a variable `a` of type `int` and assigns it a value of `10`
- `b := 20`: go can infer the type from the assigned value
- `var c float64`: declares a variable `c` of type `float64` with the zero value (0.0)
- `const d = 5`
  - you can't declare a constant that can only be computed at runtime
    - e.g. `const currentTime = time.Now()`

```go
package main

import "fmt"

func main() {

  // this is a single line comment

  /*
    this is a multi-line comment
    this text will be ignored
  */

  var a int = 10
  b := 30.5 // Type inferred as float64
  var c float64
  const d = 5
  var e = 20

  fmt.Println(a, b, c, d, e)
}
```

- `range`: used to iterate over strings, arrays or slices
  - `for index, value := range collection {}`
- blank identifier (`_`): special variable that is used to ignore a value or variable
  - normally used for error handling or discarding index in loops that use `range`

## data types

> [!IMPORTANT]
> in Golang, you cannot define a type inside a function (including `func main()`)
> types are declared at the package level, not at the function level

- basic types:
  - numeric types:
    - integers: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
      - the number after `int` is the number of bits
      - `int`: alias for `int64` in 64-bit systems and `int32` in 32-bit systems
      - `byte`: alias for `uint8`
        - `string()`: converts slice of bytes into a string
      - `rune`: alias for `int32`
	- single-quoted strings `''` are used for runes
    - floating-point numbers: `float32`, `float64`
    - complex numbers: `complex64`, `complex128`
  - Booleans: `bool`
  - Strings: `string`
    - double-quoted strings (`""`) are used for string literals
- composite types:
  - `array`
  - `slice`
  - `struct`
    - `type Car struct {}`
  - `map`
  - `chan`
- reference types:
  - `pointer`
  - `function`
- interface types:
  - `interface`
    - `type Vehicle interface {}`

> [!TIP]
> Unless you have a good performance reason, use the default types

- default types: bool, string, int, uint, byte, rune, float64, complex128

> [!TIP]
> use `reflect.TypeOf(my_variable)` to get the type of a variable

### zero values

> variables declared without an explicit initial value are given their zero value

- numerics: `0`
- strings: `""`
- booleans: false
- struct: `structName{}`
  - fills each field with the corresponding type's zero value
- array: zero value of each element
- slice: `nil`
  - `fmt.Println(zeroValueSlice)` will print `[]`
- function: `nil`
- maps: `nil`
  - `fmt.Println(zeroValueMap)` will print `map[]`
- pointers: `nil`
  - `fmt.Println(zeroValuePointer)` will print `<nil>`
- interfaces: `nil`
  - `fmt.Println(zeroValueInterface)` will print `<nil>`
- channels: `nil`
  - `fmt.Println(zeroValueChannel)` will print `<nil>`
- generics: zero value of the type parameter used

### type casting

> [!NOTE]
> to floor a division, use type casting to remove the decimal part

- it only floors the number, removing the decimal digits
- `num := int(25.99) // num = 25`

```go
package main

import "fmt"

func main() {
  var a int = 42
  var b float64 = float64(a) // Convert int to float64
  var c int = int(b)

  fmt.Println(a) // Output: 42
  fmt.Println(b) // Output: 42.0
  fmt.Println(c) // Output: 42
}
```

### type assertion

- `v`: variable of unknown type
- `person, ok := v.(Person)`
  - `person`: receives the value of `v` if `v`'s type is `Person`
  - `ok`: receives true if `v`'s type is `Person`, false otherwise


```go
package main

import "fmt"

type Person struct {
  name string
  age  int
}

func main() {
  p := Person{"John", 20}

  // Type assertion to get the Person struct
  fmt.Println(p.name) // Output: John

  // simpler example
  var x interface{} = "hello"

  value, isType := x.(string)  // Type assertion
  if isType {
    fmt.Println(value)       // Output: hello
  }
}
```

use type switch for multiple type assertions:

```go
switch v := num.(type) {
case int:
  fmt.Printf("%T\n", v)
case string:
  fmt.Printf("%T\n", v)
default:
  fmt.Printf("%T\n", v)
}
```

### type alias

> [!IMPORTANT]
> Type names and aliases should start with a capital letter (PascalCase)
> variable names should start with a lowercase letter (camelCase)

```go
package main

import "fmt"

func main() {
  type Numeric int
  var num Numeric = 2
  fmt.Printf("%T", num)     // Output: main.Numeric
}
```

### enums

- `iota`: increments automatically each new constant in a block
- `const()`: declares multiple constants in a single statement

```go
package main

import "fmt"

type Weekday int

// Define an enumeration for weekdays
const (
  Sunday Weekday = iota  // iota starts at 0
  Monday
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday
)

func main() {
  fmt.Println(Sunday, Monday, Tuesday)  // Output: 0 1 2
}
```

### interface

[best practices for interfaces in Go](https://blog.boot.dev/golang/golang-interfaces/)

> type that defines a contract or a set of methods that must be implemented by any type that satisfies the interface

> collections of method signatures

> [!IMPORTANT]
> you don't need to name the arguments of an interface
> but it's recommended to name the arguments, to make it more readable
> you must specify the input of the method and the output type
> interfaces don't accept fields, only methods (although it can accept [constraints](#constraints))

- method signature: method name, receiver type and the parameter list
  - e.g. `func (p Person) sayHello()`
- you cannot define a type that implements an interface inside a function
- a type can implement any number of interfaces in golang
- structs implements interfaces implicitly
  - you may add methods to a type and in the process be unknowingly implementing various interfaces, and that's okay
  - if a struct implements a method that is inside an interface, then it must implement all methods in the interface

- good practices for interfaces:
  - keep interfaces small
  - name the arguments of the interface
  - interfaces should focus only on the actions or behavious it defines, not on the specific types that implement it
    - don't add type-specific methods (like IsFiretruck()) to a general interface
    - types can be aware of the interfaces they satisfy
    - interfaces can't be aware of the types they satisfy

```go
package main

import "fmt"

func main() {
  rect := Rectangle{Width: 5, Height: 3}
  fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
  fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())
}

type Shape interface {
  Area() float64
  Perimeter() float64
}

// Define a struct for a Rectangle
type Rectangle struct {
  Width  float64
  Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

// Implement the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
  return 2 * (r.Width + r.Height)
}
```

#### empty interface

- every type implements the empty interface (`interface{}`)
- the empty interface can hold values of any type
  - `var i interface{} = "hello"`
    - holds the value's type (`string`) and the value itself (`hello`)

## data structures / composite types

> [!NOTE]
> use `slices.Sort(arr1)` to sort any ordered built-in type
> requires importing `slices` package

- array: ordered collection of elements with fixed sizing
  - `var arr [5]int`
- slice: ordered collection of elements with dynamic resizing
  - `s1 := []int{1, 2, 3}`
  - `s2 := make([]int, 5, 10)`: initializes an object of type: `slice`, `map` or `chan`
    - length of s2 = 5
      - number of elements that are initialized with zero values
    - capacity of s2 = 10
- map: key-value pairs hash table
  - `m1 := map[string]int{ "apple":  1, "banana": 2, }`
  - `m2 := make(map[string]int)`
    - for maps, there's no need to define length
- struct: collection of fields that can be of different type

```go
type Person struct {
    name string
    age  int
}
p := new(Person)
```

- `make(type, length, capacity)`: function creates slices, maps and channels
  - type of data structure to create
  - length: number of elements that are initialized with zero values
  - capacity: number of items that the slice can hold without needing to be resized
    - capacity's default value is the length

> [!NOTE]
> go doesn't have built-in stacks and queues, but they can be implemented using slices

### stack (implemented with slices)

```go
stack := []int{}
stack = append(stack, 1) // Push
stack = append(stack, 2) // Push
stack = append(stack, 3) // Push

top := stack[len(stack)-1]  // Peek
stack = stack[:len(stack)-1] // Pop
fmt.Println(top)            // 3
```

### queue (implemented with slices)

```go
queue := []int{}
queue = append(queue, 1) // Enqueue
queue = append(queue, 2) // Enqueue

front := queue[0]       // Peek
queue = queue[1:]       // Dequeue
fmt.Println(front)      // 1
```

### structs

> groups fields together

- replaces classes
- structs can have methods
- methods in Golang are just functions that have a receiver
- receiver: special parameter that syntactically goes before the name of the function
  - receivers can receive pointers or just structs
    - what is the difference?
      - pointer receivers: pass a pointer to the original structure
        - object-oriented programming
      - value receivers: pass a copy of the original structure
        - functional programming
  - pointer receiver are more widely used

- methods and constructors aren't defined inside the struct initialization
- `new()`: creates a new instance of a struct
  - `p := new(Person)`
  - `p.Name = John`

```go
package main

import "fmt"

type Person struct {
  Name string
  Age  int
}

func main() {
	// creating object
  p1 := Person{Name: "Alice", Age: 30}
  fmt.Println(p1.Name) // Output: Alice
  fmt.Println(p1.Age)  // Output: 30

  // calling method
  p1.Greet() // Output: Hello, my name is Alice and I am 30 years old.

  // using constructor
  p2 := NewPerson("Charlie", 20)
  p2.SetAge(45)
  p2.Greet()
}

// creating a method with a value receiver
func (p Person) Greet() {
  fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// creating a method with a pointer receiver
func (p *Person) SetAge(age int) {
  p.Age = age
}

// constructor
func NewPerson(name string, age int) *Person {
  return &Person{Name: name, Age: age}
}

```

#### embedded struct

- just add the struct inside the parent struct without the type
- can be used to emulate inheritance

```go
type car struct {
  make string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}
```

#### empty struct

> smallest possible type in golang

- it occupies 0 bytes of memory
- `struct{}` is the type (empty struct)
- `struct{}{}` is the value (empty struct literal)

```go
// anonymous empty struct type
empty := struct{}{}

// named empty struct type
type emptyStruct struct{}
empty := emptyStruct{}
```

#### anonymous struct

> struct that is defined without a name

- can't be referenced elsewhere in the code
- anonymous structs can be nested inside anonymous structs
- when should i use anonymous structures?
  - generally try not to use it
  - only use it when you won't reuse the struct

```go
myCar := struct {
  make string
  model string
} {
  make: "tesla",
  model: "model 3",
}
```

### arrays

> [!IMPORTANT]
> in golang, arrays are value types, instead of reference types
> arrays are passed as values to a function, instead of a reference

```go
package main

import "fmt"

func main() {
  var arr [6]int = [6]int{1, 2, 3, 4, 5}
  arr[0] = 6
  fmt.Println(arr) // Output: [6 2 3 4 5 0]
}
```

### slices

> [!IMPORTANT]
> slices are reference types
> slices are passed as references to a function
> if you change a slice that points to an array, you will change the array
> to avoid this accident: `copy(sliceScopy, arr[:])`

- `len(arr1)`: returns length
- `cap(arr1)`: returns capacity

```go
package main

import "fmt"

func main() {

  s1 := make([]int, 5)    // length 5, capacity 5
  s2 := make([]int, 3, 5) // length 3, capacity 5

  s3 := []int{1, 2, 3, 4, 5} // Slice literal

  s3 = append(s3, 6) // Append an element
  fmt.Println("s1 =", s1)
  fmt.Println("s2 =", s2)
  fmt.Println("s3 =", s3)

  // output:
  // s1 = [0 0 0 0 0]
  // s2 = [0 0 0]
  // s3 = [1 2 3 4 5 6]
}
```

- array: `arr := [3]int{1, 2, 3}`
- slice: `arr := []int{1, 2, 3}`

> [!WARNING]
> `slice = append(slice, 1, 2, 3, 4)`: add items to end of slice
> `append()` changes the underlying array of its parameter AND returns a new slice
> ALWAYS ASSIGN THE RESULT OF APPEND() BACK TO THE SLICE TO ENSURE THE CHANGES ARE APPLIED CORRECTLY
> `append()` only works for slices
> don't do this: `someSlice = append(otherSlice, element)`
> you can accidentally declare multiple slices that share the same underlying array
> if this happens, when you change one slice, you can change other slices without knowing

#### matrix (slice of slices)

declaring an empty matrix:
1. initialize the matrix: `matrix := make([][]int, rows)`
1. loop to initialize each row of the matrix
1. initialize each row: `matrix[i] = make([]int, cols)`
1. loop to define each element's new value: `matrix[rowIdx][colIdx] = rowIdx * colIdx`

```go
func createMatrix(rows, cols int) [][]int {
  matrix := make([][]int, rows)
  for i := 0; i < rows; i++ {
    matrix[i] = make([]int, cols)
    for j := 0; j < cols; j++ {
      matrix[i][j] = i * j
    }
  }
  return matrix
}
```

### maps

> data structure with key -> value mapping

> [!IMPORTANT]
> maps are reference types
> maps are passed as references to a function
> you need to initialize a map before you can add key-value pairs to it
> `map1 := make(map[string]int)`

- similar to python's dictionaries and javascript's objects
- O(1) lookup
- maps can contain maps, also called nested maps
- map keys can be of any type that is comparable
  - structs can be keys
    - `users := make(map[User]int)`

- delete element from map: `delete(map1, key)`
- checks if a map contains a key: `elem, ok := map1[key]`
  - if `key` is in `map1`, then `ok` is `true` and `elem` is the value as expected
  - if `key` is not in `map1`, then `ok` is `false` and `elem` is the zero value for the map's element type

> [!NOTE]
> when trying to access a value in a map using a key that doesn't exist, the map returns the "zero value" for that type
> you don't need to initialize a new entry of a map in order to increment it

```go
m := make(map[string]int)
m["foo"] = 1
m["bar"] = 2

// example of map literal
ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
```

## packages

> collection of go files

- multiple packages on the same directory are **NOT** allowed
- all .go files in a single directory must belong to the same package
  - if they don't, the compiler will throw an error

2 types of packages:
- `package main`: contains exactly 1 `main()` function, which is the entry point of the program
- library package: any package that isn't the main package
  - can only export functionality to be used by other packages
  - has no entry point

> [!IMPORTANT]
> each package has its own namespace
> which means you can define variables, functions and types without worrying about conflicts with other packages

requirements for one package to import another package's variable:
- variable's first letter must be capitalized
- the package that has the variable must be imported by the package that is trying to access the variable

visibility is controlled by capitalization of the first letter:
- identifiers that start with a capital letter are exported from the package
  - can be accessed by anyone outside the package that declares it
  - Capitalized = public
- identifiers that don't start with a capital letter are not exported and are private to the package
  - not-capitalized = private

- what can be exported by capitalizing the first letter?
  - variables
  - constants
  - functions
  - structs
  - enums
  - interfaces

- best practices for golang packages:
  - hide internal logic
  - don't change APIs
    - avoid breaking changes
  - don't export functions from the main package
    - the main package isn't a library, there's no need to export functions from it
  - packages shouldn't know about dependents
    - a package should never have specific knowledge about a particular application that uses it

### visibility rules

- fields and functions that start with an uppercase letter are accessible from outside the package they are defined
- fields and functions that start with an lowercase letter are private to the package

- exported fields can be accessed by external code
  - external code: standard libraries
- unexported fields cannot be accessed outside the package, including by standard library

```go
package main
import "fmt"

type parameters struct {
  Name  string
  grade string // Unexported field
}

func main() {
  p := parameters{
    Name:  "Alice",
    grade: "A",
  }
  fmt.Println(p) // Output: {Alice A}
}
```
## modules

> collection of go packages stored in a file tree with a go.mod file at its root

- `go.mod`: file that stores:
  - module path
  - version of go language the project requires
  - all external package dependencies
- variables, constants, functions and types defined in one source file are visible to all other source files in same package
- repository: contains 1 or more modules

- module path: string used to import a package
  - e.g. `github.com/google/go-cmp`
  - packages in the standard library do not have a module path prefix
  - usually there is only one module per repo

- `GOPATH`: environment variable that is deprecated, avoid using it

## strings

> strings are structs with a byte pointer and the length of the string

- the size of the pointer depends on architecture of the system
  - 32 bit => 4 bytes pointer
  - 64 bit => 8 bytes pointer

```go
struct {
  *byte   // pointer to the first byte of the string
  int     // length of the string
}
```

- concatenate strings with `+`
  - `str1 := str2 + str3`

```go
package main

import "fmt"

func main() {
  var str1 string = "hello"
  var str2 string = " world"
  fmt.Println(str1 + str2)
}
```

## if else

> [!NOTE]
> you can add an initial statement before the condition
> a variable created in the initial statement is only defined within the scope of the `if` body

```go
package main

import "fmt"

func main() {

  if x := 10; x > 5 {
    fmt.Println("x is greater than 5")
  } else if x == 5 {
    fmt.Println("x is equal to 5")
  } else {
    fmt.Println("x is less than 5")
  }
}
```

## loops

```go
package main

import "fmt"

func main() {
  // Basic for loop
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  // While-style loop
  for x := 0; x < 5 {
    fmt.Println(x)
    x++
  }

  // Infinite loop
  for {
    fmt.Println("Infinite loop")
  }

  numbers := []int{1, 2, 3, 4, 5}
  for i, value := range numbers {
      fmt.Println(value)
  }
}
```

loops can have multiple statements:

```go
for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
    fmt.Println(i, j)
}
```

## switch statement

```go
package main

import "fmt"

func main() {
  day := 2
	
  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  default:
    fmt.Println("Another day")
  }
}
```

## defer

- `defer`: delay the execution of a function until the surrounding function returns
  - deferred function call are pushed on a stack
  - when function ends, the deferred call are executed in LIFO order

```go
package main

import "fmt"

func main() {
  defer fmt.Print("world")
  fmt.Print("hello ")
}
// output: "hello world"
```

## functions

- can return any number of results
- return values: define the return type of a function
  - located in the function signature, after the parameters
  - can be surrounded by parentheses (optional)

> [!IMPORTANT]
> Go doesn't allow declaring named functions inside other functions
> instead, you can define an anonymous function inside a function signature with the `func` keyword
> also, you can declare functions outside of the `main()` function

```go
package main

import "fmt"

// Function that returns an integer
func add(a int, b int) int {
  return a + b
}

// Function with multiple return values
func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  sum := add(3, 4)
  fmt.Println("Sum:", sum)

  a, b := swap("hello", "world")
  fmt.Println(a, b)
}
```

### anonymous functions

> [!NOTE]
> you can define an anonymous functions function inside a function signature with the `func` keyword

```go
package main

import "fmt"

func main() {
  // Anonymous function
  add := func(a, b int) int {
    return a + b
  }

  // Call the anonymous function
  result := add(5, 3)
  fmt.Println(result) // Output: 8
}
```

## pointers

> [!NOTE]
> golang doesn't have pointer arithmetic

- `&`: returns the memory address of its operand
  - address-of operator
- `*`: returns the value stored at the memory address
  - dereference operator

> [!IMPORTANT]
> you can't use `*` to read a field of a pointer to a struct
> wrong: `value := *myStruct.myField`
> correct: `value := myStruct.myField`

```go
package main

import "fmt"

func main() {
  var x int = 5
  var p *int = &x
  fmt.Println(*p) // Output: 5
  *p = 10
  fmt.Println(x) // Output: 10
}
```

## generics

> use variables to refer to specific types

- requires prefix `[T any]` before any function that uses generics
- declaring: `var myZero T`

```go
// T1 and T2 are generics
func genericFunction[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
    // Function logic
}
```

### constraints

> interface that restricts which type a generic can be

- union operator (`|`): restricts generic to multiple types
- `~`: the ~ before a type allows a type constraint to match any type with that specified type as its underlying type

```go
type numeric interface {
    ~int | ~float64
}

func add[T numeric](value T) ([]T) {
    // use generic that is int or float64
}
```

## concurrency

go provides built-in support for concurrency using goroutines and channels

> concurrency is as simple as using the `go` keyword before a function/method

```go
go doSomething()
```

- goroutine: lightweight thread of execution
- `go`: runs following function concurrently
  - spawns a new goroutine
  - can also be used with anonymous functions
    - you can suffix the anonymous function with `()` to run it immediately after definition
  - does not block the main goroutine
    - i.e. main goroutine will continue to run and does not wait for the other goroutines to finish

goroutine's order of complete execution is unpredictable, because it depends partially on the OS' scheduler. Even in a program where you run multiple identical goroutines that only have the same print statement, the order of the print statements will change each time you run the program.

```go
package main

import (
  "fmt"
  "time"
)

func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func main() {
  go say("Hello") // Run say() in a new goroutine
  say("World")    // Run say() in the main goroutine
}
```

### channels

- typed, thread-safe queue
- `middle-man` data structure that passes values between goroutines
- FIFO
- allows different goroutines to communicate with each other
- `range` keyword works for channels
  - `for item:= range ch {}`

- my informal explanation: you have 2 types of goroutines (not mutually exclusive):
  - senders: sends values to channels
    - channels will store values that will be consumed by another goroutine
    - the sender goroutine blocks (stops execution) until a receiver receives the value
  - receivers: receive values from channels
    - when the receiver goroutine tries to receive a value, it blocks until a sender sends a value through the channel

```go
package main

import "fmt"

func main() {
  ch := make(chan int) // Create channel

  go func() {
    ch <- 42 // Send a value into the channel
  }()     // These parentheses call the function immediately after defining it

  value := <-ch // Receive a value from the channel
  fmt.Println(value)
}
```

- `ch := make(chan int)`: creates an unbuffered channel that stores int value
  - unbuffered channels can't store specific number of values before blocking
  - example of channel parameter: `func foo(num int, myChan chan struct{}) {}`
- `ch <- 42`: sends a value into the channel
  - sending a value into a channel blocks the program until another goroutine is ready to receive the value
- `value := <-ch`: receives a value from the channel and removes the value
  - this instruction blocks the program until there is a value to be read

- sending a value to a nil channel blocks the program forever
- trying to receive a value from a nil channel blocks the program forever
- sending a value to a closed channel causes a panic
- trying to receive from a closed channel returns the zero value immediately

```go
var c chan string // c is nil
c <- "let's get started" // blocks
```

- `select`: switch case for channels
  - listens to multiple channels at the same time
  - will run the `case` block from the value that's received first
  - if multiple values are received at the same time, one of them is chosen randomly
- `ok` variable refers to whether or not the channel has been closed by the sender yet
- `i` and `s` will receive values from channels if they have any
- `default`: executes immediately if there's no values in channel
  - a `default` case stops the `select` statement from blocking

```go
select {
case i, ok := <-chInts:
  fmt.Println(i)
case s, ok := <-chStrings:
  fmt.Println(s)
default:
  fmt.Println("this runs immediately if there's no values in all channels")
}
```

example of using a `select` statement within an infinite loop to continuously check multiple channels for incoming data

```go
for {
  select {
  case <-ch1: //
  fmt.Println("channel 1")
  case <-ch2:
  fmt.Println("channel 2")
  default:
  fmt.Println("this runs immediately if there's no values in all channels")
  }
}
```

#### buffered channels

> buffered channels store a specific number of values before blocking

- `emailChan := make(chan string, 100)`: creates buffered channel
  - defines number of values to store before blocking
- don't create an array for buffered channels in function signatures
  - e.g. `func addEmailsToQueue(emails []string) chan string {`

#### read-only channels

```go
var readOnlyChan <-chan int
```

#### write-only channels

```go
var writeOnlyChan chan<- int
```

- reader is a goroutine that reads from a shared resource (normally a variable or data structure)
- writer is a goroutine that writes/changes/mutates a shared resource
- to acquire the lock means to gain exclusive access a shared resource
  - preventing other goroutines from accessing it simultaneously

## mutexes

> locks access to data and controls which goroutines can access certain data at which time

- requires importing `"sync"` package
- initialize like so: `var mu sync.Mutex`
- lock current goroutine: `mu.Lock()`
- unlock current goroutine: `defer mu.Lock()`
  - always `defer mu.Unlock()` to make sure eventually the goroutine gets unlocked

> [!WARNING]
> maps are not safe for concurrent use
> you must always `mu.Lock()` before writing to a map

- group the data with its synchronization mechanism in a `struct`
  - by bundling the mutex with the data, it's harder to accidentally access the data without using the mutex

```go
type SafeCounter struct {
  mu    sync.Mutex
  value int
}
```

## standard library

- library: collection of packages

### fmt

> provides functions for formatting and printing data

- format string: string that contains placeholders for the arguments, separated by commas
  - `fmt.Println("Hello, %s!", "world")`
  - `fmt.Printf("Name: %s, Age: %d, Height: %.2f meters\n", name, age, height)`

- `Print()`: prints arguments to the standard output
- `Println()`: prints arguments to the standard output, followed by a newline character
- `Printf()`: prints arguments to the standard output and allows the use of format strings

- `Sprint()`: returns a string representation of its arguments without printing
- `Sprintf()`: returns a string representation of its arguments without printing, also formats strings

- `Errorf("failed to get user: %w", err)`: returns error with a format string description

```go
fmt.Printf("%v\n", 42)          // Output: 42
fmt.Printf("%v\n", true)        // Output: true
fmt.Printf("%v\n", "hello")     // Output: hello
```

format specifiers are used when printing variables with `fmt` package:
- `%v`: works with every type
- `+%v`: for structs, prints the field names and values
- `%T`: type of the variable
- `%d`: decimal representation of integer
- `%f`: floating point decimal numbers
- `%.2f`: limit precision to 2 decimal places
- `%s`: string
- `%t`: boolean values
- `%p`: pointer
- `%b`: binary
- `%x`: hexadecimal representation with lowercase letters
- `%c`: character corresponding the integer's Unicode code point
- `%%`: print percent sign
- `%w`: error

### errors

go does not have exceptions: instead, errors are handled using the built-in `error` type

- error is any type that implements this error interface:

```go
type error interface {
    Error() string
}
```

- you can create your own structs that implements `error interface` that returns a custom `Error() string`
- the Go standard library provides an `"errors"` package that makes it easy to deal with errors
  - easy custom errors
    - `var err error = errors.New("error description")`
- use `errors` when you encounter a programming error that can be recovered from
- avoid using `panic()`
  - why?
    - `panic()` removes control of current function, up the call stack until it reaches a function that `defers` a `recover`
    - if no function calls `recover`, the goroutine (often the entire program) crashes
  - to cleanly exit the program in an unrecoverable way, use `log.fatal()`
- it's recommended to return the "zero" values of the type when you return a non-nil error

```go
package main

import (
  "errors"
  "fmt"
)

func divide(a, b float64) (float64, error) {
  if b == 0 {
    return 0, errors.New("cannot divide by zero")
  }
  return a / b, nil
}

func main() {
  result, err := divide(4, 0)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }
}
```

### math

The `math` package provides basic constants and mathematical functions for floating-point operations.

### net

> provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets

#### http

[http](/networking.md#http-(hypertext-transfer-protocol))

```go
package main

import (
  "fmt"
  "io"
  "net/http"
)

func getItemData() ([]byte, error) {
  res, err := http.Get("https://example.com")
  if err != nil {
     return nil, fmt.Errorf("error creating request: %w", err)
  }
  defer res.Body.Close()

  data, err := io.ReadAll(res.Body)
  return data, err
}
```

- initiate http get request: `res, err := http.Get("https://example.com")`
  - `res`: HTTP response that comes back from the server
  - during get request, Go opens a connection and reads the response into the `Body`
  - `defer res.Body.Close()`: always close `body` to free system resources
- `data, err := io.ReadAll(res.Body)`: reads the response body into a slice of bytes called `data`
- does it reads response into a slice of bytes?
- use `fmt.Errorf()` to format an error message
  - `fmt.Errorf("error message: %w", err)`

HTTP request pseudocode:

1. make request with http verb (this example shows GET)
1. return error if request failed
1. close resource (use `defer` to ensure it eventually gets closed)
1. create JSON decoder to decode the body of the HTTP response
1. create nil slice of `struct` that will store the response's decoded content
1. return error in case of failed decoding
1. return data structure with response's decoded content

```go
func getUsers(url string) ([]User, error) {
  res, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  decoder := json.NewDecoder(res.Body)
  var users []User
  if err := decoder.Decode(&users); err != nil {
    fmt.Println("error decoding response body")
  }
  return users, nil
}
```

- use `json.NewDecoder.Decode` when there's an `io.Writer` or `io.Reader`
  - when in doubt, use `json.NewDecoder.Decode`
  - how does `json.NewDecoder.Decode` knows which fields to store decoded json data?
    - json decoder reads `{ "body": "Hello world" }`, goes to the go struct, sees `json:"body"` and knows where to store the json data
- use `json.Unmarshal` when there's a `[]byte` slice

- use `enconding/json` package to map JSON fields to struct fields
  - struct fields must be exported (capitalized) to decode JSON
- `json.NewDecoder(res.body)`: creates new JSON decoder that can read and decode JSON from an `io.Reader` source
  - `io.Reader` source: files, HTTP response, network connection
- `decoder.Decode(&items)`: reads the JSON-encoded value from `decoder` and stores it in the value pointed to by `&items`

```go
Price float64 `json:"price"`
```

- `Price`: field name in Golang
- `float64`: field type in Golang
- `json:"price"`: field name in the original JSON

price (JSON) => Item.Price float64 (Golang)

```go
type Item struct {
  ID    int    `json:"id"`
  Name  string `json:"name"`
  Price float64 `json:"price"`
}

var items []Item
decoder := json.NewDecoder(res.Body)      // res is a successful `http.Response`
if err := decoder.Decode(&items); err != nil {
  fmt.Println("error decoding response body")
  return
}
```

> HTTP POST request sends data to a server, typically to create e new resource

- `http.Post` functions it's limited (?)
- `http.NewRequest` is better (?)

- payload: body of the request sent to the server
- content-type: format of the body (e.g. application/JSON)

POST request pseudocode:

1. encode data (normally a struct) as JSON
1. return error if encoding failed
1. create new request
1. return error if new request failed to be created
1. set headers
1. create new client and make the request
1. decode json data from the response
1. return decoded json data


- short networking review:
  - data is transmitted over the network as raw binary
  - Computers communicate using bytes
  - body of HTTP request needs to be in a readable stream of bytes
    - because HTTP transmits data over a network as byte streams
  - formats like JSON, text, or images must be converted (encoded) into a byte stream to be transmitted over the internet
    - JSON serialization: process that converts object/struct into a JSON string
  - the byte stream will be received and decoded on the receiving end
  - byte streams are highly efficient, because:
    - it's impractical to send large amounts of data at once
    - instead, data is broken into chunks (a stream) that can be sent gradually
    - this allows the transmission of large amounts of data without memory constraints
  - HTTP is built on top of TCP/IP, which is a byte-oriented protocol
  - all data in an HTTP request or response (headers, body, etc.) must ultimately be encoded into bytes
    - since that's the format in which TCP transmits data

what does it mean to say: HTTP is built on top of TCP/IP?
how is this data sent over the network with ip address to its destination?
if the data sent over a network is wrapped in someway, explain and illustrate how
give me example of http request and response

- before sending a Golang struct, it needs to be serialized into JSON
- `bytes.NewBuffer(jsonData)`: wraps JSON data into a buffer to be sent with the request body
  - what does it mean to wrap JSON data into a buffer?
    - convert JSON string to bytes
    - `bytes.Buffer` holds those bytes in memory, allowing you to pass the data to the HTTP request body
    - the buffer allows the request body to be treated as a readable stream of data
      - now HTTP client can read from buffer and send the content over the network in chunks
  - `json.Marshal(data)`: converts the Golang struct into a JSON-formatted string
- `bytes.NewBuffer(jsonData)`:converts JSON string into an instance of `bytes.Buffer`
  - `bytes.Buffer`: it's a Go type that provides an efficient way to manage data streams
  - buffer: in-memory queue where the bytes can be written to or read from
    - convenient way to hold and send the request body data

```go
func createUser(url, apiKey string, data User) (User, error) {

  // encode the data as json
  jsonData, err := json.Marshal(data)
  if err != nil {
    return User{}, err
  }

  // create new request
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
  if err != nil {
    return User{}, err
  }

  // set request headers
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-API-Key", apiKey)

  // create new client and make the request
  client := &http.Client{}
  res, err := client.Do(req)

  if err != nil {
    return User{}, err
  }
  defer res.Body.Close()

  var user User
  decoder := json.NewDecoder(res.Body)
  err = decoder.Decode(&user)
  if err != nil {
    return User{}, err
  }
  return user, nil
}
```

- `http.Response` has a `.StatusCode` property that contains the status code of the response

```go
func getUserCode(url string) int {
    res, err := http.Get(url)
    if err != nil {
        return 0
    }
    defer res.Body.Close()
    return res.StatusCode
}
```

- `http.StatusOk` is a constant that represents the status code 200
- `http.StatusCreated` is a constant that represents the status code 201
- `http.StatusNotFound` is a constant that represents the status code 404

#### API functions

[example of API functions](./code/golang/example1/http.go)

- unlike `GET` and `POST` there's no `http.Put` function
- for `PUT` requests, create raw `http.Request` that an http.Client can `myClient.Do(myRequest)`

##### PUT request implementation

1. get required data: endpoint URL, id, data that is being sent
1. convert data into json with `json.Marshal(data)`
1. create new HTTP request with `http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))`
1. set headers for request
1. create new `http.Client` and `client.Do(req)` to execute request
1. `defer res.Body.Close()` to close connection between client and server

TODO POST request implementation

##### server implementation

[example of server that doesn't serve files](./code/golang/example2/main.go)
[example of static file server](./code/golang/example3/main.go)

1. create http request multiplexer
1. program handler functions for different paths
1. start server and listen on port for requests

- `http.NewServeMux`: creates new http request multiplexer (or router)
  - multiplexer: device that combines multiple input signals into a single output signal
    - in the context of HTTP, a multiplexer is a router that directs incoming HTTP requests to the appropriate handler
    - server checks the `ServeMux` (`mux` in the example) to see if there's a handler for the path in the request
-`r *http.Request`: `r` is a pointer to an `http.Request` struct
  - `http.Request` struct contains all the information about the incoming HTTP request
- `mux.HandleFunc("/path", func())`: registers an anonymous handler function for the given path and responds to the request
- `mux.Handle("/", http.FileServer(http.Dir(".")))`: serves files from the current directory for root url `"/"`

##### handling endpoints

> endpoints are the URLs that your API exposes

- common operations:
  - GET resource: retrieve a resource (uses resource ID)
  - GET resources: retrieve a list of resources
  - POST resource: create a new resource
  - PUT resource: update an existing resource (uses resource ID)
  - DELETE resource: delete an existing resource (uses resource ID)

[example of server that serves files](./code/golang/example4/main.go)

- `http.Handler` is just an interface
- you need to strip the prefix from the file path before serving the file
  - `mux.Handle("/app/assets/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))`
    - removes the prefix `/app` from the request URL before passing it to the handler
      - after stripping the prefix, the request URL becomes `/assets/logo.jpeg`
    - serves files from `./assets` directory
- `http.FileServer(http.Dir("."))` serves files from the current directory
- it looks for files relative to `"."` (current directory)
- the remaining path `/assets/logo.jpeg` is used to find the file

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

- `ResponseWriter` interface is used to construct an HTTP response
  - `myResponseWriter.WriteHeader(http.StatusOK)`: sets the status code of the response
  - `myResponseWriter.Write([]byte(fmt.Sprintf("Hits: %d", hits)))`: writes the response body
- `Request` struct contains all the information about the incoming HTTP request

- `http.Handler` vs `http.HandlerFunc`:
  - `http.Handler` is any defined type that implements the interface above
  - use `http.Handler` for more complex use cases, that use routers, middleware, or custom logic
- `http.HandlerFunc` is a type that is a function with the signature `func(ResponseWriter, *Request)`
  - `http.HandlerFunc` is a type that implements the `http.Handler` interface
  - use `http.HandlerFunc` for simple use cases, like serving static files

##### middleware

[example of server with middleware](https://github.com/troclaux/chirpy/main.go)

request -> middleware -> handler -> response

1. use `atomic.Int32` to create a counter
1. middleware function does something before the handler
1. handler function does something with the request

- patterns: string that define which URL paths `ServeMux` should match to route HTTP requests
  - generally look like this `[METHOD] [HOST]/[PATH]`, where all parts are optional
    - e.g. `GET http://www.example.com/api/users`
    - METHOD: HTTP method (e.g. `GET`, `POST`, `PUT`, `DELETE`)
    - HOST: domain name or IP address (e.g. `http://www.example.com`)
    - PATH: path to the resource on the server (e.g. `/api/users`)
  - example
    - pattern: `/app/`
    - matches: `/app/`, `/app/img.png`, `/app/css/style.css`
    - doesn't match: `/app`
  - longer patterns take precedence over shorter patterns
    - example
      - pattern: `/` and `/app/`
      - request: `/app/img.png`
      - matched pattern: `/app/`, because it's longer than `/`
  - patterns can include a hostname
    - example
      - pattern: `www.example.com/about`
      - matches: `http://www.example.com/about`
      - doesn't match: `http://api.example.com/about`
  - think of patterns as signs on a highway:
    - a fixed path (e.g., `/about`) is like a specific exit
    - a subtree path (e.g., `/images/`) is like an entire district with many roads inside it
    - the longest match rule ensures you take the most specific road available

#### url

> package that parses URL

```go
parsedURL, err := url.Parse("https://cartoonwebsite.com/toons")
if err != nil {
  fmt.Println("error parsing url:", err)
  return
}

hostname := parsedURL.Hostname()
fmt.Println(hostname)     // cartoonwebsite.com
```

### enconding

#### json

[JSON](/javascript.md#json-javascript-object-notation)

[json.NewDecoder.Decode vs json.Unmarshal](https://stackoverflow.com/questions/21197239/decoding-json-using-json-unmarshal-vs-json-newdecoder-decode)

reading json data and storing it in golang struct:

1. declare struct that will store decoded json data
1. create json decoder
1. initialize empty struct defined in step 1
1. attempt to decode and read json data and write the decoded data into a golang struct
1. congratulations, you now have json data available in a golang variable if no errors occurred

[practical example](https://github.com/troclaux/chirpy/blob/main/handler_users_create.go):

```go
type User struct {
  ID        uuid.UUID `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Email     string    `json:"email"`
  Password  string    `json:"password"`
}

func (cfg *apiConfig) handleUsersCreate(w http.ResponseWriter, r *http.Request) {
  // creates new JSON decoder to read the request body
  decoder := json.NewDecoder(r.Body)
  // create empty User struct to store the decoded JSON
  reqUser := User{}
  // attempts to decode the JSON from the request body and store it in the reqUser struct
  if err := decoder.Decode(&reqUser); err != nil {
    log.Printf("error decoding user: %v", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
}
```

- `json.NewDecoder(input)`: creates new JSON decoder that will read from the input
- `json.NewDecoder(r.Body).Decode(&post)`: decodes the json data from the `decoder` into `post` struct
- `json.NewEncoder(w).Encode(User{key: value})`: encodes `User` struct as json and writes it to the http response `w`


- Capitalize first letter of each field, to allow `json` package to access each one
- add a tag before each field (e.g. `json:` to specify the name of the field when it is serialized to JSON
  - serialize: process of converting an object or data structure into a format that can be sent over a network


### string

### testing

> built-in package that provides essential tools to define and run tests

- test files should end with `_test.go`
  - e.g. `calculator.go` => `calculator_test.go`
- each test function should start with `Test` followed by a descriptive name
  - `func TestAdd(t *testing.T)`
- running tests: `go test`
  - this command looks for `_test.go` files
  - print code coverage: `go test . -cover`
- Go doesn't come with assertion libraries built-in
  - checks results with simple conditionals and `t.Errorf()` or `t.Fatalf()`

```go
func TestAdd(t *testing.T) {
  result := Add(2, 3)
  expected := 5
  if result != expected {
    t.Errorf("Add(2, 3) = %d; want %d", result, expected)
  }
}
```
- common pattern in go is to use table-driven tests
  - this means you define a set of inputs and expected outputs, then loop over them
  - makes it easy to add new cases or modify old ones

```go
func TestAdd(t *testing.T) {
  tests := []struct {
    name     string
    a, b     int
    expected int
  }{
    {"both positive", 2, 3, 5},
    {"one negative", -2, 3, 1},
    {"both negative", -2, -3, -5},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := Add(tt.a, tt.b)
      if result != tt.expected {
        t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
      }
    })
  }
}
```

### bufio
### io
### os
### time
### sql
### csv

### golang useful tools

- staticcheck: linting checks
  - install staticcheck: `go install honnef.co/go/tools/cmd/staticcheck@latest`
  - run staticcheck: `staticcheck ./...`
- gosec: security chacks
  - install gosec: `go install github.com/securego/gosec/v2/cmd/gosec@latest`
  - run gosec: `gosec ./...`

## Golang proverbs by Rob Pike

Don't communicate by sharing memory, share memory by communicating.

Concurrency is not parallelism.

Channels orchestrate; mutexes serialize.

The bigger the interface, the weaker the abstraction.

Make the zero value useful.

interface{} says nothing.

Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.

A little copying is better than a little dependency.

Syscall must always be guarded with build tags.

Cgo must always be guarded with build tags.

Cgo is not Go.

With the unsafe package there are no guarantees.

Clear is better than clever.

Reflection is never clear.

Errors are values.

Don't just check errors, handle them gracefully.

Design the architecture, name the components, document the details.

Documentation is for users.

Don't panic.
