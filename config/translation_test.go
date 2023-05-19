package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type TranslationConfigTestSuite struct {
	suite.Suite
	config *TranslationConfig
}

func (suite *TranslationConfigTestSuite) SetupTest() {
	_ = os.Setenv(translationsEnabled, fmt.Sprintf("%t", true))
	_ = os.Setenv(translationsPath, fmt.Sprintf("%s", "./i18n"))
	_ = os.Setenv(translationsWhitelistedLocales, fmt.Sprintf("%s", "en_ID, id_ID, en_MY, my_MY"))
	_ = os.Setenv(translationsDefaultLocale, fmt.Sprintf("%s", "en_EN"))

	viper.New()
	viper.AutomaticEnv()
	suite.config = translationsConfig()

}

func (suite *TranslationConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv(translationsEnabled)
	_ = os.Unsetenv(translationsPath)
	_ = os.Unsetenv(translationsWhitelistedLocales)
	_ = os.Unsetenv(translationsDefaultLocale)
}

func (suite *TranslationConfigTestSuite) TestAllConfigs() {
	suite.Equal(true, suite.config.Enabled())
	suite.Equal("./i18n", suite.config.Path())
	suite.Equal([]string{"en_ID", "id_ID", "en_MY", "my_MY"}, suite.config.WhitelistedLocales())
	suite.Equal("en_EN", suite.config.DefaultLocale())
}

func TestTranslationTestSuite(t *testing.T) {
	suite.Run(t, new(TranslationConfigTestSuite))
}
