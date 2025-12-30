## `os` Package

The `os` package provides a platform-independent interface to core operating-system features. It enables file operations, environment variable management, directory manipulation, process interaction, and access to standard I/O.

This guide serves as a structured reference, including tables, examples, and usage explanations.

---

## Quick Reference Table

| Package | Name       | Kind      | Signature / Type                                               | Description                                       |
|---------|------------|-----------|----------------------------------------------------------------|---------------------------------------------------|
| os      | Create     | Function  | `func Create(name string) (*File, error)`                      | Creates or truncates a named file.                |
| os      | Open       | Function  | `func Open(name string) (*File, error)`                        | Opens a file for reading.                         |
| os      | OpenFile   | Function  | `func OpenFile(name string, flag int, perm FileMode) (*File, error)` | Opens a file with specific flags and permissions. |
| os      | ReadFile   | Function  | `func ReadFile(name string) ([]byte, error)`                   | Reads an entire file into a byte slice.           |
| os      | WriteFile  | Function  | `func WriteFile(name string, data []byte, perm FileMode) error` | Writes data to a file.                            |
| os      | Mkdir      | Function  | `func Mkdir(name string, perm FileMode) error`                 | Creates a new directory.                          |
| os      | MkdirAll   | Function  | `func MkdirAll(path string, perm FileMode) error`              | Recursively creates directories.                  |
| os      | Remove     | Function  | `func Remove(name string) error`                               | Removes a file or empty directory.                |
| os      | RemoveAll  | Function  | `func RemoveAll(path string) error`                            | Recursively removes a directory and contents.     |
| os      | Rename     | Function  | `func Rename(oldpath, newpath string) error`                   | Renames a file or directory.                      |
| os      | Stat       | Function  | `func Stat(name string) (FileInfo, error)`                     | Retrieves file metadata.                          |
| os      | IsNotExist | Function  | `func IsNotExist(err error) bool`                              | Checks if an error indicates a missing file.      |
| os      | Getenv     | Function  | `func Getenv(key string) string`                               | Retrieves an environment variable.                |
| os      | LookupEnv  | Function  | `func LookupEnv(key string) (string, bool)`                    | Retrieves an env variable + existence flag.       |
| os      | Setenv     | Function  | `func Setenv(key, value string) error`                         | Sets an environment variable.                     |
| os      | ExpandEnv  | Function  | `func ExpandEnv(s string) string`                              | Expands `$vars` in a string.                      |
| os      | Getwd      | Function  | `func Getwd() (string, error)`                                 | Returns working directory.                        |
| os      | Chdir      | Function  | `func Chdir(dir string) error`                                 | Changes working directory.                        |
| os      | Exit       | Function  | `func Exit(code int)`                                          | Exits program immediately.                        |
| os      | Args       | Variable  | `[]string`                                                     | Command-line arguments.                           |
| os      | Stdin      | Variable  | `*File`                                                        | Standard input descriptor.                        |
| os      | Stdout     | Variable  | `*File`                                                        | Standard output descriptor.                       |
| os      | Stderr     | Variable  | `*File`                                                        | Standard error descriptor.                        |
| os      | File       | Type      | `struct{ ... }`                                                | Represents an open file descriptor.               |
| os      | FileInfo   | Interface | `interface { Name(); Size()... }`                              | Metadata about files.                             |

---

# ✏️ 1. File Creation & Opening

### `os.Create()`

```go
f, err := os.Create("notes.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

f.WriteString("Hello, World!")
```

### `os.Open()`

```go
f, err := os.Open("notes.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

### `os.OpenFile()`

```go
f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer f.Close()

f.WriteString("New log entry\n")
```

---

# 📄 2. Reading & Writing Files 

### `os.ReadFile()`

```go
data, err := os.ReadFile("config.json")
if err != nil {
    log.Fatal(err)
}
println(string(data))
```

### `os.WriteFile()`

```go
content := []byte("This replaces the file content.")
err := os.WriteFile("output.txt", content, 0644)
if err != nil {
    log.Fatal(err)
}
```

---

# 📁 3. Directory Management

### Create directory

```go
err := os.Mkdir("data", 0755)
```

### Recursive create

```go
os.MkdirAll("logs/2023/jan", 0755)
```

### Delete items

```go
os.Remove("old_notes.txt")
os.RemoveAll("logs")
```

---

# 📊 4. File Metadata

```go
info, err := os.Stat("image.png")
if err != nil {
    if os.IsNotExist(err) {
        println("File does not exist")
    }
} else {
    println(info.Name(), info.Size())
}
```

### Rename:

```go
os.Rename("temp.tmp", "final.jpg")
```

---

# 🌐 5. Environment Variables

### Read

```go
path := os.Getenv("PATH")
```

### Lookup

```go
val, ok := os.LookupEnv("MY_API_KEY")
```

### Set

```go
os.Setenv("APP_MODE", "production")
```

---

# ⚙️ 6. Process & User Info

```go
dir, _ := os.Getwd()
os.Chdir("/tmp")
os.Exit(1)
```

---

# 📌 7. Standard Streams

```go
os.Stdout.WriteString("This goes to stdout\n")
os.Stderr.WriteString("This is an error message\n")
```

---
