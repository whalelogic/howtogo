---
author: Keith Thomson
categories:
- Programming
tags:
- go
- error-handling
- golang
- programming
title: 🩺 Error Handling in Go
---
# Error Handling in Go {#error-handling-in-go .title}

Go's approach to error handling is explicit and straightforward. Errors
are values that are returned from functions, not thrown as exceptions.

## The error Type {#the-error-type .subtitle}

The error type is a built-in interface:

    type error interface {
        Error() string
    }

    // Using errors
    import "errors"

    func divide(a, b float64) (float64, error) {
        if b == 0 {
            return 0, errors.New("division by zero")
        }
        return a / b, nil
    }

    func main() {
        result, err := divide(10, 0)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println("Result:", result)
    }

## Creating Custom Errors {#creating-custom-errors .subtitle}

Create custom error types by implementing the error interface:

    type ValidationError struct {
        Field string
        Message string
    }

    func (e *ValidationError) Error() string {
        return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
    }

    func validateAge(age int) error {
        if age < 0 {
            return &ValidationError{
                Field: "age",
                Message: "cannot be negative",
            }
        }
        return nil
    }

## Error Wrapping {#error-wrapping .subtitle}

Wrap errors to add context using fmt.Errorf with %w:

    import (
        "fmt"
        "errors"
    )

    func readConfig(filename string) error {
        _, err := os.Open(filename)
        if err != nil {
            return fmt.Errorf("failed to read config: %w", err)
        }
        return nil
    }

    func main() {
        err := readConfig("config.json")
        if err != nil {
            fmt.Println(err)
            
            // Unwrap to check original error
            if errors.Is(err, os.ErrNotExist) {
                fmt.Println("Config file not found")
            }
        }
    }

## Error Checking Patterns {#error-checking-patterns .subtitle}

Common patterns for checking specific errors:

    // errors.Is - check if error matches a value
    if errors.Is(err, os.ErrNotExist) {
        // Handle file not found
    }

    // errors.As - check if error is of a specific type
    var validationErr *ValidationError
    if errors.As(err, &validationErr) {
        fmt.Println("Field:", validationErr.Field)
    }

    // Type assertion
    if netErr, ok := err.(*net.Error); ok {
        if netErr.Timeout() {
            // Handle timeout
        }
    }

## Panic and Recover {#panic-and-recover .subtitle}

For unrecoverable errors, use panic and recover:

    func riskyOperation() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered from:", r)
            }
        }()
        
        // This will panic
        panic("something went wrong")
    }

    // Use panic sparingly, only for truly exceptional situations
    func mustConnect(addr string) *Connection {
        conn, err := connect(addr)
        if err != nil {
            panic(fmt.Sprintf("failed to connect to %s: %v", addr, err))
        }
        return conn
    }

## Best Practices {#best-practices .subtitle}

Guidelines for effective error handling:

    // 1. Always check errors
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer f.Close()

    // 2. Provide context with error messages
    if err != nil {
        return fmt.Errorf("failed to process user %d: %w", userID, err)
    }

    // 3. Handle errors at the appropriate level
    func processFile(filename string) error {
        // Handle specific errors here
        if err := validateFile(filename); err != nil {
            return err
        }
        
        // Let unexpected errors bubble up
        return processData(filename)
    }

    // 4. Use named return values for error handling
    func operation() (result int, err error) {
        defer func() {
            if err != nil {
                err = fmt.Errorf("operation failed: %w", err)
            }
        }()
        
        // Do work...
        return
    }
