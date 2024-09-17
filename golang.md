
# golang

## properties

- functions in go are first-class citizens and can be used as values
- statically typed: each variable must have a type
- compiled language
- has strong standard library
- syntax is similar to C
- no implicit type conversion
- doesn't have `while`, uses `for` instead

## go cli commands

- `go mod init`:
- `go fmt`:

## basics

- `package`: packages group related Go code together
  - `package main`: contains main func
- `import`: 
  - `import "fmt"`: import package with basic functions
    - 

```go

package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

```


## basic go commands

- `go build`: compiles the go package
- `go run`: compiles and runs go programs
- `go fmt`: formats go source code
- `go test`: runs unit tests
- `go get`: downloads and installs packages and dependencies
- `go mod`: manages dependencies (modules)



## Declaring Variables

```go
package main

import "fmt"

func main() {
    // Using var keyword
    var a int = 10
    var b = 20 // Type inference
    var c float64

    // Short variable declaration
    d := 30.5 // Type inferred as float64

    fmt.Println(a, b, c, d)
}
```

- `var a int = 10`: declares a variable `a` of type `int` and assigns it a value of `10`
- `var b = 20`: go can infer the type from the assigned value
- `var c float64`: declares a variable `c` of type `float64` with the zero value (0.0)
- `d := 30.5`: the shorthand `:=` declares and initializes a variable without explicitly specifying the type (it is inferred from the assigned value)

## basic data types

- Integers: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- Floating-point numbers: `float32`, `float64`
- Complex numbers: `complex64`, `complex128`
- Booleans: `bool`
- Strings: `string`
## type casting

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

## control structures

### if else

```go
package main

import "fmt"

func main() {
    x := 10

    if x > 5 {
        fmt.Println("x is greater than 5")
    } else if x == 5 {
        fmt.Println("x is equal to 5")
    } else {
        fmt.Println("x is less than 5")
    }
}
```

### for Loop

```go
package main

import "fmt"

func main() {
    // Basic for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // While-style loop
    x := 0
    for x < 5 {
        fmt.Println(x)
        x++
    }

    // Infinite loop
    // for {
    //     fmt.Println("Infinite loop")
    // }
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

## arrays, slices, and maps

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

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5} // Slice
    s = append(s, 6)          // Append an element
    fmt.Println(s)
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

## concurrency with goroutines and channels

Go provides built-in support for concurrency using goroutines and channels

### goroutines

A goroutine is a lightweight thread managed by the Go runtime

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
