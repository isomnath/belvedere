package store

import (
	"os"
	"testing"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/instrumentation"
	"github.com/isomnath/belvedere/log"

	"github.com/stretchr/testify/suite"
)

type PostgresTestSuite struct {
	suite.Suite
}

func (suite *PostgresTestSuite) SetupSuite() {
	config.LoadBaseConfig()
	instrumentation.StartDDTracer()
	defer instrumentation.StopDDTracer()
	log.Setup()
}

func (suite *PostgresTestSuite) TestPostgresConnectionSuccess() {
	_ = os.Setenv("POSTGRES_HOST", "localhost")
	config.LoadPostgresConfig()
	pgConfig := config.GetPostgresConfig()
	err := PostgresConnect(pgConfig)
	suite.NoError(err)
	suite.NotNil(GetPostgresClient())
}

func (suite *PostgresTestSuite) TestPostgresPingError() {
	_ = os.Setenv("POSTGRES_HOST", "invalid")
	defer os.Unsetenv("POSTGRES_HOST")
	config.LoadPostgresConfig()
	pgConfig := config.GetPostgresConfig()
	err := PostgresConnect(pgConfig)
	suite.Contains(err.Error(), "dial tcp: lookup invalid")
}

func TestPostgresTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}
