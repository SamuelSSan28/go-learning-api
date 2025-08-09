package main

// Imports are the packages we bring into this file
import (
	"log"        // write logs to stdout/stderr
	"net/http"   // HTTP server primitives
	"os"         // environment variables, files, etc.

	// our packages (local code)
	"study-go/internal/handlers"  // HTTP request handlers
	"study-go/pkg/middleware"     // request interceptors (middlewares)

	// third-party router
	"github.com/gorilla/mux"
)

func main() {
	// 1) CREATE ROUTER
	// Router directs incoming requests to the right handler
	r := mux.NewRouter()

	// 2) ATTACH MIDDLEWARES
	// Middlewares run before handlers for cross-cutting concerns
	r.Use(middleware.LoggingMiddleware)  // logs each request
	r.Use(middleware.CORSMiddleware)     // sets CORS headers

	// 3) REGISTER ROUTES
	// Centralized in handlers.SetupRoutes
	handlers.SetupRoutes(r)

	// 4) PICK SERVER PORT
	// Read from env var PORT or fallback to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 5) START THE SERVER
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
} 