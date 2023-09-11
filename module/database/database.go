package database

import (
	"database/sql"

	sqlc "github.com/nguyendhst/lagile/sqlc/generated"
)

type Database interface {
	Connect() error
	Close() error
	GetDBConnection() *sql.DB
	GetQueries() *sqlc.Queries
	String() string
}
