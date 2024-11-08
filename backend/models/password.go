package models

import (
	"time"
)

const NUM_ITER int = 32

// Password represents a password entry in the database
type Password struct {
	ID                int       `json:"password_id"` // Unique identifier for the password
	Website           string    `json:"website"`     // The website associated with the password
	Description       string    `json:"description"` // Optional description of the password entry
	Email             string    `json:"email"`       // Email associated with the account
	Username          string    `json:"username"`    // Username for the account
	PlainPassword     string    `json:"plain_password"`
	EncryptedPassword string    // Encrypted password
	AccessCount       int       `json:"access_count"` // Number of times the password has been accessed
	CreatedAt         time.Time // Timestamp when the password entry was created
	UpdatedAt         time.Time // Timestamp when the password entry was last updated
	UserID            int       // Foreign key reference to the user who owns the password
}

// NewPassword creates a new Password instance with the provided values.
// It sets the CreatedAt and UpdatedAt fields to the current time and initializes the AccessCount to 0.
func NewPassword(website, description, email, username, password string, accessCount int, user User) *Password {
	encryption_key := DeriveKey(user.MasterPassword, user.Salt, NUM_ITER)
	encrypted_password, _ := Encrypt([]byte(password), encryption_key)

	return &Password{
		Website:           website,
		Description:       description,
		Email:             email,
		Username:          username,
		PlainPassword:     password,
		EncryptedPassword: string(encrypted_password), // This should be the encrypted password
		AccessCount:       0,                          // Default to 0
		CreatedAt:         time.Now(),                 // Set created time to now
		UpdatedAt:         time.Now(),                 // Set updated time to now
		UserID:            user.ID,
	}
}
func (password *Password) IncrementAccessCount() {
	password.AccessCount += 1
}
