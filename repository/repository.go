package repository

import (
	"github.com/nguyendhst/lagile/module/config"
	user_postgres "github.com/nguyendhst/lagile/repository/user/postgres"
	constant "github.com/nguyendhst/lagile/shared/constant"
	sqlc "github.com/nguyendhst/lagile/sqlc/generated"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewUserRepository,
	),
)

type (
	params struct {
		fx.In
		// should be DB Module instead -?
		Postgres *sqlc.Queries `optional:"true"`
		Mongo    *mongo.Client `optional:"true"`
		Configs  *config.Config
	}
)

func NewUserRepository(p params) UserRepository {
	s := p.Configs.Repository.User.Store
	if s == constant.STORE_TYPE_POSTGRES {
		return user_postgres.NewRepository(p.Postgres)
	}
	if s == constant.STORE_TYPE_MONGO {
		return nil
	}
	panic("no database connection")
}
