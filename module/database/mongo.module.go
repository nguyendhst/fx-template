package database

import (
	"context"
	"fmt"
	"time"

	"github.com/nguyendhst/lagile/module/config"
	"github.com/nguyendhst/lagile/shared/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient(env *config.Env) (*mongo.Client, error) {
	withRetry := 3
	uri := getMongoDSN(env)

	var client *mongo.Client
	var err error

	if err := util.Retry(withRetry, 5*time.Second, func() error {
		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}

		return client.Ping(context.Background(), readpref.Primary())
	}); err != nil {
		return nil, err
	}

	return client, nil
}

func getMongoDSN(env *config.Env) string {
	return fmt.Sprintf("mongodb://%s:%s@%s/?retryWrites=true&w=majority",
		env.Database.Mongo.User,
		env.Database.Mongo.Password,
		env.Database.Mongo.Host,
	)
}