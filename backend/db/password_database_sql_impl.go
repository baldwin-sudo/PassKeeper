package db

import (
	"backend/models"
	"errors"
)

// PasswordServiceSql provides methods for managing passwords
type PasswordServiceSql struct {
	dataSource DataSource // The data source to use (MySQL or SQLite)
}

// NewPasswordService creates a new PasswordService
func NewPasswordServiceSql(ds DataSource) *PasswordServiceSql {
	return &PasswordServiceSql{dataSource: ds}
}

// Create saves a new password entry in the database
func (ps *PasswordServiceSql) Create(password *models.Password) error {
	query := "INSERT INTO passwords (website, description, email, username, password, access_count, created_at, updated_at, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := ps.dataSource.Execute(query, password.Website, password.Description, password.Email, password.Username, password.EncryptedPassword, password.AccessCount, password.CreatedAt, password.UpdatedAt, password.UserID)
	return err
}

// GetByID retrieves a password entry by ID
func (ps *PasswordServiceSql) GetByID(id int) (*models.Password, error) {
	query := "SELECT id, website, description, email, username, password, access_count, created_at, updated_at, user_id FROM passwords WHERE id = ?"
	rows, err := ps.dataSource.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var password models.Password
		if err := rows.Scan(&password.ID, &password.Website, &password.Description, &password.Email, &password.Username, &password.EncryptedPassword, &password.AccessCount, &password.CreatedAt, &password.UpdatedAt, &password.UserID); err != nil {
			return nil, err
		}
		return &password, nil
	}

	return nil, errors.New("password not found")
}

// Update modifies an existing password entry in the database
func (ps *PasswordServiceSql) Update(password *models.Password) error {
	query := "UPDATE passwords SET website = ?, description = ?, email = ?, username = ?, password = ?, access_count = ?, updated_at = ? WHERE id = ?"
	_, err := ps.dataSource.Execute(query, password.Website, password.Description, password.Email, password.Username, password.EncryptedPassword, password.AccessCount, password.UpdatedAt, password.ID)
	return err
}

// Delete removes a password entry by ID
func (ps *PasswordServiceSql) Delete(id int) error {
	query := "DELETE FROM passwords WHERE id = ?"
	_, err := ps.dataSource.Execute(query, id)
	return err
}

// GetAllByUserID retrieves all password entries for a specific user
func (ps *PasswordServiceSql) GetAllByUserID(user models.User) ([]*models.Password, error) {
	query := "SELECT id, website, description, email, username, password, access_count, created_at, updated_at ,user_id FROM passwords WHERE user_id = ?"
	rows, err := ps.dataSource.Query(query, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwords []*models.Password
	for rows.Next() {
		var password models.Password
		if err := rows.Scan(&password.ID, &password.Website, &password.Description, &password.Email, &password.Username, &password.EncryptedPassword, &password.AccessCount, &password.CreatedAt, &password.UpdatedAt, &password.UserID); err != nil {
			return nil, err
		}
		encryption_key := models.DeriveKey(user.MasterPassword, user.Salt, models.NUM_ITER)
		PlainPasswordBytes, _ := models.Decrypt(password.EncryptedPassword, encryption_key)
		password.PlainPassword = string(PlainPasswordBytes)
		passwords = append(passwords, &password)
	}
	return passwords, nil
}
