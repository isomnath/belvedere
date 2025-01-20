# Belvedere

---

![Workflow Status](https://github.com/isomnath/belvedere/actions/workflows/main.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/isomnath/belvedere)](https://goreportcard.com/report/github.com/isomnath/belvedere)

<div align="center">
    <img src="assets/belvedere.jpeg" title="Belvedere" alt="Belvedere Logo">
</div>

Full-featured library to bootstrap golang applications lightning quick.

- [Features](#features)
- [Development](#development)
- [Installation & usage](#installation--usage)


## Development ##
#### Local Dev Setup ###
Pre-Requisites:
1. Add `$GOPATH/bin` and `$GOROOT/bin` to `$PATH` variable

Run Following Commands to set up local dev environment
1. `make setup` - To download all pre-requisites
2. `make all`

## Features ##

Consider Belvedere as your applications' butler providing the following capabilities:

- Configuration Management
    - Basic Application Configurations
    - Data Store Connectivity Configurations (Optional)
        - PostgreSQL
        - MongoDB
        - Redis
    - Instrumentation Connectivity Configurations (Optional)
        - DataDog
        - Sentry
        - New Relic
    - i18N Translation Configurations (Optional)
        - Locale based translations
- Instrumentation Support
    - StatsD
        - HTTP Recorders
        - Postgres Query Recorders
        - Mongo Query Recorders
        - Redis Statement Recorders
- Data Store Connectivity Support
    - PostgreSQL
    - MongoDB
    - Redis
- JSON Formatted Logging - With contextual logging 
- Console Based Application Support
  - Postgres Migrations Support - Up and Down
  - MongoDB Index Creation Support
  - Web Server Starter Kit 
    - Logger and instrumentation wrapping
    - Global Panic Recovery
    - Graceful server shutdown 

## Installation & Usage ##

Install Belvedere with:

```sh
go get -u github.com/isomnath/belvedere
```

Sample usage of configuration package.

#### Configuration Usage

```yaml
# Mandatory Configs
APP_NAME
APP_VERSION
APP_ENVIRONMENT
APP_LOG_LEVEL

# Optional Configs
APP_WEB_PORT
APP_NON_WEB_PORT
APP_HEALTH_CHECK_API_PATH
APP_SWAGGER_ENABLED
APP_SWAGGER_DOCS_DIRECTORY

POSTGRES_HOST
POSTGRES_PORT
POSTGRES_DB_NAME
POSTGRES_USERNAME
POSTGRES_PASSWORD
POSTGRES_POOL_SIZE
POSTGRES_MIGRATIONS_DIRECTORY

MONGO_HOSTS
MONGO_DB_NAME
MONGO_USERNAME
MONGO_PASSWORD
MONGO_POOL_SIZE
MONGO_SOCKET_TIMEOUT
MONGO_CONNECTION_TIMEOUT

REDIS_HOST
REDIS_PORT
REDIS_USERNAME
REDIS_PASSWORD

NEW_RELIC_ENABLED
NEW_RELIC_LICENSE_KEY

SENTRY_ENABLED
SENTRY_DSN

DATA_DOG_ENABLED
DATA_DOG_HOST
DATA_DOG_PORT
DATA_DOG_FLUSH_PERIOD_SECONDS

TRANSLATIONS_ENABLED
TRANSLATIONS_PATH
TRANSLATIONS_WHITELISTED_LOCALES
TRANSLATIONS_DEFAULT_LOCALE
```

```go
package main

import "github.com/isomnath/belvedere/config"

func dummyFunc() {
	type CustomConfig struct {
		KeyOne   int      `mapstructure:"KEY_ONE"`
		KeyTwo   string   `mapstructure:"KEY_TWO"`
		KeyThree bool     `mapstructure:"KEY_THREE"`
		KeyFour  []string `mapstructure:"KEY_FOUR"`
		KeyFive  []int    `mapstructure:"KEY_FIVE"`
	}
	var custom CustomConfig
	config.LoadBaseConfig()         // Load Base Application Config - Mandatory
	config.LoadTranslationsConfig() // Load Translation Config - Optional(to support i18N language support)
	config.LoadPostgresConfig()     // Load PostgresDB Connection Configs - Optional
	config.LoadMongoConfig()        // Load MongoDB Connection Configs - Optional
	config.LoadRedisConfig()        // Load Redis Connection Configs - Optional
	config.LoadCustomConfig(custom) // Load Custom Connection Configs - Optional

	// Get Base Application Configs
	config.GetAppName()
	config.GetAppVersion()
	config.GetAppEnvironment()
	config.GetAppWebPort()
	config.GetAppHealthCheckAPIPath()
	config.GetAppLogLevel()
	config.GetAppNonWebPort()
	config.GetSwaggerEnabled()
	config.GetSwaggerDocsDirectory()

	// Get Instrumentation Configs
	config.GetDataDogConfig()
	config.GetSentryConfig()
	config.GetNewRelicConfig()

	// Get Data Store Configs
	config.GetPostgresConfig()
	config.GetMongoConfig()
	config.GetRedisConfig()

	// Get Translations Configs
	config.GetTranslationConfig()

	// Get Custom Configs
	config.GetCustomConfig()
}
```

#### Logger Usage

```go
package main

import (
	"net/http"

	"github.com/isomnath/belvedere/log"
)

func dummyFn() {
	// Pre-requisite - Load Base Config
	log.Setup() //To initialize the logger

	format := "prefix %s and suffix %s"
	args1 := "args1"
	args2 := "args2"
	request := http.Request{}

	log.Logger.Fatalf(format, args1, args2)
	log.Logger.Errorf(format, args1, args2)
	log.Logger.Infof(format, args1, args2)
	log.Logger.Warnf(format, args1, args2)

	log.Logger.HTTPErrorf(&request, format, args1, args2)
	log.Logger.HTTPInfof(&request, format, args1, args2)
	log.Logger.HTTPWarnf(&request, format, args1, args2)

	log.Logger.PostgresErrorf(format, args1, args2)
	log.Logger.PostgresInfof(format, args1, args2)
	log.Logger.PostgresWarnf(format, args1, args2)

	log.Logger.MongoErrorf(format, args1, args2)
	log.Logger.MongoInfof(format, args1, args2)
	log.Logger.MongoWarnf(format, args1, args2)

	log.Logger.RedisErrorf(format, args1, args2)
	log.Logger.RedisInfof(format, args1, args2)
	log.Logger.RedisWarnf(format, args1, args2)
}
```

#### Data Store Usage

```go
package main

import (
	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/console"
	"github.com/isomnath/belvedere/store"
)

func dummyFn() {
	config.LoadBaseConfig()
	config.LoadPostgresConfig()
	config.LoadMongoConfig()
	config.LoadRedisConfig()

	store.PostgresConnect(config.GetPostgresConfig()) // Connects to Postgres DB server
	store.GetPostgresClient()                         // Returns Postgres Client
	store.MongoConnect(config.GetMongoConfig())       // Connects to Mongo DB server
	store.GetMongoClient()                            // Returns Mongo Client
	store.RedisConnect(config.GetRedisConfig(), 0)    // Connects to Redis server
	store.GetRedisClient()                            // Returns Redis Client
}


```

#### Instrumentation Usage

```go
package main

import (
	"errors"
	
	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/instrumentation"
)

func dummyFn() {
	config.LoadBaseConfig()

	// Init functions
	instrumentation.InitializeDataDogClient(config.GetAppName(), config.GetDataDogConfig())
	instrumentation.InitializeSentry(config.GetSentryConfig())

	// Get Clients - Use the clients to raise custom metrics
	instrumentation.GetDataDogClient()
	instrumentation.GetSentryClient()

	// Data Dog Wrappers
	instrumentation.RecordInboundHTTPStat(200, "/path", "GET", 10.2)
	instrumentation.RecordOutboundHTTPStat(200, "/path", "GET", 10.2)

	instrumentation.RecordPostgresSuccessStat("public", "test_table", "GET_RECORD", 3.2)
	instrumentation.RecordPostgresErrorStat("public", "test_table", "GET_RECORD", 3.2)

	instrumentation.RecordMongoSuccessStat("test_collection", "INSERT_RECORD", 1.2)
	instrumentation.RecordMongoErrorStat("test_collection", "INSERT_RECORD", 1.2)

	instrumentation.RecordRedisSuccessStat(0, "SET_KEY", 1.2)
	instrumentation.RecordRedisErrorStat(0, "SET_KEY", 1.2)

	// Sentry Wrappers
	instrumentation.CaptureError(errors.New("some error"))
	instrumentation.CaptureErrorWithTags(errors.New("some error"), map[string]string{"tag_1": "value_1"})
	instrumentation.CaptureWarn(errors.New("some error"))
}
```

#### Instrumentation Usage

```go
package main

import (
	"errors"
	
	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/instrumentation"
)

func dummyFn() {
	config.LoadBaseConfig()

	// Init functions
	instrumentation.InitializeDataDogClient(config.GetAppName(), config.GetDataDogConfig())
	instrumentation.InitializeSentry(config.GetSentryConfig())

	// Get Clients - Use the clients to raise custom metrics
	instrumentation.GetDataDogClient()
	instrumentation.GetSentryClient()

	// Data Dog Wrappers
	instrumentation.RecordInboundHTTPStat(200, "/path", "GET", 10.2)
	instrumentation.RecordOutboundHTTPStat(200, "/path", "GET", 10.2)

	instrumentation.RecordPostgresSuccessStat("public", "test_table", "GET_RECORD", 3.2)
	instrumentation.RecordPostgresErrorStat("public", "test_table", "GET_RECORD", 3.2)

	instrumentation.RecordMongoSuccessStat("test_collection", "INSERT_RECORD", 1.2)
	instrumentation.RecordMongoErrorStat("test_collection", "INSERT_RECORD", 1.2)

	instrumentation.RecordRedisSuccessStat(0, "SET_KEY", 1.2)
	instrumentation.RecordRedisErrorStat(0, "SET_KEY", 1.2)

	// Sentry Wrappers
	instrumentation.CaptureError(errors.New("some error"))
	instrumentation.CaptureErrorWithTags(errors.New("some error"), map[string]string{"tag_1": "value_1"})
	instrumentation.CaptureWarn(errors.New("some error"))
}
```
