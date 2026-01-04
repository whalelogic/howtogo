---
author: Keith Thomson
categories:
- Programming
tags:
- go
- structs
- golang
- programming
title: 🏛️Structs in Go
---
# Structs in Go {#structs-in-go .title}

A struct is a composite data type that groups together variables under a
single name. These variables, known as fields, can be of different
types. Structs are a fundamental concept in Go for creating complex data
structures.

## Defining and Using Structs {#defining-and-using-structs .subtitle}

You can define a struct using the `type` and `struct` keywords. Once a
struct type is defined, you can create instances (values) of that
struct.

    package main

    import "fmt"

    // Define a struct type
    type Person struct {
        Name string
        Age  int
    }

    func main() {
        // Create an instance of the Person struct
        p1 := Person{Name: "Alice", Age: 30}
        fmt.Println(p1)

        // Accessing struct fields
        fmt.Println("Name:", p1.Name)
        fmt.Println("Age:", p1.Age)

        // You can also create a pointer to a struct
        p2 := &Person{Name: "Bob", Age: 25}
        fmt.Println(p2.Name) // Go automatically dereferences the pointer
    }

## Methods on Structs {#methods-on-structs .subtitle}

Go allows you to define methods on struct types. A method is a function
with a special receiver argument. The receiver appears in its own
argument list between the `func` keyword and the method name.

    package main

    import "fmt"

    type Rectangle struct {
        Width, Height float64
    }

    // A method for the Rectangle struct
    func (r Rectangle) Area() float64 {
        return r.Width * r.Height
    }

    func main() {
        rect := Rectangle{Width: 10, Height: 5}
        fmt.Println("Area:", rect.Area())
    }

## Embedding Structs {#embedding-structs .subtitle}

Go does not provide inheritance in the traditional sense, but it does
support composition through struct embedding. You can embed a struct
within another struct to "inherit" its fields and methods.

    package main

    import "fmt"

    type Engine struct {
        Horsepower int
    }

    func (e Engine) Start() {
        fmt.Println("Engine started.")
    }

    type Car struct {
        Engine // Embed the Engine struct
        Make   string
        Model  string
    }

    func main() {
        car := Car{
            Engine: Engine{Horsepower: 300},
            Make:   "Toyota",
            Model:  "Camry",
        }

        // The Car struct now has the fields and methods of Engine
        fmt.Println("Horsepower:", car.Horsepower)
        car.Start()
    }

This is a powerful feature that allows for flexible and reusable code.

## Key Features of Structs {#key-features-of-structs .subtitle}

- **Composition:** Structs are the primary way to build complex data
  types in Go.
- **Methods:** You can define methods on structs to add behavior.
- **Embedding:** Structs can be embedded within other structs to promote
  code reuse.
- **Value vs. Pointer Receivers:** Methods can have either a value or a
  pointer receiver, which affects whether the method can modify the
  original struct.
