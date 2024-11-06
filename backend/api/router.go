package api

import (
	"backend/db"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(userService *db.UserServiceSql, passwordService *db.PasswordServiceSql) *mux.Router {
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
	r.HandleFunc("/users/{username}", services.authenticateUser).Methods("POST")

	// Password routes
	r.HandleFunc("/passwords/{user_id}", services.createPassword).Methods("POST")
	r.HandleFunc("/passwords/{user_id}", services.getAllPasswords).Methods("GET")
	r.HandleFunc("/passwords/{user_id}/{password_id}", services.updatePassword).Methods("PUT")
	r.HandleFunc("/passwords/{user_id}/{password_id}", services.deletePassword).Methods("DELETE")

	return r
}
