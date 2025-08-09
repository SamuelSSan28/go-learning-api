# Go Study API – Central Guide

A simple REST API in Go designed as a learning path. This README is the central, didactic guide for the repository.

- Project goal: learn the Go toolchain, project layout, HTTP server, routing, middleware, and JSON handling
- Tech: Go 1.22, net/http, Gorilla Mux

## Repository Map

```
study-go/
├── cmd/api/main.go              # Application entry point
├── internal/
│   ├── handlers/handlers.go     # HTTP handlers (users, health)
│   └── models/user.go           # Data models
├── pkg/middleware/middleware.go # Middlewares (logging, CORS)
├── docs/                        # Auto-generated Swagger docs
├── go.mod / go.sum              # Dependencies
├── README.md                    # Central, didactic guide (this file)
├── LICENSE                      # MIT license
└── .gitignore                   # Git ignores
```

## Quickstart

1) Install deps
```bash
go mod tidy
```
2) Run the API
```bash
go run cmd/api/main.go
```
3) Test
```bash
curl http://localhost:8080/health
```

## Core Concepts (with small code excerpts)

### 1) Imports in Go
Imports bring packages (standard library, third‑party, or your own) into the file.
```go
import (
  "log"        // logging to stdout/stderr
  "net/http"   // HTTP server and client
  "os"         // environment variables, files, etc.

  // our code
  "study-go/internal/handlers"
  "study-go/pkg/middleware"

  // third-party router
  "github.com/gorilla/mux"
)
```
- "log", "net/http", "os" come from the Go standard library
- "study-go/..." are local project packages
- "github.com/gorilla/mux" is a third‑party dependency (declared in go.mod)

### 2) Router: mux.NewRouter()
Go does not have classes; it has packages, functions, and structs. `mux.NewRouter()` is a function from Gorilla Mux that returns a `*mux.Router` struct instance you use to declare routes.
```go
r := mux.NewRouter()
```
You then attach middlewares and routes to this router and pass it to the HTTP server.

### 3) Middlewares
Middlewares are request interceptors that run before your handler.
```go
r.Use(middleware.LoggingMiddleware)
r.Use(middleware.CORSMiddleware)
```
- LoggingMiddleware: logs method, path, remote addr, duration
- CORSMiddleware: sets CORS headers and replies to OPTIONS

### 4) Routes
Routes map HTTP methods and paths to handlers.
```go
handlers.SetupRoutes(r) // centralizes all route bindings
```
Example (from handlers.go):
```go
r.HandleFunc("/api/users", GetUsers).Methods("GET")
r.HandleFunc("/health",   HealthCheck).Methods("GET")
```

### 5) The short variable declaration operator :=
`:=` declares and initializes a new variable in a single step (function scope only).
```go
port := os.Getenv("PORT") // declare 'port' and assign its value
if port == "" {
  port = "8080" // re-assign (already declared above)
}
```
- Use `:=` inside functions to both declare and assign
- Use `var` for zero‑values, wider scope, or when you need an explicit type
```go
var count int      // 0 by default
var name = "Go"   // type inferred
value := 42        // short declaration (function scope)
```

### 6) Starting the HTTP server
```go
log.Printf("Server starting on port %s", port)
log.Fatal(http.ListenAndServe(":"+port, r))
```
- `http.ListenAndServe` blocks and serves requests using your router
- `log.Fatal` logs and exits if the server fails to start

### 7) Request flow (high level)
```
Client → CORS middleware → Logging middleware → Router → Handler → JSON response
```

## API Endpoints

- GET `/health` → health check
- GET `/api/users` → list users
- POST `/api/users` → create user
- GET `/api/users/{id}` → get user by id
- PUT `/api/users/{id}` → update user
- DELETE `/api/users/{id}` → delete user

Quick tests
```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/users
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

## Add Swagger/OpenAPI (Step-by-step)

Swagger gives you interactive API docs at `/swagger/index.html`.

1) Install the CLI (once):
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
Ensure your PATH includes `$(go env GOPATH)/bin`.

2) Generate docs (creates `./docs`):
```bash
swag init --parseDependency --parseInternal -g cmd/api/main.go
```

3) Run the server and open:
- `http://localhost:8080/swagger/index.html`

Notes:
- Minimal annotations were added to handlers (e.g., `@Summary`, `@Tags`, `@Router`)
- If build breaks due to swag versions, pin it: `go get github.com/swaggo/swag@v1.16.6`

## Learning Checklist

- Imports: standard vs. local vs. external
- Router: create (`mux.NewRouter`), attach middlewares, register routes
- Middlewares: cross‑cutting concerns (logging, CORS)
- Handlers: read request, write JSON response, status codes
- JSON: `encoding/json` for encoding/decoding
- Variables: `:=` vs. `var`
- Project layout: `cmd/`, `internal/`, `pkg/`
- Swagger: annotations, generate docs, serve UI

## Next Steps

- Add database (PostgreSQL/MySQL) and repositories
- Add validation and authentication (JWT)
- Add tests (unit/integration)
- Add Swagger/OpenAPI docs for all endpoints
- Containerize with Docker

## License

MIT (see LICENSE) 