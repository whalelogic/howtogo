# 🧵 Channels & Mutexes 🔒

### A Go Reference & Guide

A thorough reference and practical guide for using channels and mutexes in Go. This document collects syntax, keywords, type/method tables, idioms, patterns, pitfalls, recipes, and advanced topics with examples and recommendation checklists.

Contents
- Table of Contents
- Overview
- Channels (syntax, semantics, examples, patterns)
- Channel Quick Reference Table
- Channel Patterns & Recipes
- Mutexes (sync.Mutex, sync.RWMutex, sync.Cond)
- Mutex Quick Reference Table
- Mutex Patterns & Recipes
- Channels vs Mutexes — When to use what
- Concurrency Gotchas, Debugging & Performance
- FAQs & Checklist
- Further reading

---

## 📚 Table of Contents
- [Overview](#-overview)
- [Channels — Concepts & Syntax](#-channels--concepts--syntax)
  - [Creation & Types](#creation--types)
  - [Send / Receive / Comma-ok](#send--receive--comma-ok)
  - [Directional channel types](#directional-channel-types)
  - [Buffered vs Unbuffered](#buffered-vs-unbuffered)
  - [Closing channels](#closing-channels)
  - [Range over channel](#range-over-channel)
  - [select and non-blocking ops](#select-and-non-blocking-ops)
  - [Nil channels & semantics](#nil-channels--semantics)
  - [Channel Examples](#channel-examples)
- [Channel Quick Reference Table](#-channel-quick-reference-table)
- [Channel Patterns & Recipes](#-channel-patterns--recipes)
  - [Producer-consumer / Pipeline](#producer-consumer--pipeline)
  - [Worker pool / Fan-out fan-in](#worker-pool--fan-out-fan-in)
  - [Bounded buffer with context cancellation](#bounded-buffer-with-context-cancellation)
  - [Broadcast via close](#broadcast-via-close)
- [Mutexes — Concepts & Syntax](#-mutexes--concepts--syntax)
  - [sync.Mutex](#syncmutex)
  - [sync.RWMutex](#syncrwmutex)
  - [sync.Cond](#synccond)
  - [Best practices & common pitfalls](#best-practices--common-pitfalls)
  - [Mutex Examples](#mutex-examples)
- [Mutex Quick Reference Table](#-mutex-quick-reference-table)
- [Mutex Patterns & Recipes](#-mutex-patterns--recipes)
  - [Thread-safe counters & maps](#thread-safe-counters--maps)
  - [Read-mostly caches with RWMutex](#read-mostly-caches-with-rwmutex)
  - [Condition variable example](#condition-variable-example)
- [Channels vs Mutexes — Choosing & Hybrids](#-channels-vs-mutexes--choosing--hybrids)
- [Concurrency Gotchas & Debugging](#-concurrency-gotchas--debugging)
- [Performance & Alternatives](#-performance--alternatives)
- [FAQs & Checklist](#-faqs--checklist)
- [Further reading](#-further-reading)

---

## 🔎 Overview

Go provides two broad primitives to coordinate concurrent goroutines and manage access to shared state:

- Channels: typed pipes to send/receive values between goroutines. They embody message-passing and are excellent for pipelines, worker pools, and ownership transfer.
- Mutexes: low-level locks (sync.Mutex, sync.RWMutex) for synchronized access to shared in-memory state.

This guide explains syntax, semantics, idiomatic usage, and patterns for both.

---

## 🧭 Channels — Concepts & Syntax

Channels let goroutines communicate safely and synchronize without explicit locks.

### Creation & Types

- make(chan T) — unbuffered (synchronous)
- make(chan T, n) — buffered (capacity n)
- Zero value is nil: var ch chan T

Examples:
```go
ch := make(chan int)      // unbuffered
buf := make(chan string, 10) // buffered with capacity 10
var nilch chan int        // nil channel; blocks on recv/send
```

### Send / Receive / Comma-ok

- Send: ch <- v
- Receive: v := <-ch
- Discard receive: <-ch
- Comma-ok: v, ok := <-ch  // ok == false when channel closed and drained

Example:
```go
ch <- 42            // send
x := <-ch           // receive
v, ok := <-ch       // ok false if closed
```

### Directional channel types

- Bidirectional: chan T
- Send-only: chan<- T
- Receive-only: <-chan T

Function signatures:
```go
func producer(out chan<- int) { out <- 1 }
func consumer(in <-chan int)  { fmt.Println(<-in) }
```

### Buffered vs Unbuffered

- Unbuffered: send and receive must synchronize; useful for handoff.
- Buffered: send succeeds immediately until buffer fills; decouples sender/receiver timing.

### Closing channels

- close(ch) signals no more values will be sent.
- Receivers read remaining values; subsequent reads yield zero values with ok==false.
- Only the sender should close; sending on closed channel panics; closing twice panics.

```go
close(ch)
v, ok := <-ch // ok == false if drained
```

### Range over channel

Range receives until ch is closed and drained:

```go
for v := range ch {
    fmt.Println(v)
}
```

### select and non-blocking ops

select multiplexes channel operations:

```go
select {
case v := <-ch1:
    // handle v
case ch2 <- x:
    // sent x
case <-time.After(time.Second):
    // timeout
default:
    // non-blocking fallback
}
```

Non-blocking send/receive idiom:
```go
select {
case ch <- v:
    // sent
default:
    // would block
}
```

### Nil channels & semantics

- A nil channel blocks forever on send and receive.
- Useful for dynamic select: set case channel to nil to disable it.

```go
var ch chan int
// <-ch // blocks forever
```

### Panics & gotchas

- Sending on closed channel => panic
- Closing closed channel => panic
- Closing nil channel => panic
- Range without close (no sender closes) => deadlock if reader expects termination
- len(cap) are snapshot-ish and race-prone; avoid relying on them for synchronization

---

## 🧪 Channel Examples

Simple handoff:
```go
ch := make(chan string)
go func() { ch <- "hello" }()
fmt.Println(<-ch) // prints "hello"
```

Buffered:
```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
// no blocking yet
fmt.Println(<-ch, <-ch)
```

Timeout with select:
```go
select {
case v := <-ch:
    fmt.Println("got", v)
case <-time.After(500 * time.Millisecond):
    fmt.Println("timeout")
}
```

Producer-consumer pipeline:
```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n*n
        }
    }()
    return out
}
```

---

## 📋 Channel Quick Reference Table

| Concept / Function | Syntax | Notes |
|---|---:|---|
| Create unbuffered | `make(chan T)` | Blocks on send until recv ready |
| Create buffered | `make(chan T, n)` | Buffer capacity n |
| Send | `ch <- v` | Blocks if unbuffered or buffer full |
| Receive | `v := <-ch` | Blocks if empty |
| Receive discard | `<-ch` | Receive and drop value |
| Comma-ok | `v, ok := <-ch` | ok==false when closed and drained |
| Close | `close(ch)` | Panic if closed twice or nil |
| Range | `for v := range ch {}` | Stops when closed and drained |
| Directional types | `chan<- T`, `<-chan T` | Use in function signatures |
| len/cap | `len(ch)`, `cap(ch)` | Snapshot; not synchronization tools |
| Nil channel | `var ch chan T` | Blocks forever on ops |
| select | `select { case ... }` | Multiplexing, default for non-blocking |

---

## 🔧 Channel Patterns & Recipes

### Producer-consumer / Pipeline

- Use chaining functions that return receive-only channels.
- Always close the channel from the producer when done.
- Keep pipelines cancellable using context.Context.

Example:
```go
func generator(ctx context.Context, nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            select {
            case <-ctx.Done():
                return
            case out <- n:
            }
        }
    }()
    return out
}
```

### Worker pool (fan-out / fan-in)

- Fan-out: multiple workers read from same input channel.
- Fan-in: aggregate worker results into a single output channel.

Sketch:
```go
func worker(id int, jobs <-chan Job, results chan<- Result) {
    for j := range jobs {
        // process
        results <- result
    }
}

jobs := make(chan Job)
results := make(chan Result)
for w := 0; w < nW; w++ {
    go worker(w, jobs, results)
}
// send jobs then close jobs
```

Use sync.WaitGroup to close results after all workers finish:
```go
var wg sync.WaitGroup
for i := 0; i < n; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // worker loop
    }()
}
go func() {
    wg.Wait()
    close(results)
}()
```

### Bounded buffer + cancellation

Combine buffered channels with context to build cancelable queues.

### Broadcast via close

- For one-time broadcast (no more values), send values into a channel, then close; receivers do range and exit when drained.
- To broadcast the same value to many receivers, use a separate channel per receiver or use a "fan-out" goroutine that duplicates the stream.

---

## 🔒 Mutexes — Concepts & Syntax

Mutexes are for protecting shared memory when multiple goroutines access or mutate it.

### sync.Mutex

- Type: sync.Mutex
- Methods:
  - mu.Lock()
  - mu.Unlock()
- Zero value is an unlocked Mutex.

Idioms:
```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
// critical section
```

Rules:
- Do not copy a Mutex after first use (copying breaks it).
- Locks are not reentrant — Lock twice in the same goroutine deadlocks.

### sync.RWMutex

- Type: sync.RWMutex
- Methods:
  - mu.RLock()
  - mu.RUnlock()
  - mu.Lock()
  - mu.Unlock()

Use RLock for read-only access when many readers and few writers.

Pitfalls:
- Avoid upgrading from RLock to Lock while holding RLock — deadlock risk.
- Prefer simple Mutex if RW complexity doesn't help.

### sync.Cond

- Use for wait/notify semantics with an associated Locker (often *sync.Mutex).
- Create: cond := sync.NewCond(&mu)
- Methods:
  - cond.Wait() — atomically release lock and wait; reacquire lock before returning.
  - cond.Signal() — wake one waiter.
  - cond.Broadcast() — wake all waiters.

Pattern:
```go
mu := &sync.Mutex{}
cond := sync.NewCond(mu)
mu.Lock()
for !ready {
    cond.Wait()
}
mu.Unlock()
```

### Best practices & common pitfalls

- Always Unlock with defer immediately after Lock.
- Keep critical sections small and avoid blocking operations while holding a mutex (I/O, network, time.Sleep).
- Never copy Mutex/RWMutex once used (store in struct, pass pointer to struct if necessary).
- Use RWMutex when reads greatly outnumber writes and contention matters.
- Prefer channels for coordination (signaling, ownership) and mutexes for protecting shared state.

---

## 🧪 Mutex Examples

Counter:
```go
type Counter struct {
    mu sync.Mutex
    n  int
}
func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}
func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.n
}
```

Read-mostly cache with RWMutex:
```go
type SafeCache struct {
    mu sync.RWMutex
    m  map[string]string
}
func (c *SafeCache) Get(k string) (string, bool) {
    c.mu.RLock()
    v, ok := c.m[k]
    c.mu.RUnlock()
    return v, ok
}
func (c *SafeCache) Set(k, v string) {
    c.mu.Lock()
    c.m[k] = v
    c.mu.Unlock()
}
```

Condition variable (producer/consumer):
```go
type Queue struct {
    mu    sync.Mutex
    cond  *sync.Cond
    items []int
}
func NewQueue() *Queue {
    q := &Queue{}
    q.cond = sync.NewCond(&q.mu)
    return q
}
func (q *Queue) Push(x int) {
    q.mu.Lock()
    q.items = append(q.items, x)
    q.mu.Unlock()
    q.cond.Signal()
}
func (q *Queue) Pop() int {
    q.mu.Lock()
    for len(q.items) == 0 {
        q.cond.Wait()
    }
    x := q.items[0]
    q.items = q.items[1:]
    q.mu.Unlock()
    return x
}
```

---

## 📋 Mutex Quick Reference Table

| Type | Methods / Ops | Zero-value usable? | Notes |
|---|---:|:---:|---|
| sync.Mutex | Lock(), Unlock() | Yes | Not reentrant; don't copy |
| sync.RWMutex | RLock(), RUnlock(), Lock(), Unlock() | Yes | Readers/writers semantics |
| sync.Cond | Wait(), Signal(), Broadcast() | N/A (needs Locker) | Use with a Locker (often *sync.Mutex) |

Common patterns:
- mu.Lock(); defer mu.Unlock()
- mu.RLock(); defer mu.RUnlock()

---

## 🛠️ Mutex Patterns & Recipes

### Thread-safe map / cache

Go's standard maps are not safe for concurrent writes. Use a mutex:

```go
type SafeMap struct {
    mu sync.Mutex
    m  map[string]int
}
func (s *SafeMap) Load(k string) (int, bool) {
    s.mu.Lock()
    v, ok := s.m[k]
    s.mu.Unlock()
    return v, ok
}
func (s *SafeMap) Store(k string, v int) {
    s.mu.Lock()
    s.m[k] = v
    s.mu.Unlock()
}
```

For high read concurrency, use sync.RWMutex or sync.Map (concurrent map in stdlib for specific patterns).

### Double-checked locking (careful!)

Avoid naïve double-checked locking. Use sync.Once or channels.

Use sync.Once for lazy initialization:
```go
var once sync.Once
var cfg *Config
func GetConfig() *Config {
    once.Do(func() { cfg = loadConfig() })
    return cfg
}
```

### Condition variable example

See queue example above.

---

## ⚖️ Channels vs Mutexes — Choosing & Hybrids

When to use channels:
- You can model the problem as passing ownership or messages between goroutines.
- Pipelines, worker pools, timeouts, throttling, bounded queues.
- Prefer when you want to avoid shared mutable state.

When to use mutexes:
- You have shared memory (counters, caches, maps) with fine-grained reads/writes.
- Performance-sensitive sections where channel overhead is too high.
- Complex structures where explicit locking is simpler.

Hybrid:
- Use channels for control flow and mutexes for protecting in-memory caches behind the scenes.
- Use channels to serialize access (single goroutine actor model) if suitable: one goroutine owns the state and services request messages on a channel (avoids mutexes).

Example actor pattern:
```go
type request struct{ key string; resp chan<- int }
func actor(reqs <-chan request) {
    store := make(map[string]int)
    for r := range reqs {
        r.resp <- store[r.key]
    }
}
```

---

## ⚠️ Concurrency Gotchas & Debugging

- Race conditions: use `go test -race` or `go run -race` to detect data races.
- Deadlocks: common causes include waiting on channels that no goroutine will write to, or double-locking mutexes.
- Panic on send to closed channel or closing already closed channel.
- Blocking on nil channels causes goroutine hang.
- Avoid long blocking operations while holding a mutex.

Debugging tips:
- Reproduce with race detector.
- Add logging with goroutine IDs (use third-party libraries or attach request IDs).
- Break problems into smaller testable components.
- Use context.Context for cancellation propagation.

---

## 🚀 Performance & Alternatives

- Channels are slightly heavier than raw mutex when used for extremely hot-path synchronization. For counters and simple state, consider `sync/atomic`.
- `sync.Map` is optimized for certain read-heavy concurrent maps.
- `sync.Pool` for object pooling.
- Keep locks small and critical sections short. Avoid I/O in critical section.
- For very high throughput, microbench with `go test -bench` and `-benchmem` to compare channels vs mutexes vs atomics.

---

## ✅ FAQs & Checklist

Q: Who should close a channel?  
A: The sender. Close signals no more values.

Q: Can I read from a closed channel?  
A: Yes; receives return zero value and ok==false.

Q: Is Mutex zero value usable?  
A: Yes.

Q: Should I use defer Unlock always?  
A: Usually yes — it ensures unlock on all return paths. In super-hot code, sometimes explicit Unlock is used to avoid defer cost; measure first.

Q: How to avoid deadlock while upgrading RLock to Lock?  
A: Don't upgrade. Instead, release RLock and then Lock (with potential retry), or use a different design.

Checklist before shipping concurrent code:
- [ ] Run tests with -race
- [ ] Add timeouts / context where blocking can occur
- [ ] Keep critical sections short
- [ ] Avoid copying Mutex-containing structs
- [ ] Ensure all senders close channels or use other termination signals
- [ ] Use sync.Once for one-time initialization

---

## 📖 Further reading

- Official Go blog — [Go Concurrency Patterns](https://go.dev/blog/pipelines)  
- sync package docs — [sync](https://pkg.go.dev/sync)  
- Effective Go — [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)  
- Go Memory Model — [The Go Memory Model](https://go.dev/ref/mem)

---

