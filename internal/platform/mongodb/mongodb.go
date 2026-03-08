package mongodb

import (
	"context"
	"time"

	"github.com/rdl/core/internal/platform/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client = mongo.Client
type Database = mongo.Database
type Collection = mongo.Collection

func Connect(ctx context.Context, cfg config.MongoDBConfig) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(cfg.URI)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

func Database_(client *mongo.Client, cfg config.MongoDBConfig) *mongo.Database {
	return client.Database(cfg.Database)
}
