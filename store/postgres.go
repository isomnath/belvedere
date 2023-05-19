package store

import (
	"context"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq" //Needed for postgres DB Connection
	sqlTracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	ddTracerSqlx "gopkg.in/DataDog/dd-trace-go.v1/contrib/jmoiron/sqlx"
)

type PostgresStore struct {
	client *sqlx.DB
}

var pg *PostgresStore

// TODO: Add tests for error blocks
func pgInitialize(cnf *config.PostgresConfig) (*sqlx.DB, error) {
	ctx := context.Background()
	log.Log.PostgresInfof(ctx, "attempting to connect to postgres database: %s at host: %s, port: %d",
		cnf.DbName(), cnf.Host(), cnf.Port())

	sqlTracer.Register("postgres", &pq.Driver{}, sqlTracer.WithServiceName(config.GetAppName()))
	db, err := ddTracerSqlx.MustOpen("postgres", cnf.ConnectionString())
	if err != nil {
		log.Log.PostgresErrorf(ctx, "failed to connect to postgres server: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Log.PostgresErrorf(ctx, "ping to postgres server failed: %v", err)
		return nil, err
	}

	db.SetMaxOpenConns(cnf.PoolSize())
	log.Log.PostgresInfof(ctx, "successfully connected")
	return db, nil
}

// PostgresConnect - Connects to mongo and initializes the DB client to be reused across the application
func PostgresConnect(config *config.PostgresConfig) error {
	pgClient, err := pgInitialize(config)
	if err != nil {
		return err
	}
	pg = &PostgresStore{client: pgClient}
	return err
}

// GetPostgresClient - Returns the instance of postgres DB client created when PostgresConnect was invoked
func GetPostgresClient() *sqlx.DB {
	return pg.client
}
