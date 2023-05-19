package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	name                 string
	version              string
	environment          string
	webPort              int
	nonWebPort           int
	logLevel             string
	healthCheckAPIPath   string
	swaggerEnabled       bool
	swaggerDocsDirectory string
	postgres             *PostgresConfig
	mongo                *MongoConfig
	redis                *RedisConfig
	newRelic             *NewRelicConfig
	dataDog              *DataDogConfig
	dataDogTracer        *DataDogTraceConfig
	sentry               *SentryConfig
	translation          *TranslationConfig
	custom               interface{}
}

var baseConfig *Config

func LoadBaseConfig() {
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
	viper.AddConfigPath("../../../../../")
	viper.SetConfigType("yaml")
	_ = viper.ReadInConfig()

	baseConfig = &Config{
		name:                 getString(appName, true),
		version:              getString(appVersion, true),
		environment:          getString(appEnvironment, true),
		webPort:              getInt(appWebPort, false),
		nonWebPort:           getInt(appNonWebPort, false),
		logLevel:             getString(appLogLevel, true),
		healthCheckAPIPath:   getString(appHealthCheckAPIPath, false),
		swaggerEnabled:       getBool(appSwaggerEnabled, false),
		swaggerDocsDirectory: getString(appSwaggerDocsDirectory, false),
		newRelic:             newRelicConfig(),
		dataDog:              dataDogConfig(),
		dataDogTracer:        dataDogTraceConfig(),
		sentry:               sentryConfig(),
	}
}

func LoadPostgresConfig() {
	baseConfig.postgres = postgresConfig()
}

func LoadMongoConfig() {
	baseConfig.mongo = mongoConfig()
}

func LoadRedisConfig() {
	baseConfig.redis = redisConfig()
}

func LoadTranslationsConfig() {
	baseConfig.translation = translationsConfig()
}

func LoadCustomConfig(v interface{}) {
	customConfig(v)
	baseConfig.custom = v
}

func GetAppName() string {
	return baseConfig.name
}

func GetAppVersion() string {
	return baseConfig.version
}

func GetAppEnvironment() string {
	return baseConfig.environment
}

func GetAppWebPort() int {
	return baseConfig.webPort
}

func GetAppNonWebPort() int {
	return baseConfig.nonWebPort
}

func GetAppLogLevel() string {
	return baseConfig.logLevel
}

func GetAppHealthCheckAPIPath() string {
	return baseConfig.healthCheckAPIPath
}

func GetSwaggerEnabled() bool {
	return baseConfig.swaggerEnabled
}

func GetSwaggerDocsDirectory() string {
	return baseConfig.swaggerDocsDirectory
}

func GetPostgresConfig() *PostgresConfig {
	return baseConfig.postgres
}

func GetMongoConfig() *MongoConfig {
	return baseConfig.mongo
}

func GetRedisConfig() *RedisConfig {
	return baseConfig.redis
}

func GetNewRelicConfig() *NewRelicConfig {
	return baseConfig.newRelic
}

func GetDataDogConfig() *DataDogConfig {
	return baseConfig.dataDog
}

func GetDataDogTracerConfig() *DataDogTraceConfig {
	return baseConfig.dataDogTracer
}

func GetSentryConfig() *SentryConfig {
	return baseConfig.sentry
}

func GetTranslationConfig() *TranslationConfig {
	return baseConfig.translation
}

func GetCustomConfig() interface{} {
	return baseConfig.custom
}
