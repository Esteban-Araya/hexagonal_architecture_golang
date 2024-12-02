package storage

import (
	"database/sql"
)

type DBinterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}
