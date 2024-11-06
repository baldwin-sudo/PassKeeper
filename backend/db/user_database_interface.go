package db

import "backend/models"

// UserRepository provides an interface for user data access
type UserDatabaseInterface interface {
	Create(user *models.User) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetById(id int) (*models.User, error)
}
