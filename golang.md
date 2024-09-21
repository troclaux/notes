
# golang

## properties

- has strong standard library
- syntax is similar to C
- doesn't use semicolons
- uses structs instead of classes
- doesn't have `while`, uses `for` instead
    - `for sum < 10 { sum += 1 }`

- statically typed
- compiled language
- has garbage collection
- functions are first-class citizens
- hoisted declarations: variables and functions can be used before they are declared
- no implicit type casting
- mutable strings
- multiplatform

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

## data types

- basic types:
  - numeric types:
    - integers: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
    - floating-point numbers: `float32`, `float64`
    - complex numbers: `complex64`, `complex128`
    - others:
      - `byte` (alias for `uint8`)
      - `rune` (alias for `int32`)
  - Booleans: `bool`
  - Strings: `string`
- composite types:
  - `array`
  - `slice`
  - `struct`
  - `map`
  - `chan`
- reference types:
  - `pointer`
  - `function`
- interface types:
  - `interface`

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

## data structures / composite types

- array: ordered collection of elements with fixed sizing
  - `var arr [5]int`
- slice: ordered collection of elements with dynamic resizing
  - `s1 := []int{1, 2, 3}`
  - `s2 := make([]int, 5, 10)`
    - length of s2 = 5
    - capacity of s2 = 10
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

- `make(type, size)`: function that can be used to create slices, maps and channels
  - type of data structure to create
  - size of data structure to create (can be length or capacity)

> [!NOTE]
> go doesn't have built-in stacks and queues, but they can be implemented using slices

### stack

```go
stack := []int{}
stack = append(stack, 1) // Push
stack = append(stack, 2) // Push
stack = append(stack, 3) // Push

top := stack[len(stack)-1]  // Peek
stack = stack[:len(stack)-1] // Pop
fmt.Println(top)            // 3
```

### queue

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
- methods and constructors aren't defined inside the struct initialization
- `new()`: creates a new instance of a struct
  - `p := new(Person)`
  - `p.Name = John`

```go

type Person struct {
  Name string
  Age  int
}

// creating object
p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name)  // Output: Alice
fmt.Println(p.Age)   // Output: 30

// creating a method
func (p *Person) Greet() {
  fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// calling method
p.Greet()  // Output: Hello, my name is Alice and I am 30 years old.

// constructor
func NewPerson(name string, age int) *Person {
  return &Person{Name: name, Age: age}
}

// using constructor
p := NewPerson("Charlie", 25)

```

### arrays

```go
package main

import "fmt"

func main() {
  var arr [5]int // Array of 5 integers
  arr[0] = 1
  fmt.Println(arr)
}
```

### slices

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

## maps

```go
package main

import "fmt"

func main() {
  m := make(map[string]int)
  m["foo"] = 1
  m["bar"] = 2
  fmt.Println(m)
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

## string operations

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
- you can define the return type of a function with **return values**
  - in the function initialization, after the parameters
  - can be surrounded by parentheses

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

## Error Handling

go does not have exceptions: instead, errors are handled using the built-in `error` type

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

## standard libraries

- library: collection of packages

### fmt

> provides functions for formatting and printing data

- format string: string that contains placeholders for the arguments, separated by commas
  - `fmt.Println("Hello, %s!", "world")`
  - `fmt.Printf("Name: %s, Age: %d, Height: %.2f meters\n", name, age, height)`

- `Print()`: prints arguments to the standard output
- `Println()`: prints arguments to the standard output, followed by a newline character
- `Printf()`: prints arguments to the standard output and allows the use of format strings

- `Sprint()`: returns a string representation of its arguments, without printing
- `Sprintf()`: similar to `Sprint()`, but allows the use of format strings

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

### math

The `math` package provides basic constants and mathematical functions for floating-point operations.

### os
### testing
### time
### bufio
### sql
### json
### csv
### net
