package db

import (
	"backend/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMySQLDataSource_Connect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a database connection", err)
	}
	defer db.Close()

	mysqlDS := &MySQLDataSource{db: db}
	err = mysqlDS.Connect("user:password@/dbname")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Ensure that the database is properly closed
	if err := mysqlDS.Close(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Expect no queries
	mock.ExpectClose()
}

func TestSQLiteDataSource_Connect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a database connection", err)
	}
	defer db.Close()

	sqliteDS := &SQLiteDataSource{db: db}
	err = sqliteDS.Connect("file::memory:?cache=shared")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Ensure that the database is properly closed
	if err := sqliteDS.Close(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Expect no queries
	mock.ExpectClose()
}

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a database connection", err)
	}
	defer db.Close()

	userRepo := NewUserServiceSql(&SQLiteDataSource{db: db})

	user := &models.User{
		Username:             "testuser",
		HashedMasterPassword: "hashedpassword",
		Salt:                 []byte("somesalt"),
		CreatedAt:            time.Now(),
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Username, user.HashedMasterPassword, user.Salt, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if _, err := userRepo.Create(user); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a database connection", err)
	}
	defer db.Close()

	userRepo := NewUserServiceSql(&MySQLDataSource{db: db})

	expectedUser := &models.User{
		ID:                   1,
		Username:             "testuser",
		HashedMasterPassword: "hashedpassword",
		Salt:                 []byte("somesalt"),
		CreatedAt:            time.Now(),
	}

	mock.ExpectQuery("SELECT id, username, hashed_master_password, salt, created_at").
		WithArgs(expectedUser.Username).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "hashed_master_password", "salt", "created_at"}).
			AddRow(expectedUser.ID, expectedUser.Username, expectedUser.HashedMasterPassword, expectedUser.Salt, expectedUser.CreatedAt))

	user, err := userRepo.GetByUsername("testuser")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if user.Username != expectedUser.Username {
		t.Errorf("expected username %s, got %s", expectedUser.Username, user.Username)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Mock User for testing
func mockUser() models.User {
	return models.User{
		ID:                   1,
		Username:             "testuser",
		HashedMasterPassword: "hashedpassword",   // Assume this is already hashed
		Salt:                 []byte("somesalt"), // Mock salt
	}
}

// TestNewPassword tests the NewPassword function for creating a new Password instance
func TestNewPassword(t *testing.T) {
	user := mockUser()
	website := "example.com"
	description := "Test account"
	email := "test@example.com"
	username := "testuser"
	password := "securepassword"
	accessCount := 0

	pass := models.NewPassword(website, description, email, username, password, accessCount, user)

	if pass.Website != website {
		t.Errorf("expected website to be %s, got %s", website, pass.Website)
	}
	if pass.Description != description {
		t.Errorf("expected description to be %s, got %s", description, pass.Description)
	}
	if pass.Email != email {
		t.Errorf("expected email to be %s, got %s", email, pass.Email)
	}
	if pass.Username != username {
		t.Errorf("expected username to be %s, got %s", username, pass.Username)
	}
	if pass.PlainPassword != "" {
		t.Error("expected PlainPassword to be empty")
	}
	if pass.AccessCount != accessCount {
		t.Errorf("expected access count to be %d, got %d", accessCount, pass.AccessCount)
	}
	if pass.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set, but it is zero")
	}
	if pass.UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be set, but it is zero")
	}
	if pass.UserID != user.ID {
		t.Errorf("expected UserID to be %d, got %d", user.ID, pass.UserID)
	}
}

// TestIncrementAccessCount tests the incrementAccessCount method
func TestIncrementAccessCount(t *testing.T) {
	pass := &models.Password{
		AccessCount: 0,
	}

	pass.IncrementAccessCount()
	if pass.AccessCount != 1 {
		t.Errorf("expected access count to be 1, got %d", pass.AccessCount)
	}

	pass.IncrementAccessCount()
	if pass.AccessCount != 2 {
		t.Errorf("expected access count to be 2, got %d", pass.AccessCount)
	}
}
