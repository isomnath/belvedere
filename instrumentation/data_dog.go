package instrumentation

import (
	"fmt"
	"log"
	"time"

	"github.com/isomnath/belvedere/config"

	"github.com/DataDog/datadog-go/v5/statsd"
)

var dataDog *statsd.Client

func InitializeDataDogClient(appName string, config *config.DataDogConfig) {
	const (
		retryIntervalSeconds = 2
		maxRetries           = 5
	)

	var err error
	retries := 0

	for {
		dataDog, err = statsd.New(fmt.Sprintf("%s:%d", config.Host(), config.Port()),
			statsd.WithBufferFlushInterval(config.FlushPeriod()),
			statsd.WithNamespace(appName))
		if err != nil {
			if retries < maxRetries {
				retries++
				time.Sleep(retryIntervalSeconds * time.Second)
				continue
			}
			log.Panicf("error initializing data dog client: %v", err)
		}
		break
	}
}

func GetDataDogClient() *statsd.Client {
	return dataDog
}

func CloseDataDogDClient() {
	if dataDog != nil {
		_ = dataDog.Close()
	}
}

func recordHTTPStats(metricCounter, metricTiming string, status int, path, method string, duration float64) {
	if dataDog == nil {
		return
	}

	statusTag := fmt.Sprintf("%s.%d", "status", status)
	methodTag := fmt.Sprintf("%s.%s", "method", method)
	pathTag := fmt.Sprintf("%s.%s", "path", path)

	_ = dataDog.Incr(metricCounter, []string{statusTag, methodTag, pathTag}, 1)
	_ = dataDog.Histogram(metricTiming, duration, []string{methodTag, pathTag}, 1)
}

// RecordInboundHTTPStat records inbound http request/base stats
func RecordInboundHTTPStat(status int, path, method string, duration float64) {
	const (
		metricCounter = "http.inbound"
		metricTiming  = "http.inbound.time"
	)

	recordHTTPStats(metricCounter, metricTiming, status, path, method, duration)
}

// RecordOutboundHTTPStat records outbound http request/base stats
func RecordOutboundHTTPStat(status int, path, method string, duration float64) {
	const (
		metricCounter = "http.outbound"
		metricTiming  = "http.outbound.time"
	)
	recordHTTPStats(metricCounter, metricTiming, status, path, method, duration)
}

func recordPostgresQueryStat(schema, table, queryAlias, status string, duration float64) {
	const (
		metricCounter = "postgres"
		metricTiming  = "postgres.time"
	)

	if dataDog == nil {
		return
	}

	if schema == "" {
		schema = "public"
	}

	schemaTag := fmt.Sprintf("%s.%s", "schema", schema)
	tableTag := fmt.Sprintf("%s.%s", "table", table)
	statusTag := fmt.Sprintf("%s.%s", "status", status)
	queryAliasTag := fmt.Sprintf("%s.%s", "queryAlias", queryAlias)

	_ = dataDog.Incr(metricCounter, []string{schemaTag, tableTag, statusTag, queryAliasTag}, 1)
	_ = dataDog.Histogram(metricTiming, duration, []string{schemaTag, tableTag, queryAliasTag}, 1)
}

// RecordPostgresSuccessStat records postgres success query stats
func RecordPostgresSuccessStat(schema, table, queryAlias string, duration float64) {
	recordPostgresQueryStat(schema, table, queryAlias, "success", duration)
}

// RecordPostgresErrorStat records postgres error query stats
func RecordPostgresErrorStat(schema, table, queryAlias string, duration float64) {
	recordPostgresQueryStat(schema, table, queryAlias, "error", duration)
}

func recordMongoQueryStat(collection, queryAlias, status string, duration float64) {
	const (
		metricCounter = "mongo"
		metricTiming  = "mongo.time"
	)

	if dataDog == nil {
		return
	}

	collectionTag := fmt.Sprintf("%s.%s", "collection", collection)
	statusTag := fmt.Sprintf("%s.%s", "status", status)
	queryAliasTag := fmt.Sprintf("%s.%s", "queryAlias", queryAlias)

	_ = dataDog.Incr(metricCounter, []string{collectionTag, statusTag, queryAliasTag}, 1)
	_ = dataDog.Histogram(metricTiming, duration, []string{collectionTag, queryAliasTag}, 1)
}

// RecordMongoSuccessStat records mongo success query stats
func RecordMongoSuccessStat(collection, queryAlias string, duration float64) {
	recordMongoQueryStat(collection, queryAlias, "success", duration)
}

// RecordMongoErrorStat records mongo error query stats
func RecordMongoErrorStat(collection, queryAlias string, duration float64) {
	recordMongoQueryStat(collection, queryAlias, "error", duration)
}

func recordRedisQueryStat(databaseID int, queryAlias, status string, duration float64) {
	const (
		metricCounter = "redis"
		metricTiming  = "redis.time"
	)

	if dataDog == nil {
		return
	}

	databaseIDTag := fmt.Sprintf("%s.%d", "databaseID", databaseID)
	statusTag := fmt.Sprintf("%s.%s", "status", status)
	queryAliasTag := fmt.Sprintf("%s.%s", "queryAlias", queryAlias)

	_ = dataDog.Incr(metricCounter, []string{databaseIDTag, statusTag, queryAliasTag}, 1)
	_ = dataDog.Histogram(metricTiming, duration, []string{databaseIDTag, queryAliasTag}, 1)
}

// RecordRedisSuccessStat records redis success query stats
func RecordRedisSuccessStat(databaseID int, queryAlias string, duration float64) {
	recordRedisQueryStat(databaseID, queryAlias, "success", duration)
}

// RecordRedisErrorStat records mongo error query stats
func RecordRedisErrorStat(databaseID int, queryAlias string, duration float64) {
	recordRedisQueryStat(databaseID, queryAlias, "error", duration)
}
