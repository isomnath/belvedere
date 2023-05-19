package config

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
	config *Config
}

func (suite *ConfigTestSuite) TestBaseConfigs() {
	LoadBaseConfig()

	suite.Equal("sample-app", GetAppName())
	suite.Equal("0.0.1", GetAppVersion())
	suite.Equal("staging", GetAppEnvironment())
	suite.Equal(8181, GetAppWebPort())
	suite.Equal(9191, GetAppNonWebPort())
	suite.Equal("DEBUG", GetAppLogLevel())
	suite.Equal("/ping", GetAppHealthCheckAPIPath())
	suite.Equal(true, GetSwaggerEnabled())
	suite.Equal("/docs/", GetSwaggerDocsDirectory())
	suite.Equal(true, GetNewRelicConfig().Enabled())
	suite.Equal("dummy", GetNewRelicConfig().LicenseKey())
	suite.Equal(true, GetDataDogConfig().Enabled())
	suite.Equal("localhost", GetDataDogConfig().Host())
	suite.Equal(8125, GetDataDogConfig().Port())
	suite.Equal(time.Duration(20), GetDataDogConfig().FlushPeriod())
	suite.Equal(true, GetDataDogTracerConfig().Enabled())
	suite.Equal("localhost", GetDataDogTracerConfig().Host())
	suite.Equal(8126, GetDataDogTracerConfig().Port())
	suite.Equal("DEBUG", GetDataDogTracerConfig().LogLevel())
	suite.Equal(true, GetDataDogTracerConfig().LogInjectionEnabled())
	suite.Equal(true, GetSentryConfig().Enabled())
	suite.Equal("dummy_DSN", GetSentryConfig().DSN())
}

func (suite *ConfigTestSuite) TestAllConfigs() {
	type TestStruct struct {
		TestKeyOne   int      `mapstructure:"TEST_KEY_ONE"`
		TestKeyTwo   string   `mapstructure:"TEST_KEY_TWO"`
		TestKeyThree bool     `mapstructure:"TEST_KEY_THREE"`
		TestKeyFour  []string `mapstructure:"TEST_KEY_FOUR"`
		TestKeyFive  []int    `mapstructure:"TEST_KEY_FIVE"`
	}

	_ = os.Setenv("TEST_KEY_ONE", fmt.Sprintf("%d", 20))
	defer os.Unsetenv("TEST_KEY_ONE")
	_ = os.Setenv("TEST_KEY_TWO", fmt.Sprintf("%s", "test_value"))
	defer os.Unsetenv("TEST_KEY_TWO")
	_ = os.Setenv("TEST_KEY_THREE", fmt.Sprintf("%t", true))
	defer os.Unsetenv("TEST_KEY_THREE")
	_ = os.Setenv("TEST_KEY_FOUR", fmt.Sprintf("%s", "val_1,val_2,val_3"))
	defer os.Unsetenv("TEST_KEY_FOUR")
	_ = os.Setenv("TEST_KEY_FIVE", fmt.Sprintf("%s", "1,2,3"))
	defer os.Unsetenv("TEST_KEY_FIVE")

	var ts TestStruct

	LoadBaseConfig()
	LoadTranslationsConfig()
	LoadRedisConfig()
	LoadPostgresConfig()
	LoadMongoConfig()
	LoadCustomConfig(&ts)

	suite.Equal("sample-app", GetAppName())
	suite.Equal("0.0.1", GetAppVersion())
	suite.Equal("staging", GetAppEnvironment())
	suite.Equal(8181, GetAppWebPort())
	suite.Equal(9191, GetAppNonWebPort())
	suite.Equal("DEBUG", GetAppLogLevel())
	suite.Equal("/ping", GetAppHealthCheckAPIPath())
	suite.Equal(true, GetSwaggerEnabled())
	suite.Equal("/docs/", GetSwaggerDocsDirectory())
	suite.Equal(true, GetNewRelicConfig().Enabled())
	suite.Equal("dummy", GetNewRelicConfig().LicenseKey())
	suite.Equal(true, GetDataDogConfig().Enabled())
	suite.Equal("localhost", GetDataDogConfig().Host())
	suite.Equal(8125, GetDataDogConfig().Port())
	suite.Equal(time.Duration(20), GetDataDogConfig().FlushPeriod())
	suite.Equal(true, GetSentryConfig().Enabled())
	suite.Equal("dummy_DSN", GetSentryConfig().DSN())

	suite.Equal(true, GetTranslationConfig().Enabled())
	suite.Equal("./i18n", GetTranslationConfig().Path())
	suite.Equal([]string{"en_ID", "id_ID", "en_MY", "my_MY", "en_SG", "ch_SG"}, GetTranslationConfig().WhitelistedLocales())
	suite.Equal("en_EN", GetTranslationConfig().DefaultLocale())

	suite.Equal("localhost", GetRedisConfig().Host())
	suite.Equal(6379, GetRedisConfig().Port())
	suite.Equal("redis", GetRedisConfig().Username())
	suite.Equal("redis", GetRedisConfig().Password())

	suite.Equal("localhost", GetPostgresConfig().Host())
	suite.Equal(5432, GetPostgresConfig().Port())
	suite.Equal("belvedere", GetPostgresConfig().DbName())
	suite.Equal("belvedere", GetPostgresConfig().Username())
	suite.Equal("password", GetPostgresConfig().Password())
	suite.Equal(20, GetPostgresConfig().PoolSize())
	suite.Equal("./pg_migrations", GetPostgresConfig().MigrationsDirectory())

	suite.Equal("localhost:27017", GetMongoConfig().Hosts())
	suite.Equal("admin", GetMongoConfig().DbName())
	suite.Equal("belvedere", GetMongoConfig().Username())
	suite.Equal("belvedere", GetMongoConfig().Password())
	suite.Equal(20, GetMongoConfig().PoolSize())
	suite.Equal(time.Duration(500), GetMongoConfig().SocketTimeout())
	suite.Equal(time.Duration(1000), GetMongoConfig().ConnectionTimeout())

	suite.Equal(&TestStruct{
		TestKeyOne:   20,
		TestKeyTwo:   "test_value",
		TestKeyThree: true,
		TestKeyFour:  []string{"val_1", "val_2", "val_3"},
		TestKeyFive:  []int{1, 2, 3},
	}, GetCustomConfig())

}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
