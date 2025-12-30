---
author: Keith Thomson
categories:
- Programming
tags:
- go
- interfaces
- golang
- programming
title: Interfaces in Go
---
# Interfaces in Go {#interfaces-in-go .title}

An interface in Go is a type that specifies a set of method signatures.
A type implements an interface by implementing its methods. There is no
explicit declaration of intent, no "implements" keyword.

## Defining and Implementing Interfaces {#defining-and-implementing-interfaces .subtitle}

You can define an interface using the `type` and `interface` keywords.
Any type that has all the methods of the interface is said to implement
that interface.

    package main

    import (
        "fmt"
        "math"
    )

    // Define an interface
    type Shape interface {
        Area() float64
    }

    // Define a struct
    type Rectangle struct {
        Width, Height float64
    }

    // Implement the Area method for Rectangle
    func (r Rectangle) Area() float64 {
        return r.Width * r.Height
    }

    // Define another struct
    type Circle struct {
        Radius float64
    }

    // Implement the Area method for Circle
    func (c Circle) Area() float64 {
        return math.Pi * c.Radius * c.Radius
    }

    // A function that takes an interface type
    func PrintArea(s Shape) {
        fmt.Println("Area:", s.Area())
    }

    func main() {
        rect := Rectangle{Width: 10, Height: 5}
        circ := Circle{Radius: 7}

        PrintArea(rect)
        PrintArea(circ)
    }

In this example, both `Rectangle` and `Circle` implement the `Shape`
interface because they both have an `Area()` method.

## The Empty Interface {#the-empty-interface .subtitle}

The interface type that specifies zero methods is known as the **empty
interface**, written as `interface{}`. An empty interface may hold
values of any type, because every type has zero or more methods.

    package main

    import "fmt"

    func describe(i interface{}) {
        fmt.Printf("(%v, %T)\n", i, i)
    }

    func main() {
        var i interface{}
        describe(i)

        i = 42
        describe(i)

        i = "hello"
        describe(i)
    }

The empty interface is often used to handle values of unknown type.

## Type Assertions {#type-assertions .subtitle}

A type assertion provides access to an interface value\'s underlying
concrete value. A type assertion takes the form `t := i.(T)`, where `i`
is an interface value and `T` is the asserted type.

    package main

    import "fmt"

    func main() {
        var i interface{} = "hello"

        s := i.(string)
        fmt.Println(s)

        s, ok := i.(string)
        fmt.Println(s, ok)

        f, ok := i.(float64)
        fmt.Println(f, ok)

        // This will panic because i does not hold a float64
        // f = i.(float64)
        // fmt.Println(f)
    }

## Key Features of Interfaces {#key-features-of-interfaces .subtitle}

- **Implicit Implementation:** A type implements an interface simply by
  possessing all the methods the interface requires.
- **Polymorphism:** Interfaces allow you to write functions that can
  work with multiple types.
- **Decoupling:** Interfaces help to decouple different parts of your
  code, making it more modular and maintainable.