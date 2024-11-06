package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// MySQLDataSource implements the DataSource interface for MySQL
type MySQLDataSource struct {
	db *sql.DB
}

// NewMySQLDataSource initializes a new MySQLDataSource and connects to the database
func NewMySQLDataSource(dsn string) (*MySQLDataSource, error) {
	dataSource := &MySQLDataSource{}
	if err := dataSource.Connect(dsn); err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL database: %w", err)
	}
	return dataSource, nil
}

// Connect establishes a connection to the MySQL database
func (m *MySQLDataSource) Connect(dsn string) error {
	var err error
	m.db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return m.db.Ping()
}

// Close closes the database connection
func (m *MySQLDataSource) Close() error {
	return m.db.Close()
}

// Execute executes a query with optional arguments and returns the result
func (m *MySQLDataSource) Execute(query string, args ...interface{}) (sql.Result, error) {
	return m.db.Exec(query, args...)
}

// Query performs a query with optional arguments and returns the rows
func (m *MySQLDataSource) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.db.Query(query, args...)
}
