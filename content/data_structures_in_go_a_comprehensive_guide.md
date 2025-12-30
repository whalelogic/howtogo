---
title: Data Structures in Go: A Comprehensive Guide
author: Keith Thomson
description: Go (Golang) provides a rich set of built-in and library-> supported data structures that make it powerful for both systems programming and application development.
tags: [go, golang, data-structures, programming, guide]
---


#### **Go** (Golang) provides a rich set of built-in and library-> supported data structures that make it powerful for both systems programming and application development. 

In this guide, we’ll explore the core data structures available in Go, explain how they work, and show practical code examples.  


## 🔢 Arrays  

An **array** in Go is a fixed-size, ordered collection of elements.  

```go
package main

import "fmt"

func main() {
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr)

    // Iterate
    for i, v := range arr {
        fmt.Printf("Index %d = %d\n", i, v)
    }
}
```

- Arrays have a fixed length.  
- Useful when the size is known and constant.  



## 📐 Slices  

A **slice** is a dynamically-sized, flexible view into an array. Slices are the most commonly used data structure in Go.  

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    slice = append(slice, 6)
    fmt.Println(slice)
    fmt.Println("Length:", len(slice), "Capacity:", cap(slice))
}
```

- Built on top of arrays.  
- Support dynamic resizing with `append`.  
- Preferred over arrays in most cases.  



## 🗺 Maps  

A **map** is Go’s built-in hash table implementation for key-value pairs.  

```go
package main

import "fmt"

func main() {
    m := map[string]int{
        "Alice": 25,
        "Bob":   30,
    }
    m["Charlie"] = 35

    for k, v := range m {
        fmt.Printf("%s is %d years old\n", k, v)
    }
}
```

- Keys must be comparable (e.g., strings, ints).  
- Lookups are O(1) on average.  



## 🏗 Structs  

A **struct** groups fields together, making it Go’s way of creating custom data types.  

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name, "is", p.Age)
}
```

- Structs are used for modeling entities.  
- They are the building blocks of more complex data structures.  


## 🔗 Linked Lists  

Go’s standard library provides a **doubly linked list** via `container/list`.  

```go
package main

import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    l.PushBack(1)
    l.PushBack(2)
    l.PushFront(0)

    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}
```

- Each element points to the next and previous nodes.  
- Efficient for insertions/removals in the middle.  



## 📚 Stacks  

A **stack** is a LIFO (Last In, First Out) structure. Implemented easily with slices.  

```go
package main

import "fmt"

type Stack []int

func (s *Stack) Push(v int) {
    *s = append(*s, v)
}

func (s *Stack) Pop() int {
    if len(*s) == 0 {
        panic("stack is empty")
    }
    val := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return val
}

func main() {
    var s Stack
    s.Push(10)
    s.Push(20)
    fmt.Println(s.Pop()) // 20
}
```

- Built using slices.  
- Great for recursion-like problems, parsing, and backtracking.  

---

## 📬 Queues  

A **queue** is a FIFO (First In, First Out) structure. Also implemented with slices.  

```go
package main

import "fmt"

type Queue []int

func (q *Queue) Enqueue(v int) {
    *q = append(*q, v)
}

func (q *Queue) Dequeue() int {
    if len(*q) == 0 {
        panic("queue is empty")
    }
    val := (*q)[0]
    *q = (*q)[1:]
    return val
}

func main() {
    var q Queue
    q.Enqueue(1)
    q.Enqueue(2)
    fmt.Println(q.Dequeue()) // 1
}
```

- Useful for scheduling and breadth-first search (BFS).  



## ⛰ Heaps & Priority Queues  

Go provides heap operations in the `container/heap` package.  

```go
package main

import (
    "container/heap"
    "fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    h := &IntHeap{3, 1, 4}
    heap.Init(h)
    heap.Push(h, 2)
    fmt.Println(heap.Pop(h)) // 1 (smallest element)
}
```

- Implements a min-heap by default.  
- Can be adapted into a priority queue.  

---


## 🏁 Conclusion  

Go provides both **high-level abstractions** (slices, maps, structs) and **low-level control** (linked lists, heaps).  
By mastering these data structures, you’ll be ready to build efficient algorithms, design scalable applications, and handle both systems and business logic effectively.  

💡 *Next Step:* Try implementing algorithms like BFS, DFS, and Dijkstra’s algorithm using these data structures to strengthen your understanding.
