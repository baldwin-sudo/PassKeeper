package api

import (
	"backend/db"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(userService *db.UserServiceSql, passwordService *db.PasswordServiceSql) http.Handler {
	r := mux.NewRouter()

	services := &StorageServices{
		UserService:     userService,
		PasswordService: passwordService,
	}

	// Health check route
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// User routes
	r.HandleFunc("/users", services.createUser).Methods("POST")
	r.HandleFunc("/login/{username}", services.authenticateUser).Methods("POST")

	// Password routes
	r.HandleFunc("/passwords/create", services.createPassword).Methods("POST")
	//TODO MODIFY : REMOVE
	r.HandleFunc("/passwords/{user_id}", services.getAllPasswords).Methods("GET")
	r.HandleFunc("/passwords/{user_id}/{password_id}", services.updatePassword).Methods("PUT")
	r.HandleFunc("/passwords/{user_id}/{password_id}", services.deletePassword).Methods("DELETE")

	// Wrap the router with CORS to allow requests from any origin
	corsRouter := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Add frontend origin here
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	return corsRouter
}
