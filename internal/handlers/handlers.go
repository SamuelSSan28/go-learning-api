package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"study-go/internal/models"

	"github.com/gorilla/mux"
)

// SetupRoutes registers all API routes
func SetupRoutes(r *mux.Router) {
	// User routes
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	// Health check route
	r.HandleFunc("/health", HealthCheck).Methods("GET")
}

// GetUsers returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Mary Smith", Email: "mary@example.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simulate creation (in a real project we would persist to a DB)
	user.ID = 3 // simulated ID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser returns a specific user
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	// Simulate lookup (in a real project we would query the DB)
	user := models.User{ID: id, Name: "User " + strconv.Itoa(id), Email: "user" + strconv.Itoa(id) + "@example.com"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
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

// DeleteUser removes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	// Simulate deletion (in a real project we would delete from the DB)
	w.WriteHeader(http.StatusNoContent)
}

// HealthCheck verifies API is up
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "OK",
		"message": "API is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
} 