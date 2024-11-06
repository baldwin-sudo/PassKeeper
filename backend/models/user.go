package models

import (
	"time"
)

const SALT_SIZE int = 16

// User represents a user in the system with a hashed master password and salt
type User struct {
	ID                   int       `json:"-"`
	Username             string    `json:"username"`
	MasterPassword       string    `json:"master_password"`
	HashedMasterPassword string    `json:"-"`
	Salt                 []byte    `json:"-"`
	CreatedAt            time.Time `json:"-"`
}

func NewUser(username string, password string) (*User, error) {
	salt, err := GenerateSalt(SALT_SIZE)
	if err != nil {
		return nil, err // Return an error if salt generation fails
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err // Return an error if password hashing fails
	}

	return &User{
		Username:             username,
		Salt:                 salt,
		MasterPassword:       password,
		HashedMasterPassword: hashedPassword,
		CreatedAt:            time.Now(),
	}, nil // Return the new User instance and nil error
}
func (user *User) CheckPassword(password string) bool {

	return CheckPassword(user.HashedMasterPassword, password)
}
func (user *User) SetMasterPassword(master_password string) {
	user.MasterPassword = master_password
}
