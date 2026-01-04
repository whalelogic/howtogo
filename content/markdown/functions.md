---
author: Keith Thomson
categories:
- Programming
tags:
- go
- functions
- golang
- programming
title: 🔭 Functions in Go
---
# Functions in Go {#functions-in-go .title}

Functions are central to Go. They are the building blocks of a Go
program and are used to perform specific tasks. This section covers the
basics of creating and using functions in Go.

## Defining and Calling Functions {#defining-and-calling-functions .subtitle}

A function is defined with the `func` keyword, followed by the function
name, a list of parameters, the return type(s), and the function body.

    package main

    import "fmt"

    // A simple function with two parameters and one return value
    func add(x int, y int) int {
        return x + y
    }

    func main() {
        result := add(42, 13)
        fmt.Println("42 + 13 =", result)
    }

When you have multiple consecutive parameters of the same type, you can
omit the type from all but the last one.

    func add(x, y int) int {
        return x + y
    }

## Multiple Return Values {#multiple-return-values .subtitle}

A function in Go can return any number of results. This is often used to
return both a result and an error value.

    package main

    import (
        "fmt"
        "errors"
    )

    func divide(a, b float64) (float64, error) {
        if b == 0 {
            return 0, errors.New("division by zero")
        }
        return a / b, nil
    }

    func main() {
        result, err := divide(10, 2)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("Result:", result)
        }

        result, err = divide(10, 0)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("Result:", result)
        }
    }

## Variadic Functions {#variadic-functions .subtitle}

A function that can be called with a varying number of arguments is
called a variadic function. The last parameter of a variadic function is
of the form `...T`, which means it can accept any number of arguments of
type `T`.

    package main

    import "fmt"

    func sum(nums ...int) {
        fmt.Print(nums, " ")
        total := 0
        for _, num := range nums {
            total += num
        }
        fmt.Println(total)
    }

    func main() {
        sum(1, 2)
        sum(1, 2, 3)

        nums := []int{1, 2, 3, 4}
        sum(nums...)
    }

## Key Features of Functions {#key-features-of-functions .subtitle}

- **First-Class Citizens:** Functions are first-class citizens in Go,
  meaning they can be assigned to variables, passed as arguments to
  other functions, and returned from other functions.
- **Anonymous Functions:** You can create functions without a name,
  known as anonymous functions or function literals.
- **Closures:** An anonymous function can form a closure, which means it
  has access to the variables in its lexical scope.
