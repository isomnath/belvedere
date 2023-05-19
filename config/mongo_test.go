package config

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/spf13/viper"
)

type MongoConfigTestSuite struct {
	suite.Suite
	config *MongoConfig
}

func (suite *MongoConfigTestSuite) SetupTest() {
	_ = os.Setenv(mongoHosts, fmt.Sprintf("%s", "localhost:27017"))
	_ = os.Setenv(mongoDbName, fmt.Sprintf("%s", "dummy_db"))
	_ = os.Setenv(mongoUsername, fmt.Sprintf("%s", "dummy_username"))
	_ = os.Setenv(mongoPassword, fmt.Sprintf("%s", "dummy_password"))
	_ = os.Setenv(mongoPoolSize, fmt.Sprintf("%s", "20"))
	_ = os.Setenv(mongoSocketTimeout, fmt.Sprintf("%s", "500"))
	_ = os.Setenv(mongoConnectionTimeout, fmt.Sprintf("%s", "1000"))

	viper.New()
	viper.AutomaticEnv()
	suite.config = mongoConfig()
}

func (suite *MongoConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv(mongoHosts)
	_ = os.Unsetenv(mongoDbName)
	_ = os.Unsetenv(mongoUsername)
	_ = os.Unsetenv(mongoPassword)
	_ = os.Unsetenv(mongoPoolSize)
	_ = os.Unsetenv(mongoSocketTimeout)
	_ = os.Unsetenv(mongoConnectionTimeout)
}

func (suite *MongoConfigTestSuite) TestAllConfigs() {
	suite.Equal("localhost:27017", suite.config.Hosts())
	suite.Equal("dummy_db", suite.config.DbName())
	suite.Equal("dummy_username", suite.config.Username())
	suite.Equal("dummy_password", suite.config.Password())
	suite.Equal(20, suite.config.PoolSize())
	suite.Equal(time.Duration(500), suite.config.SocketTimeout())
	suite.Equal(time.Duration(1000), suite.config.ConnectionTimeout())
	suite.Equal("mongodb://dummy_username:dummy_password@localhost:27017/dummy_db", suite.config.ConnectionURL())
}

func TestMongoConfigTestSuite(t *testing.T) {
	suite.Run(t, new(MongoConfigTestSuite))
}
