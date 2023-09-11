package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/nguyendhst/clean-architecture-skeleton/module/config"
	sqlc "github.com/nguyendhst/clean-architecture-skeleton/sqlc/generated"
)

type (
	database struct {
		*sql.DB
		*config.Env
		*sqlc.Queries
	}

	table struct {
		*sql.DB
		query string
	}
)

func NewPostgresDatabase(env *config.Env) (Database, error) {
	db := &database{
		Env: env,
		DB:  nil,
	}
	err := db.Connect()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *database) String() string {
	return "postgres"
}

func (db *database) Connect() error {
	var err error
	var dsn = getDSN(db.Env)
	fmt.Println(dsn)
	db.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	db.Queries = sqlc.New(db.DB)
	return nil
}

func (db *database) Close() error {
	return db.DB.Close()
}

func (db *database) GetDBConnection() *sql.DB {
	return db.DB
}

func (db *database) GetQueries() *sqlc.Queries {
	return db.Queries
}

func getDSN(env *config.Env) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)
}
