package main

// Imports são como "bibliotecas" que importamos para usar no código
import (
	"log"        // Para registrar mensagens no console
	"net/http"   // Para criar servidor HTTP
	"os"         // Para acessar variáveis de ambiente

	// Nossos próprios pacotes (código que criamos)
	"study-go/internal/handlers"  // Funções que processam requisições HTTP
	"study-go/pkg/middleware"     // Interceptadores de requisições

	// Biblioteca externa para roteamento HTTP
	"github.com/gorilla/mux"
)

func main() {
	// 1. CRIAR O ROTEADOR
	// Um roteador é como um "guia" que direciona as requisições para as funções corretas
	r := mux.NewRouter()

	// 2. CONFIGURAR MIDDLEWARES
	// Middlewares são como "filtros" que processam cada requisição antes de chegar ao destino
	r.Use(middleware.LoggingMiddleware)  // Registra informações de cada requisição
	r.Use(middleware.CORSMiddleware)     // Permite requisições de outros sites

	// 3. CONFIGURAR ROTAS
	// Define quais URLs fazem o quê (ex: /api/users → lista usuários)
	handlers.SetupRoutes(r)

	// 4. CONFIGURAR PORTA DO SERVIDOR
	// Pega a porta do ambiente ou usa 8080 como padrão
	port := os.Getenv("PORT")  // Ex: se você definir PORT=3000, usa 3000
	if port == "" {
		port = "8080"  // Se não definir, usa 8080
	}

	// 5. INICIAR O SERVIDOR
	// Mostra mensagem de que está iniciando
	log.Printf("Servidor iniciando na porta %s", port)
	
	// Inicia o servidor e fica esperando requisições
	// Se der erro, para a aplicação e mostra o erro
	log.Fatal(http.ListenAndServe(":"+port, r))
} 