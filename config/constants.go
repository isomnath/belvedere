package config

const (
	appName                 = "APP_NAME"
	appVersion              = "APP_VERSION"
	appEnvironment          = "APP_ENVIRONMENT"
	appWebPort              = "APP_WEB_PORT"
	appNonWebPort           = "APP_NON_WEB_PORT"
	appLogLevel             = "APP_LOG_LEVEL"
	appHealthCheckAPIPath   = "APP_HEALTH_CHECK_API_PATH"
	appSwaggerEnabled       = "APP_SWAGGER_ENABLED"
	appSwaggerDocsDirectory = "APP_SWAGGER_DOCS_DIRECTORY"
)

const (
	translationsEnabled            = "TRANSLATIONS_ENABLED"
	translationsPath               = "TRANSLATIONS_PATH"
	translationsWhitelistedLocales = "TRANSLATIONS_WHITELISTED_LOCALES"
	translationsDefaultLocale      = "TRANSLATIONS_DEFAULT_LOCALE"
)

const (
	dataDogEnabled            = "DATA_DOG_ENABLED"
	dataDogHost               = "DATA_DOG_HOST"
	dataDogPort               = "DATA_DOG_PORT"
	dataDogFlushPeriodSeconds = "DATA_DOG_FLUSH_PERIOD_SECONDS"
)

const (
	dataDogTraceAgentEnabled   = "DD_TRACE_AGENT_ENABLED"
	dataDogTraceAgentHost      = "DD_AGENT_HOST"
	dataDogTraceAgentPort      = "DD_TRACE_AGENT_PORT"
	dataDogTraceLogLevel       = "DD_TRACE_LOG_LEVEL"
	dataDogLogInjectionEnabled = "DD_LOGS_INJECTION"
)

const (
	sentryEnabled = "SENTRY_ENABLED"
	sentryDSN     = "SENTRY_DSN"
)

const (
	newRelicEnabled    = "NEW_RELIC_ENABLED"
	newRelicLicenseKey = "NEW_RELIC_LICENSE_KEY"
)

const (
	redisHost     = "REDIS_HOST"
	redisPort     = "REDIS_PORT"
	redisUsername = "REDIS_USERNAME"
	redisPassword = "REDIS_PASSWORD"
)

const (
	postgresHost                = "POSTGRES_HOST"
	postgresPort                = "POSTGRES_PORT"
	postgresDbName              = "POSTGRES_DB_NAME"
	postgresUsername            = "POSTGRES_USERNAME"
	postgresPassword            = "POSTGRES_PASSWORD"
	postgresPoolSize            = "POSTGRES_POOL_SIZE"
	postgresMigrationsDirectory = "POSTGRES_MIGRATIONS_DIRECTORY"
)

const (
	mongoHosts             = "MONGO_HOSTS"
	mongoDbName            = "MONGO_DB_NAME"
	mongoUsername          = "MONGO_USERNAME"
	mongoPassword          = "MONGO_PASSWORD"
	mongoPoolSize          = "MONGO_POOL_SIZE"
	mongoSocketTimeout     = "MONGO_SOCKET_TIMEOUT"
	mongoConnectionTimeout = "MONGO_CONNECTION_TIMEOUT"
)
