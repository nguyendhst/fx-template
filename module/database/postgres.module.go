package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/shared/util"
	sqlc "github.com/nguyendhst/lagile/sqlc/generated"
)

// New Postgres client
func NewPostgresClient(env *config.Env) (*sqlc.Queries, error) {
	// ctx := context.Background()
	withRetry := 3
	uri := getDSN(env)

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

func getDSN(env *config.Env) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.Database.Postgres.User,
		env.Database.Postgres.Password,
		env.Database.Postgres.Host,
		strconv.Itoa(env.Database.Postgres.Port),
		env.Database.Postgres.Name,
	)
}
