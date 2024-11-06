package main

import (
	"backend/api"
	"backend/db"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	dataSource, err := db.NewSQLiteDataSource("../db/password_manager.db") // Replace "passkeeper.db" with your SQLite file
	if err != nil {
		log.Fatalf("Failed to initialize data source: %v", err)
	}
	defer dataSource.Close()

	// Initialize services
	userService := db.NewUserServiceSql(dataSource)
	passwordService := db.NewPasswordServiceSql(dataSource)

	// Initialize the router with services
	router := api.NewRouter(userService, passwordService)

	// Start the server
	port := ":8080" // Use any port you prefer
	log.Printf("Starting server on %s...", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
