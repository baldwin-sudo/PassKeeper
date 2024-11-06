package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver

	"fmt"
)

// SQLiteDataSource implements the DataSource interface
type SQLiteDataSource struct {
	db *sql.DB
}

// NewSQLiteDataSource initializes a new SQLiteDataSource and connects to the database
func NewSQLiteDataSource(dsn string) (*SQLiteDataSource, error) {
	dataSource := &SQLiteDataSource{}
	if err := dataSource.Connect(dsn); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return dataSource, nil
}

// Connect establishes a connection to the SQLite database
func (s *SQLiteDataSource) Connect(dsn string) error {
	var err error
	s.db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}
	return s.db.Ping()
}

// Close closes the database connection
func (s *SQLiteDataSource) Close() error {
	return s.db.Close()
}

// Execute executes a query with optional arguments and returns the result
func (s *SQLiteDataSource) Execute(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

// Query performs a query with optional arguments and returns the rows
func (s *SQLiteDataSource) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}
