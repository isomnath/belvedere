package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type DataDogTraceConfigTestSuite struct {
	suite.Suite
	config *DataDogTraceConfig
}

func (suite *DataDogTraceConfigTestSuite) SetupTest() {
	_ = os.Setenv(dataDogTraceAgentEnabled, fmt.Sprintf("%t", true))
	_ = os.Setenv(dataDogTraceAgentHost, fmt.Sprintf("%s", "localhost"))
	_ = os.Setenv(dataDogTraceAgentPort, fmt.Sprintf("%s", "8126"))
	_ = os.Setenv(dataDogTraceLogLevel, fmt.Sprintf("%s", "DEBUG"))
	_ = os.Setenv(dataDogLogInjectionEnabled, fmt.Sprintf("%t", true))

	viper.New()
	viper.AutomaticEnv()
	suite.config = dataDogTraceConfig()
}

func (suite *DataDogTraceConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv(dataDogTraceAgentEnabled)
	_ = os.Unsetenv(dataDogTraceAgentHost)
	_ = os.Unsetenv(dataDogTraceAgentPort)
	_ = os.Unsetenv(dataDogTraceLogLevel)
	_ = os.Unsetenv(dataDogLogInjectionEnabled)
}

func (suite *DataDogTraceConfigTestSuite) TestAllConfigs() {
	suite.Equal(true, suite.config.Enabled())
	suite.Equal("localhost", suite.config.Host())
	suite.Equal(8126, suite.config.Port())
	suite.Equal("DEBUG", suite.config.LogLevel())
	suite.Equal(true, suite.config.LogInjectionEnabled())
}

func TestDataDogTraceConfigTestSuite(t *testing.T) {
	suite.Run(t, new(DataDogTraceConfigTestSuite))
}
