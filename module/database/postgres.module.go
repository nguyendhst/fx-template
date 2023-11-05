package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/nguyendhst/fx-template/module/config"
	"github.com/nguyendhst/fx-template/shared/util"
	sqlc "github.com/nguyendhst/fx-template/sqlc/generated"
)

// New Postgres client
func NewPostgresClient(cfg *config.Config) (*sqlc.Queries, error) {
	// ctx := context.Background()
	withRetry := 3
	uri := getDSN(cfg)

	var client *sqlc.Queries
	var sqldb *sql.DB
	var err error

	if err := util.Retry(withRetry, 5*time.Second, func() error {
		sqldb, err = sql.Open("postgres", uri)
		if err != nil {
			return err
		}
		client = sqlc.New(sqldb)
		return sqldb.Ping()
	}); err != nil {
		return nil, err
	}

	return client, nil
}

func getDSN(cfg *config.Config) string {
	fmt.Printf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Env.Database.Postgres.User,
		cfg.Env.Database.Postgres.Password,
		cfg.Env.Database.Postgres.Host,
		strconv.Itoa(cfg.Env.Database.Postgres.Port),
		cfg.Env.Database.Postgres.Name)
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Env.Database.Postgres.User,
		cfg.Env.Database.Postgres.Password,
		cfg.Env.Database.Postgres.Host,
		strconv.Itoa(cfg.Env.Database.Postgres.Port),
		cfg.Env.Database.Postgres.Name,
	)
}
