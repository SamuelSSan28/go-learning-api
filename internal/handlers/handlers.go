package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"study-go/internal/models"

	"github.com/gorilla/mux"
)

// SetupRoutes configura todas as rotas da API
func SetupRoutes(r *mux.Router) {
	// Rotas de usuários
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	// Rota de health check
	r.HandleFunc("/health", HealthCheck).Methods("GET")
}

// GetUsers retorna todos os usuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{ID: 1, Name: "João Silva", Email: "joao@email.com"},
		{ID: 2, Name: "Maria Santos", Email: "maria@email.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser cria um novo usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simular criação (em um projeto real, salvaria no banco)
	user.ID = 3 // ID simulado

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser retorna um usuário específico
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Simular busca (em um projeto real, buscaria no banco)
	user := models.User{ID: id, Name: "Usuário " + strconv.Itoa(id), Email: "user" + strconv.Itoa(id) + "@email.com"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser atualiza um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser remove um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Simular remoção (em um projeto real, removeria do banco)
	w.WriteHeader(http.StatusNoContent)
}

// HealthCheck verifica se a API está funcionando
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "OK",
		"message": "API funcionando corretamente",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
} 