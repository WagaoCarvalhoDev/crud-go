package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongodbUrl       = "MONGODB_URL"
	mongodb_users_db = "MONGODB_USERS_DB"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(mongodbUrl)
	mongodb_database := os.Getenv(mongodb_users_db)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
