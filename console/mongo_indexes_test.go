package console

import (
	"context"
	"testing"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/instrumentation"
	"github.com/isomnath/belvedere/log"
	"github.com/isomnath/belvedere/store"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type MongoIndexSupportTestSuite struct {
	suite.Suite
	collectionName string
	indexName      string
	dbClient       *mongo.Client
}

func (suite *MongoIndexSupportTestSuite) SetupSuite() {
	config.LoadBaseConfig()
	config.LoadMongoConfig()
	instrumentation.StartDDTracer()
	log.Setup()
	_ = store.MongoConnect(config.GetMongoConfig())
	suite.dbClient = store.GetMongoClient()
	suite.collectionName = "test_collection"
	suite.collectionName = "idx_test_collection"
}

func (suite *MongoIndexSupportTestSuite) TearDownSuite() {
	instrumentation.StopDDTracer()
}

func (suite *MongoIndexSupportTestSuite) TearDownTest() {
	_, _ = suite.dbClient.Database(config.GetMongoConfig().DbName()).Collection(suite.collectionName).Indexes().DropAll(context.Background())
}

func (suite *MongoIndexSupportTestSuite) TestCreateIndexesSuccessfully() {
	indexes := []mongo.IndexModel{
		{
			Keys: bsonx.MDoc{
				"test_key": bsonx.Int64(1),
			},
		},
	}
	err := CreateIndexes(suite.collectionName, indexes)
	suite.NoError(err)

	createdIndexes := suite.dbClient.Database(config.GetMongoConfig().DbName()).Collection(suite.collectionName).Indexes()
	suite.NotNil(createdIndexes)
}

func (suite *MongoIndexSupportTestSuite) TestCreateIndexesFails() {
	indexes := []mongo.IndexModel{
		{
			Keys: bsonx.MDoc{
				"test_key": bsonx.Int64(1),
			},
		},
	}
	_ = CreateIndexes(suite.collectionName, indexes)
	indexes = []mongo.IndexModel{
		{
			Keys: bsonx.MDoc{
				"test_key": bsonx.Timestamp(uint32(1), uint32(12)),
			},
		},
	}
	err := CreateIndexes(suite.collectionName, indexes)
	suite.Equal("invalid index value", err.Error())
}

func TestMongoIndexSupportTestSuite(t *testing.T) {
	suite.Run(t, new(MongoIndexSupportTestSuite))
}
