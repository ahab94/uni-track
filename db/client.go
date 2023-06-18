package db

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ahab94/uni-track/config"
)

type MongoClient struct {
	client   *mongo.Client
	database *mongo.Database
}

type Collection struct {
	collection *mongo.Collection
}

func NewMongoClient(ctx context.Context) (*MongoClient, error) {
	dbName := viper.GetString(config.DbName)
	dbHost := viper.GetString(config.DbHost)
	dbPort := viper.GetString(config.DbPort)

	mongoURL := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	return &MongoClient{
		client:   client,
		database: client.Database(dbName),
	}, nil
}

func (mc *MongoClient) Collection(name string) *Collection {
	collection := mc.database.Collection(name)

	return &Collection{
		collection: collection,
	}
}

func (mc *MongoClient) Disconnect(ctx context.Context) error {
	return mc.client.Disconnect(ctx)
}
