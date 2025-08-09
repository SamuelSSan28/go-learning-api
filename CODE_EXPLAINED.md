# Code Explained (Deep Dive)

This document expands on the core ideas shown in README with extra context and small examples.

## Imports

```go
import (
  "log"        // logging to stdout/stderr
  "net/http"   // HTTP server and client
  "os"         // env vars, files

  // local packages (our code)
  "study-go/internal/handlers"
  "study-go/pkg/middleware"

  // third‑party
  "github.com/gorilla/mux"
)
```
- Standard library: `log`, `net/http`, `os`
- Local: `study-go/...`
- Third‑party: `github.com/gorilla/mux`

## Router (no classes in Go)
Go uses packages, functions, and structs (not classes). `mux.NewRouter()` constructs a router:
```go
r := mux.NewRouter()
```
You register routes and middlewares on `r` and pass it to `http.ListenAndServe`.

## Middlewares
Middlewares run before your handler and can modify the request/response.
```go
r.Use(middleware.LoggingMiddleware)  // logs method, path, duration
r.Use(middleware.CORSMiddleware)     // sets CORS headers, handles OPTIONS
```

## Routes
```go
r.HandleFunc("/api/users", GetUsers).Methods("GET")
r.HandleFunc("/health", HealthCheck).Methods("GET")
```
A route matches method + path and invokes the handler function.

## Short variable declaration (:=)
Declares and initializes a variable in one step (function scope only):
```go
port := os.Getenv("PORT") // declare + assign
if port == "" {
  port = "8080" // re-assign
}
```
`var` is used when you want a zero value or explicit type:
```go
var n int       // 0 by default
var s = "go"   // inferred type string
x := 42         // short declaration
```

## Starting the server
```go
log.Printf("Server starting on %s", port)
log.Fatal(http.ListenAndServe(":"+port, r))
```
- `ListenAndServe` blocks and serves requests with router `r`
- `log.Fatal` logs any startup error and exits

## Request Flow
```
Client → CORS middleware → Logging middleware → Router → Handler → JSON response
```

## Handler example (simplified)
```go
func GetUsers(w http.ResponseWriter, r *http.Request) {
  users := []User{{ID: 1, Name: "John"}}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(users)
}
```

## Why this layout
- Separation of concerns (handlers, models, middleware)
- Scales as features grow
- Easy to test and refactor

For the full project flow and runnable code, see README. 