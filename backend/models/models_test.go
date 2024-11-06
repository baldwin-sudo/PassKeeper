package models

import (
	"testing"
)

// TestNewUser tests the creation of a new User
func TestNewUser(t *testing.T) {
	username := "testuser"
	password := "securepassword"

	user, err := NewUser(username, password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if user.Username != username {
		t.Fatalf("Expected username %s, got %s", username, user.Username)
	}
	if user.CreatedAt.IsZero() {
		t.Fatal("Expected CreatedAt to be set, but it is zero")
	}
}

// TestCheckPassword tests the password checking function
func TestCheckPassword(t *testing.T) {
	// For the sake of the test, we will directly set the HashedMasterPassword
	user := User{Username: "testuser", MasterPassword: "securepassword"}
	user.HashedMasterPassword, _ = HashPassword(user.MasterPassword)
	if !user.CheckPassword("securepassword") {
		t.Fatal("Expected password check to pass")
	}
	if user.CheckPassword("hashedwrongpassword") {
		t.Fatal("Expected password check to fail for incorrect password")
	}
}
func TestGenerateSalt(t *testing.T) {
	saltSize := 16
	salt, err := GenerateSalt(saltSize)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(salt) != saltSize {
		t.Fatalf("Expected salt size %d, got %d", saltSize, len(salt))
	}
}

func TestHashPassword(t *testing.T) {
	password := "securepassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !CheckPassword(hashedPassword, password) {
		t.Fatal("Expected password check to pass for correct password")
	}
	if CheckPassword(hashedPassword, "wrongpassword") {
		t.Fatal("Expected password check to fail for incorrect password")
	}
}

func TestDeriveKey(t *testing.T) {
	password := "securepassword"
	salt, err := GenerateSalt(16)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	iterations := 10000
	key := DeriveKey(password, salt, iterations)
	if len(key) != 32 {
		t.Fatalf("Expected key length to be 32 bytes, got %d", len(key))
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("examplekey12345678901234567890°") // 32 bytes for AES-256
	plainText := []byte("Hello, World!")

	cipherText, err := Encrypt(plainText, key)
	if err != nil {
		t.Fatalf("Expected no error during encryption, got %v", err)
	}

	decryptedText, err := Decrypt(string(cipherText), key)
	if err != nil {
		t.Fatalf("Expected no error during decryption, got %v", err)
	}

	if string(decryptedText) != string(plainText) {
		t.Fatalf("Expected decrypted text to be %s, got %s", plainText, decryptedText)
	}
}
