package store

import (
	"context"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	client *mongo.Client
}

var mg *MongoStore

// TODO: Add tests for error blocks
// TODO: Add client options for socketTimeout and connectionTimeout
func mgInitialize(config *config.MongoConfig) (*mongo.Client, error) {
	ctx := context.Background()
	log.Log.MongoInfof(ctx, "attempting to connect: %s", config.Hosts())

	clientOptions := options.Client().ApplyURI(config.ConnectionURL())

	poolSize := uint64(config.PoolSize())
	clientOptions.MaxPoolSize = &poolSize

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Log.MongoErrorf(ctx, "failed to connect to mongo server: %v", err)
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Log.MongoErrorf(ctx, "ping to mongo server failed: %v", err)
		return nil, err
	}

	log.Log.MongoInfof(ctx, "successfully connected")
	return client, nil
}

// MongoConnect - Connects to mongo and initializes the DB client to be reused across the application
func MongoConnect(config *config.MongoConfig) error {
	mgClient, err := mgInitialize(config)
	if err != nil {
		return err
	}
	mg = &MongoStore{client: mgClient}
	return nil
}

// GetMongoClient - Returns the instance of mongo DB client created when MongoConnect was invoked
func GetMongoClient() *mongo.Client {
	return mg.client
}
