package api

import (
	"backend/db"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Services holds the services and the session store
type Services struct {
	UserService     *db.UserServiceSql
	PasswordService *db.PasswordServiceSql
	SessionStore    *sessions.CookieStore
}

// NewRouter initializes the router with routes and middleware
func NewRouter(userService *db.UserServiceSql, passwordService *db.PasswordServiceSql) http.Handler {
	r := mux.NewRouter()

	// Load session key from environment variables for bmust etter security
	sessionKey := "session-key"
	if sessionKey == "" {
		// You can use a fallback value or panic if the key is not set (ensure it's set in production)
		panic("SESSION_ENCRYPTION_KEY is not set")
	}

	// Secure the session store with more options
	store := sessions.NewCookieStore([]byte(sessionKey))
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,  // Ensure the cookie is not accessible via JavaScript
		Secure:   false, // Ensure cookies are sent only over HTTPS
		MaxAge:   3600,
	}

	services := &Services{
		UserService:     userService,
		PasswordService: passwordService,
		SessionStore:    store, // Use the store in services

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
	r.HandleFunc("/passwords", services.getAllPasswords).Methods("GET")
	r.HandleFunc("/passwords/{password_id}", services.updatePassword).Methods("PUT")
	r.HandleFunc("/passwords/{password_id}", services.deletePassword).Methods("DELETE")

	// Wrap the router with CORS to allow requests from any origin
	corsRouter := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Add frontend origin here
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)(r)

	return corsRouter
}
