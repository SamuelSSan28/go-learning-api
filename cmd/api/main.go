package main

import (
	"log"
	"net/http"
	"os"

	"study-go/internal/handlers"
	"study-go/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Configurar roteador
	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSMiddleware)

	// Rotas
	handlers.SetupRoutes(r)

	// Porta do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
} 