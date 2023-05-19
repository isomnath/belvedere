package store

import (
	"testing"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/instrumentation"
	"github.com/isomnath/belvedere/log"

	"github.com/stretchr/testify/suite"
)

type RedisTestSuite struct {
	suite.Suite
	redisConfig *config.RedisConfig
}

func (suite *RedisTestSuite) SetupTest() {
	config.LoadBaseConfig()
	config.LoadRedisConfig()
	instrumentation.StartDDTracer()
	defer instrumentation.StopDDTracer()
	log.Setup()
	suite.redisConfig = config.GetRedisConfig()
}

func (suite *RedisTestSuite) TestSuccessfulUnsecuredRedisConnection() {
	RedisConnect(suite.redisConfig, 0)
	suite.NotNil(GetRedisClient())
}

func TestRedisTestSuite(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}
