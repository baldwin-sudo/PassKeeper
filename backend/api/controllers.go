package api

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const SESSION_NAME = "user-session"

// Handler to create a new user
func (router *Services) createUser(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"master_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := models.NewUser(userInput.Username, userInput.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	user, err = router.UserService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Handler to authenticate a user
func (router *Services) authenticateUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	var requestBody struct {
		MasterPassword string `json:"master_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := router.UserService.GetByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !user.CheckPassword(requestBody.MasterPassword) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	session, err := router.SessionStore.Get(r, SESSION_NAME)
	fmt.Println("SESSSION ID :", session.ID)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}
	fmt.Println("user id", user.ID)

	session.Values["user_id"] = user.ID
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "failed to save session", http.StatusInternalServerError)
		return
	}
	fmt.Println("auth session", session.Values)
	user.MasterPassword = "YOU CAN'T SEE THIS DATA"
	json.NewEncoder(w).Encode(user)
}

// Handler to create a new password entry
func (router *Services) createPassword(w http.ResponseWriter, r *http.Request) {
	session, err := router.SessionStore.Get(r, SESSION_NAME)
	if err != nil {
		http.Error(w, "Couldn't access session", http.StatusInternalServerError)
		return
	}
	// fmt.Println("SESSSION ID :", session.ID)
	// fmt.Println("create password session", session.Values)

	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "User ID not found in session", http.StatusUnauthorized)
		return
	}

	var reqBody struct {
		Website       string `json:"website"`
		Description   string `json:"description"`
		Email         string `json:"email"`
		Username      string `json:"username"`
		PlainPassword string `json:"plain_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := router.UserService.GetById(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	password := models.NewPassword(reqBody.Website, reqBody.Description, reqBody.Email, reqBody.Username, reqBody.PlainPassword, 0, *user)

	if err := router.PasswordService.Create(password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(password)
}

// getPassword handles the GET request to retrieve a password by ID.
func (router *Services) getAllPasswords(w http.ResponseWriter, r *http.Request) {
	// Get session and check if there's an error
	session, err := router.SessionStore.Get(r, SESSION_NAME)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Couldn't access session", http.StatusInternalServerError)
		return
	}

	// Retrieve user ID from session, check if it's valid
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Invalid session or user ID", http.StatusUnauthorized)
		return
	}

	fmt.Println("User ID:", userID)

	// Fetch the user from the UserService
	user, err := router.UserService.GetById(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("User with ID %d not found", userID), http.StatusNotFound)
		return
	}

	// Retrieve all passwords for the user
	passwords, err := router.PasswordService.GetAllByUserID(*user)
	fmt.Println("passwords", passwords)
	fmt.Println("error", err)
	if err != nil {
		http.Error(w, fmt.Sprintf("Couldn't fetch passwords for user ID %d", userID), http.StatusInternalServerError)
		return
	}

	// Send the passwords in the response with a 200 OK status
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(passwords); err != nil {
		http.Error(w, "Error encoding passwords to JSON", http.StatusInternalServerError)
	}
}

// updatePassword handles the PUT request to update a password by ID.
func (router *Services) updatePassword(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here...
}

// deletePassword handles the DELETE request to remove a password by ID.
func (router *Services) deletePassword(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here...
}
