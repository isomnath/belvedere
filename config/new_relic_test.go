package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/spf13/viper"
)

type NewRelicConfigTestSuite struct {
	suite.Suite
	config *NewRelicConfig
}

func (suite *NewRelicConfigTestSuite) SetupTest() {
	_ = os.Setenv(newRelicEnabled, fmt.Sprintf("%t", true))
	_ = os.Setenv(newRelicLicenseKey, fmt.Sprintf("%s", "dummy"))

	viper.New()
	viper.AutomaticEnv()
	suite.config = newRelicConfig()
}

func (suite *NewRelicConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv(newRelicEnabled)
	_ = os.Unsetenv(newRelicLicenseKey)
}

func (suite *NewRelicConfigTestSuite) TestAllConfigs() {
	suite.Equal(true, suite.config.Enabled())
	suite.Equal("dummy", suite.config.LicenseKey())
}

func TestNewRelicConfigTestSuite(t *testing.T) {
	suite.Run(t, new(NewRelicConfigTestSuite))
}
