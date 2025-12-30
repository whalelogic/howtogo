
# Concurrency in Go

Go provides built-in support for concurrent programming, making it easy to write programs that can handle multiple tasks simultaneously. Concurrency in Go is based on two main concepts: **goroutines** and **channels**.

## Goroutines

A goroutine is a lightweight thread of execution. You can start a new goroutine by using the `go` keyword before a function call. The function will then run concurrently with the calling function.

```
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
    go say("world")
    say("hello")
}
```

In this example, `go say("world")` starts a new goroutine. The `main` function continues to execute, and both functions run concurrently.

## Channels

Channels are a typed conduit through which you can send and receive values with the `<-` operator. Channels can be used to communicate between goroutines.

```
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
}
```

This example creates two goroutines to calculate the sum of two halves of a slice. The results are sent back to the `main` function through a channel.

## Select

The `select` statement lets a goroutine wait on multiple communication operations. A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```
package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

## Sync Package

The `sync` package provides basic synchronization primitives such as mutual exclusion locks. These are useful for protecting shared data from being accessed by multiple goroutines at the same time.

*   `sync.Mutex`: A mutual exclusion lock.
*   `sync.RWMutex`: A reader/writer mutual exclusion lock.
*   `sync.WaitGroup`: Waits for a collection of goroutines to finish.
*   `sync.Once`: An object that will perform exactly one action.
*   `sync.Cond`: A condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event.