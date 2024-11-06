package db

import "backend/models"

// PasswordRepository provides an interface for password data access
type PasswordDatabaseInterface interface {
	Create(password *models.Password) error
	GetByID(id int) (*models.Password, error)
	Update(password *models.Password) error
	Delete(id int) error
	GetAllByUserID(userID int) ([]*models.Password, error)
}
