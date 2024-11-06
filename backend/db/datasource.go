package db

import "database/sql"

type DataSource interface {
	Connect(string) error
	Close() error
	Execute(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
