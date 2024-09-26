
# golang

## properties

- has strong standard library
- syntax is similar to C
- doesn't use semicolons
- is not an object-oriented language
  - uses structs instead of classes
- doesn't have `while`, uses `for` instead
  - `for sum < 10 { sum += 1 }`

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

- `go mod init`: creates a new module
  - module: collection of packages
  - `package`: directory of related .go files
  - writes a new go.mod file in the current directory
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

- `package`: directory of related .go files
  - each package can only have 1 `main()` function
  - `package main`: contains `main()` function
- `import`: imports packages
  - `import "fmt"`: import package with basic formatting functions (includes Println)
  - when importing multiple packages, don't use comma between each package
- module: collection of packages

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
    - floating-point numbers: `float32`, `float64`
    - complex numbers: `complex64`, `complex128`
    - others:
      - `byte` (alias for `uint8`)
      - `rune` (alias for `int32`)
	- single-quoted strings (`''`) are used for runes
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
}

func main() {
    p := Person{"John"}
    v := p

    // Type assertion to get the Person struct
    person, ok := v.(Person)
    if ok {
        fmt.Println(person.name) // Output: John
    }

    // simpler example
    var x interface{} = "hello"
    value, isType := x.(string)
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

### enums

- `iota`: increments automatically each new constant in a block

```go
package main
import "fmt"

// Define an enumeration for weekdays
const (
    Sunday = iota  // iota starts at 0
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

- method signature: method name, receiver type and the parameter list
  - e.g. `func (p Person) sayHello()`
- interfaces don't require variables, only methods
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
  - `s2 := make([]int, 5, 10)`
    - length of s2 = 5
    - capacity of s2 = 10
  - slice's zero value is `nil`
- map: key-value pairs hash table
  - `m1 := map[string]int{ "apple":  1, "banana": 2, }`
  - `m2 := make(map[string]int, 10)`
    - capacity of m2 = 10
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
  - length: initial length of slice
  - capacity: number of items that the slice can hold without needing to be resized
    - capacity's default value is the length

> [!WARNING]
> `slice = append(slice, 1, 2, 3, 4)`: add items to end of slice
> `append()` changes the underlying array of its parameter AND returns a new slice
> ALWAYS ASSIGN THE RESULT OF APPEND() BACK TO THE SLICE TO ENSURE THE CHANGES ARE APPLIED CORRECTLY
> `append()` only works for slices
> don't do this: `someSlice = append(otherSlice, element)`
> you can accidentally declare multiple slices that share the same underlying array
> if this happens, when you change one slice, you can change other slices without knowing

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
- zero value of map is `nil`
- maps can contain maps, also called nested maps
- map keys can be of any type that is comparable
  - structs can be keys
    - `users := make(map[User]int)`

- delete element from map: `delete(map1, key)`
- checks if a key exists: `elem, ok := map1[key]`
  - if `key` is in `map1`, then `ok` is `true` and `elem` is the value as expected
  - if `key` is not in `map1`, then `ok` is `false` and `elem` is the zero value for the map's element type

> [!NOTE]
> maps return zero value when you try to access a value with a key that doesn't exist
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

> collection of related source files that are compiled together to form a single unit of code

- 1 or more .go files that normally are in the same directory (but not necessarily)
- each package can only have 1 `main()` function

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

### if else

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

### loops

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

## functions

- can return any number of results
- return values: define the return type of a function
  - located in the function signature, after the parameters
  - can be surrounded by parentheses (optional)

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
> you can define an anonymous functions function inside a function signature

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

## concurrency with goroutines and channels

go provides built-in support for concurrency using goroutines and channels

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

## concurrent programming

- goroutine: a lightweight thread managed by the go runtime
- `go`: TODO

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

## channels

Channels are used to communicate between goroutines

```go
package main

import "fmt"

func main() {
  ch := make(chan int)

  go func() {
    ch <- 42 // Send a value into the channel
  }()

  value := <-ch // Receive a value from the channel
  fmt.Println(value)
}
```

## pointers

> [!NOTE]
> golang doesn't have pointer arithmetic

- `&`: returns the memory address of its operand
  - address-of operator
- `*`: returns the value stored at the memory address
  - dereference operator

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

## errors

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

### string
### os
### testing
### time
### bufio
### sql
### json
### csv
### net
