# Go `net/http` Package Comprehensive Reference

The **`net/http`** package provides HTTP client and server implementations.

## Server-Side Reference

| Type/Function | Signature | Description |
| :--- | :--- | :--- |
| **Handler** | `interface { ServeHTTP(ResponseWriter, *Request) }` | The primary interface for HTTP logic. |
| **HandlerFunc** | `type HandlerFunc func(ResponseWriter, *Request)` | An adapter allowing functions to act as Handlers. |
| **ListenAndServe** | `func ListenAndServe(addr string, h Handler) error` | Starts a server on the given address. |
| **ResponseWriter** | `interface { Header(); Write(); WriteHeader() }` | Used to construct the HTTP response. |
| **ServeMux** | `struct` | An HTTP request multiplexer (router). |

## Client-Side Reference

| Type/Function | Signature | Description |
| :--- | :--- | :--- |
| **Client** | `struct { Transport RoundTripper; Timeout time.Duration }` | An HTTP client for making requests. |
| **Get** | `func Get(url string) (*Response, error)` | Helper for simple GET requests. |
| **Post** | `func Post(url, contentType string, body Reader) (*Response, error)` | Helper for POST requests. |
| **NewRequest** | `func NewRequest(method, url string, body Reader) (*Request, error)` | Creates a complex request object. |

---

## Essential Structs

### http.Request
- `Method string`: GET, POST, etc.
- `URL *url.URL`: The target URL.
- `Header Header`: Map of headers.
- `Body io.ReadCloser`: The request body (must be closed by server).

### http.Response
- `Status string`: e.g., "200 OK".
- `StatusCode int`: e.g., 200.
- `Body io.ReadCloser`: The response body (must be closed by client).

---

## Practical Examples

### 1. Simple HTTP Server
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
```
#### 2. HTTP Client with Context and Timeout

```go
package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "[https://api.example.com](https://api.example.com)", nil)
	
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
