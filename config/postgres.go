package config

import "fmt"

type PostgresConfig struct {
	host                string
	port                int
	dbname              string
	username            string
	password            string
	maxPoolSize         int
	migrationsDirectory string
}

func postgresConfig() *PostgresConfig {
	return &PostgresConfig{
		host:                getString("POSTGRES_HOST", false),
		port:                getInt("POSTGRES_PORT", false),
		dbname:              getString("POSTGRES_DB_NAME", false),
		username:            getString("POSTGRES_USERNAME", false),
		password:            getString("POSTGRES_PASSWORD", false),
		maxPoolSize:         getInt("POSTGRES_POOL_SIZE", false),
		migrationsDirectory: getString("POSTGRES_MIGRATIONS_DIRECTORY", false),
	}
}

func (pg *PostgresConfig) Host() string {
	return pg.host
}

func (pg *PostgresConfig) Port() int {
	return pg.port
}

func (pg *PostgresConfig) DbName() string {
	return pg.dbname
}

func (pg *PostgresConfig) Username() string {
	return pg.username
}

func (pg *PostgresConfig) Password() string {
	return pg.password
}

func (pg *PostgresConfig) PoolSize() int {
	return pg.maxPoolSize
}

func (pg *PostgresConfig) MigrationsDirectory() string {
	return pg.migrationsDirectory
}

func (pg *PostgresConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable", pg.dbname, pg.username, pg.password, pg.host, pg.port)
}

func (pg *PostgresConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", pg.username, pg.password, pg.host, pg.port, pg.dbname)
}
