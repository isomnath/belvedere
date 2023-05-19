package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type RedisConfigTestSuite struct {
	suite.Suite
	config *RedisConfig
}

func (suite *RedisConfigTestSuite) SetupTest() {
	_ = os.Setenv(redisHost, fmt.Sprintf("%s", "localhost"))
	_ = os.Setenv(redisPort, fmt.Sprintf("%s", "6379"))
	_ = os.Setenv(redisUsername, fmt.Sprintf("%s", "dummy_username"))
	_ = os.Setenv(redisPassword, fmt.Sprintf("%s", "dummy_password"))

	viper.New()
	viper.AutomaticEnv()
	suite.config = redisConfig()
}

func (suite *RedisConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv(redisHost)
	_ = os.Unsetenv(redisPort)
	_ = os.Unsetenv(redisUsername)
	_ = os.Unsetenv(redisPassword)
}

func (suite *RedisConfigTestSuite) TestAllConfigs() {
	suite.Equal("localhost", suite.config.Host())
	suite.Equal(6379, suite.config.Port())
	suite.Equal("dummy_username", suite.config.Username())
	suite.Equal("dummy_password", suite.config.Password())
}

func TestRedisConfigTestSuite(t *testing.T) {
	suite.Run(t, new(RedisConfigTestSuite))
}
