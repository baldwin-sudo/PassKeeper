package db

import (
	"backend/models"
	"errors"
	"fmt"
)

// UserServiceSql provides methods for user management
type UserServiceSql struct {
	dataSource DataSource // The data source to use (MySQL or SQLite)
}

// NewUserService creates a new UserService
func NewUserServiceSql(ds DataSource) *UserServiceSql {
	return &UserServiceSql{dataSource: ds}
}

// Create saves a new user in the database and retrieves its ID
func (us *UserServiceSql) Create(user *models.User) (*models.User, error) {
	// Insert the new user into the database
	query := "INSERT INTO users (username, hashed_master_password, salt, created_at) VALUES (?, ?, ?, ?)"
	result, err := us.dataSource.Execute(query, user.Username, user.HashedMasterPassword, user.Salt, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	// Retrieve the ID of the newly created user using LastInsertId()
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Set the user ID after insertion
	user.ID = int(userID)

	return user, nil
}

// GetByUsername retrieves a user by username
func (us *UserServiceSql) GetByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, hashed_master_password, salt, created_at FROM users WHERE username = ?"
	rows, err := us.dataSource.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.HashedMasterPassword, &user.Salt, &user.CreatedAt); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, errors.New("user not found")
}

// GetById retrieves a user by ID from the database
func (us *UserServiceSql) GetById(id int) (*models.User, error) {
	// Prepare the query to fetch user by ID
	query := "SELECT id, username, hashed_master_password, salt, created_at FROM users WHERE id = ?"

	// Execute the query
	rows, err := us.dataSource.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Check if we got any rows
	if !rows.Next() {
		// No user found with the given ID
		return nil, fmt.Errorf("user not found with id: %d", id)
	}

	// Define a user object to store the result
	var user models.User

	// Scan the row into the user object
	err = rows.Scan(&user.ID, &user.Username, &user.HashedMasterPassword, &user.Salt, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Ensure UserService implements UserServiceInterface
var _ UserDatabaseInterface = (*UserServiceSql)(nil)
