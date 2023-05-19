package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type SentryConfigTestSuite struct {
	suite.Suite
	config *SentryConfig
}

func (suite *SentryConfigTestSuite) SetupTest() {
	_ = os.Setenv(sentryEnabled, fmt.Sprintf("%t", true))
	_ = os.Setenv(sentryDSN, fmt.Sprintf("%s", "dummy"))

	viper.New()
	viper.AutomaticEnv()
	suite.config = sentryConfig()
}

func (suite *SentryConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv(sentryEnabled)
	_ = os.Unsetenv(sentryDSN)
}

func (suite *SentryConfigTestSuite) TestAllConfigs() {
	suite.Equal(true, suite.config.Enabled())
	suite.Equal("dummy", suite.config.DSN())
}

func TestSentryConfigTestSuite(t *testing.T) {
	suite.Run(t, new(SentryConfigTestSuite))
}
