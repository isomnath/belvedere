package instrumentation

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/isomnath/belvedere/config"

	nr "github.com/newrelic/go-agent"
)

type nCtxKey int

const txKey nCtxKey = 0

var newrelicApp nr.Application

func InitNewrelic(cfg *config.NewRelicConfig) {
	if cfg.Enabled() {
		var err error
		c := nr.Config{
			Enabled: cfg.Enabled(),
			License: cfg.LicenseKey(),
		}
		newrelicApp, err = nr.NewApplication(c)
		if err != nil {
			log.Fatalf("failed to initialize nr agent: %v", err)
		}
	}
}

func ShutDownNewrelic() {
	if newrelicApp != nil {
		newrelicApp.Shutdown(time.Second)
	}
}

func GetNewrelicAgent() nr.Application {
	return newrelicApp
}

func StartPostgresDataSegmentNow(op string, tableName string, txn nr.Transaction) *nr.DatastoreSegment {
	return startDatastoreSegment(op, tableName, txn, nr.DatastorePostgres)
}

func StartMongoDBDataSegmentNow(op string, collectionName string, txn nr.Transaction) *nr.DatastoreSegment {
	return startDatastoreSegment(op, collectionName, txn, nr.DatastoreMongoDB)
}

func StartCassandraSegmentNow(op string, tableName string, txn nr.Transaction) *nr.DatastoreSegment {
	return startDatastoreSegment(op, tableName, txn, nr.DatastoreCassandra)
}

func StartRedisSegmentNow(op string, tableName string, txn nr.Transaction) *nr.DatastoreSegment {
	return startDatastoreSegment(op, tableName, txn, nr.DatastoreRedis)
}

func startDatastoreSegment(op string, tableName string, txn nr.Transaction, product nr.DatastoreProduct) *nr.DatastoreSegment {
	s := nr.DatastoreSegment{
		Product:    product,
		Collection: tableName,
		Operation:  op,
		StartTime:  nr.StartSegmentNow(txn),
	}

	return &s
}

func StartSegmentNow(name string, txn nr.Transaction) *nr.Segment {
	s := nr.Segment{
		Name:      name,
		StartTime: nr.StartSegmentNow(txn),
	}

	return &s
}

func StartKafkaPushSegment(txn nr.Transaction, topic string) *nr.MessageProducerSegment {
	s := nr.MessageProducerSegment{
		StartTime:            nr.StartSegmentNow(txn),
		Library:              "Kafka",
		DestinationType:      nr.MessageTopic,
		DestinationName:      topic,
		DestinationTemporary: false,
	}

	return &s
}

func StartRabbitmqPushSegment(txn nr.Transaction, exchange string) *nr.MessageProducerSegment {
	s := nr.MessageProducerSegment{
		StartTime:            nr.StartSegmentNow(txn),
		Library:              "RabbitMQ",
		DestinationType:      nr.MessageExchange,
		DestinationName:      exchange,
		DestinationTemporary: false,
	}

	return &s
}

func StartExternalSegmentNow(txn nr.Transaction, url string) *nr.ExternalSegment {
	s := nr.ExternalSegment{
		StartTime: nr.StartSegmentNow(txn),
		URL:       url,
	}

	return &s
}

func NewHTTPContext(ctx context.Context, w http.ResponseWriter) context.Context {
	if newrelicApp != nil {
		tx, ok := w.(nr.Transaction)
		if !ok {
			return ctx
		}
		return context.WithValue(ctx, txKey, tx)
	}
	return ctx
}

func NewContextWithTransaction(ctx context.Context, tx nr.Transaction) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func GetTx(ctx context.Context) (nr.Transaction, bool) {
	tx, ok := ctx.Value(txKey).(nr.Transaction)
	return tx, ok
}
