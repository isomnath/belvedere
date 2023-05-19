package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type PostgresConfigTestSuite struct {
	suite.Suite
	config *PostgresConfig
}

func (suite *PostgresConfigTestSuite) SetupTest() {
	_ = os.Setenv(postgresHost, fmt.Sprintf("%s", "localhost"))
	_ = os.Setenv(postgresPort, fmt.Sprintf("%s", "5432"))
	_ = os.Setenv(postgresDbName, fmt.Sprintf("%s", "dummy_db"))
	_ = os.Setenv(postgresUsername, fmt.Sprintf("%s", "dummy_username"))
	_ = os.Setenv(postgresPassword, fmt.Sprintf("%s", "dummy_password"))
	_ = os.Setenv(postgresPoolSize, fmt.Sprintf("%s", "20"))
	_ = os.Setenv(postgresMigrationsDirectory, fmt.Sprintf("%s", "./pg_migrations"))

	viper.New()
	viper.AutomaticEnv()
	suite.config = postgresConfig()
}

func (suite *PostgresConfigTestSuite) TearDownTest() {
	_ = os.Unsetenv(postgresHost)
	_ = os.Unsetenv(postgresPort)
	_ = os.Unsetenv(postgresDbName)
	_ = os.Unsetenv(postgresUsername)
	_ = os.Unsetenv(postgresPassword)
	_ = os.Unsetenv(postgresPoolSize)
}

func (suite *PostgresConfigTestSuite) TestAllConfigs() {
	suite.Equal("localhost", suite.config.Host())
	suite.Equal(5432, suite.config.Port())
	suite.Equal("dummy_db", suite.config.DbName())
	suite.Equal("dummy_username", suite.config.Username())
	suite.Equal("dummy_password", suite.config.Password())
	suite.Equal(20, suite.config.PoolSize())
	suite.Equal("./pg_migrations", suite.config.MigrationsDirectory())
	suite.Equal("dbname=dummy_db user=dummy_username password=dummy_password host=localhost port=5432 sslmode=disable", suite.config.ConnectionString())
	suite.Equal("postgres://dummy_username:dummy_password@localhost:5432/dummy_db?sslmode=disable", suite.config.ConnectionURL())
}

func TestPostgresConfigTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresConfigTestSuite))
}
