package config

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type DataDogConfigTestSuite struct {
	suite.Suite
	config *DataDogConfig
}

func (suite *DataDogConfigTestSuite) SetupTest() {
	_ = os.Setenv(dataDogEnabled, fmt.Sprintf("%t", true))
	_ = os.Setenv(dataDogHost, fmt.Sprintf("%s", "localhost"))
	_ = os.Setenv(dataDogPort, fmt.Sprintf("%s", "8125"))
	_ = os.Setenv(dataDogFlushPeriodSeconds, fmt.Sprintf("%s", "20"))

	viper.New()
	viper.AutomaticEnv()
	suite.config = dataDogConfig()
}

func (suite *DataDogConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv(dataDogEnabled)
	_ = os.Unsetenv(dataDogHost)
	_ = os.Unsetenv(dataDogPort)
	_ = os.Unsetenv(dataDogFlushPeriodSeconds)
}

func (suite *DataDogConfigTestSuite) TestAllConfigs() {
	suite.Equal(true, suite.config.Enabled())
	suite.Equal("localhost", suite.config.Host())
	suite.Equal(8125, suite.config.Port())
	suite.Equal(time.Duration(20), suite.config.FlushPeriod())
}

func TestStatsDConfigTestSuite(t *testing.T) {
	suite.Run(t, new(DataDogConfigTestSuite))
}
