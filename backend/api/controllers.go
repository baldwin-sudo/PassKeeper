package api

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// todo store a key in os.env
var store = sessions.NewCookieStore([]byte("session-encryption-key"))

type StorageServices struct {
	UserService     *db.UserServiceSql
	PasswordService *db.PasswordServiceSql
}

// Handler to create a new user
func (router *StorageServices) createUser(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`

		Password string `json:"master_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	user, err := models.NewUser(userInput.Username, userInput.Password)
	fmt.Println("created user", user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Use the UserService to create the user
	user, err = router.UserService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Handler to get a user by username
// Handler to get a user by username and check the password
func (router *StorageServices) authenticateUser(w http.ResponseWriter, r *http.Request) {
	// Retrieve the username from URL path parameters
	username := mux.Vars(r)["username"]

	// Decode the JSON body to get the password
	var requestBody struct {
		MasterPassword string `json:"master_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database
	user, err := router.UserService.GetByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Verify the password (compare the stored hash and salt with the provided password)
	if !user.CheckPassword(requestBody.MasterPassword) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	// Set up a new session
	session, err := store.Get(r, "user-session")
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// Store user information in the session (only store the user ID or other non-sensitive information)
	session.Values["master_password"] = user.MasterPassword // Save the username or any other info you'd like
	session.Save(r, w)
	// If the password matches, respond with the user details (excluding sensitive information)
	user, _ = models.NewUser(user.Username, user.MasterPassword)
	// user.MasterPassword = "" // Make sure not to send sensitive data like the password
	user.MasterPassword = "YOU CAN'T SEE THIS DATA"

	json.NewEncoder(w).Encode(user)
}

// Handler to create a new password entry
func (router *StorageServices) createPassword(w http.ResponseWriter, r *http.Request) {
	// Retrieve the id from URL path parameters
	userid := mux.Vars(r)["user_id"]
	// Retrieve the session from the request
	session, err := store.Get(r, "user-session")

	if err != nil {
		http.Error(w, "Couldnt access cookies ", http.StatusInternalServerError)
		return
	}
	var req_body struct {
		Website       string `json:"website"`     // The website associated with the password
		Description   string `json:"description"` // Optional description of the password entry
		Email         string `json:"email"`       // Email associated with the account
		Username      string `json:"username"`    // Username for the account
		PlainPassword string `json:"plain_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req_body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	userid_int, _ := strconv.Atoi(userid)
	user, _ := router.UserService.GetById(userid_int)
	// Check if the username is stored in the session
	master_password, _ := session.Values["master_password"].(string)
	user.MasterPassword = master_password
	password := models.NewPassword(req_body.Website, req_body.Website, req_body.Email, req_body.Username, req_body.PlainPassword, 0, *user)

	// Use the PasswordService to create the password
	if err := router.PasswordService.Create(password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(password)
}

// getPassword handles the GET request to retrieve a password by ID.
func (router *StorageServices) getAllPasswords(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here...
}

// updatePassword handles the PUT request to update a password by ID.
func (router *StorageServices) updatePassword(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here...
}

// deletePassword handles the DELETE request to remove a password by ID.
func (router *StorageServices) deletePassword(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here...
}
