# Go `io` Package Reference

The `io` package provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package **`os`**, into shared public interfaces that abstract the functionality.

## Core Interfaces

| Name | Signature | Description |
| :--- | :--- | :--- |
| **Reader** | `Read(p []byte) (n int, err error)` | Reads up to len(p) bytes into p. Returns bytes read and error (EOF at end). |
| **Writer** | `Write(p []byte) (n int, err error)` | Writes len(p) bytes from p to the underlying data stream. |
| **Closer** | `Close() error` | Standard interface for closing a stream, releasing resources. |
| **Seeker** | `Seek(offset int64, whence int) (int64, error)` | Sets the offset for the next Read or Write. |
| **ReadWriter** | `interface { Reader; Writer }` | Groups the basic Read and Write methods. |

---

## Utility Functions & Methods

| Function | Signature | Usage Example |
| :--- | :--- | :--- |
| **Copy** | `Copy(dst Writer, src Reader) (int64, error)` | `io.Copy(os.Stdout, strings.NewReader("Hello"))` |
| **ReadAll** | `ReadAll(r Reader) ([]byte, error)` | `data, err := io.ReadAll(response.Body)` |
| **WriteString** | `WriteString(w Writer, s string) (int, error)` | `io.WriteString(f, "Write this string to file")` |
| **LimitReader** | `LimitReader(r Reader, n int64) Reader` | `lr := io.LimitReader(largeReader, 1024)` |
| **MultiReader** | `MultiReader(readers ...Reader) Reader` | `combined := io.MultiReader(r1, r2, r3)` |
| **TeeReader** | `TeeReader(r Reader, w Writer) Reader` | `tr := io.TeeReader(r, os.Stdout) // Reads and prints simultaneously` |

---

## Practical Examples

### 1. Copying a File to Standard Output
```go
package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("This is a data stream.\n")
	// io.Copy is memory efficient as it uses a buffer internally
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		panic(err)
	}
}
```
2. Using a **LimitReader** to prevent memory exhaustion
Go
```go
func handleRequest(r io.Reader) {
	// Only allow reading 1MB of data
	limitedReader := io.LimitReader(r, 1024*1024)
	data, _ := io.ReadAll(limitedReader)
	_ = data
}
```

---

### 2. The `net/http` Package Reference
This file covers both **Client** and **Server** implementations, including the **`ResponseWriter`** and **`Request`** structs.



