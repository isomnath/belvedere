package store

import (
	"os"
	"testing"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/instrumentation"
	"github.com/isomnath/belvedere/log"

	"github.com/stretchr/testify/suite"
)

type MongoTestSuite struct {
	suite.Suite
}

func (suite *MongoTestSuite) SetupSuite() {
	config.LoadBaseConfig()
	instrumentation.StartDDTracer()
	defer instrumentation.StopDDTracer()
	log.Setup()
}

func (suite *MongoTestSuite) TestSuccessfulMongoConnection() {
	_ = os.Setenv("MONGO_HOSTS", "localhost:27017")
	config.LoadMongoConfig()
	mgConfig := config.GetMongoConfig()
	err := MongoConnect(mgConfig)
	suite.NoError(err)
	suite.NotNil(GetMongoClient())
}

func (suite *MongoTestSuite) TestMongoPingError() {
	_ = os.Setenv("MONGO_HOSTS", "invalid:27017")
	defer os.Unsetenv("MONGO_HOSTS")
	config.LoadMongoConfig()
	mgConfig := config.GetMongoConfig()
	err := MongoConnect(mgConfig)
	suite.Error(err)
}

func TestMongoTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTestSuite))
}
