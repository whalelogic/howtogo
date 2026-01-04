---
author: Keith Thomson
categories:
- Programming
tags:
- go
- variables
- golang
- programming
title: Variables in Go
---
# Variables in Go {#variables-in-go .title}

In Go, variables are explicitly declared and used by the compiler to
check the type-correctness of function calls. This section covers how to
declare, initialize, and use variables in Go.

## Declaration and Initialization {#declaration-and-initialization .subtitle}

Variables can be declared using the `var` keyword, followed by the
variable name and type. You can also initialize the variable at the time
of declaration.

    package main

    import "fmt"

    func main() {
        // Declare a variable of type string
        var name string
        name = "Go"
        fmt.Println("Hello, " + name)

        // Declare and initialize a variable
        var version = 1.17
        fmt.Println("Version:", version)

        // Multiple variables can be declared at once
        var x, y int = 10, 20
        fmt.Println(x, y)
    }

## Short Variable Declarations {#short-variable-declarations .subtitle}

Inside a function, you can use the `:=` short assignment statement to
declare and initialize a variable. The type of the variable is inferred
from the value on the right-hand side.

    package main

    import "fmt"

    func main() {
        // Short variable declaration
        language := "Go"
        fmt.Println(language)

        // Multiple assignments
        a, b, c := 1, "hello", true
        fmt.Println(a, b, c)
    }

**Note:** Short variable declarations are only available inside
functions.

## Constants {#constants .subtitle}

Constants are declared like variables, but with the `const` keyword.
They can be character, string, boolean, or numeric values. Constants
cannot be declared using the `:=` syntax.

    package main

    import "fmt"

    const Pi = 3.14

    func main() {
        const World = "World"
        fmt.Println("Hello", World)
        fmt.Println("Happy", Pi, "Day")

        const Truth = true
        fmt.Println("Go rules?", Truth)
    }

## Basic Data Types {#basic-data-types .subtitle}

Go has a variety of basic data types, including:

- **bool:** A boolean value, either `true` or `false`.
- **string:** A sequence of bytes.
- **int, int8, int16, int32, int64:** Signed integers of various sizes.
- **uint, uint8, uint16, uint32, uint64, uintptr:** Unsigned integers of
  various sizes.
- **byte:** An alias for `uint8`.
- **rune:** An alias for `int32`, represents a Unicode code point.
- **float32, float64:** Floating-point numbers.
- **complex64, complex128:** Complex numbers.