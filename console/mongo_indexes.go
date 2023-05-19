package console

import (
	"context"
	"strings"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/log"
	"github.com/isomnath/belvedere/store"

	"go.mongodb.org/mongo-driver/mongo"
)

// CreateIndexes - Allows creating mongoDB indexes
// Ideally to be used as console command
func CreateIndexes(collectionName string, models []mongo.IndexModel) error {
	ctx := context.Background()
	database := store.GetMongoClient().Database(config.GetMongoConfig().DbName())
	collection := database.Collection(collectionName)

	indexes := collection.Indexes()

	for _, model := range models {
		_, err := indexes.CreateOne(context.Background(), model)
		if err != nil {
			errMessage := err.Error()
			if !(strings.Contains(errMessage, "IndexOptionsConflict") ||
				strings.Contains(errMessage, "IndexKeySpecsConflict")) {
				log.Log.Errorf(ctx, "failed to create indexes for collection: %s with error: %v", collectionName, err)
				return err
			}
		}
	}

	log.Log.MongoInfof(ctx, "successfully created indexes for collection: %s", collectionName)
	return nil
}
