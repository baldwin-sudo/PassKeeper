package main

import (
	"backend/api"
	"backend/db"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {
	// Database file path
	dbPath := "./password_manager.db"

	// Initialize the database with schema
	dataSource, err := db.NewSQLiteDataSource(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize data source: %v", err)
	}
	defer dataSource.Close()

	// Apply schema if tables don't exist
	if err := initDBSchema(dataSource); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Initialize services
	userService := db.NewUserServiceSql(dataSource)
	passwordService := db.NewPasswordServiceSql(dataSource)

	// Initialize the router with services
	router := api.NewRouter(userService, passwordService)

	// Start the server
	port := "0.0.0.0:8080" // Use any port you prefer
	log.Printf("Starting server on %s...", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// initDBSchema applies the schema to the database if tables do not exist
func initDBSchema(db db.DataSource) error {
	// Check if any tables exist
	tablesExist, err := checkTablesExist(db)
	if err != nil {
		return fmt.Errorf("could not check if tables exist: %v", err)
	}

	// If tables do not exist, apply the schema
	if !tablesExist {
		log.Println("Tables not found. Applying schema...")

		// Read schema from file
		schema, err := ioutil.ReadFile("./db_schema.sql")
		if err != nil {
			return fmt.Errorf("could not read schema file: %v", err)
		}

		// Execute schema
		if _, err := db.Execute(string(schema)); err != nil {
			return fmt.Errorf("could not execute schema: %v", err)
		}

		log.Println("Database schema initialized successfully.")
	} else {
		log.Println("Tables already exist. Skipping schema application.")
	}
	return nil
}

// checkTablesExist checks if the required tables exist in the database
func checkTablesExist(db db.DataSource) (bool, error) {
	// Query the database to check for existing tables (e.g., `users`, `passwords`)
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name IN ('users', 'passwords');"
	rows, err := db.Query(query)
	if err != nil {
		return false, fmt.Errorf("could not check tables: %v", err)
	}
	defer rows.Close()

	// If rows are returned, it means the tables exist
	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			return false, fmt.Errorf("could not scan row: %v", err)
		}
	}

	// If there is any table found, return true
	if tableName != "" {
		return true, nil
	}

	return false, nil
}
