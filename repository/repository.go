package repository

import (
	"github.com/nguyendhst/lagile/module/config"
	user_postgres "github.com/nguyendhst/lagile/repository/user/postgres"
	constant "github.com/nguyendhst/lagile/shared/constant"
	sqlc "github.com/nguyendhst/lagile/sqlc/generated"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

type (
	params struct {
		fx.In
		// should be DB Module instead -?
		postgres *sqlc.Queries `optional:"true"`
		mongo    *mongo.Client `optional:"true"`
		cfg      *config.Env
	}
)

// Provide your repository implementations here.
func New() fx.Option {
	return fx.Provide(
		NewUserRepository,
	)
}

func NewUserRepository(p params) UserRepository {
	s := p.cfg.Repository.User.Store
	if s == constant.STORE_TYPE_POSTGRES {
		return user_postgres.NewRepository(p.postgres)
	}
	if s == constant.STORE_TYPE_MONGO {
		return nil
	}
	panic("no database connection")
}
