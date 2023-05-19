package console

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/instrumentation"
	"github.com/isomnath/belvedere/log"
	"github.com/isomnath/belvedere/store"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type PostgresMigrationSupportTestSuite struct {
	suite.Suite
	pgClient  *sqlx.DB
	tableName string
}

func (suite *PostgresMigrationSupportTestSuite) SetupSuite() {
	config.LoadBaseConfig()
	config.LoadPostgresConfig()

	instrumentation.StartDDTracer()
	log.Setup()
}

func (suite *PostgresMigrationSupportTestSuite) TearDownSuite() {
	instrumentation.StopDDTracer()
}

func (suite *PostgresMigrationSupportTestSuite) SetupTest() {
	_ = store.PostgresConnect(config.GetPostgresConfig())
	suite.pgClient = store.GetPostgresClient()
	suite.tableName = "test_table"
}

func (suite *PostgresMigrationSupportTestSuite) TearDownTest() {
	queryOne := fmt.Sprintf("DROP TABLE IF EXISTS %s", suite.tableName)
	queryTwo := fmt.Sprintf("DROP TABLE IF EXISTS %s", "schema_migrations")
	_, _ = suite.pgClient.DB.Exec(queryOne)
	_, _ = suite.pgClient.DB.Exec(queryTwo)
	suite.removeMigrationFilesDirectory()
}

func (suite *PostgresMigrationSupportTestSuite) TestRunRollUpAndRollBackSuccessfully() {
	suite.prepareMigrationFiles()
	type testStruct struct {
		ID      int64  `db:"id"`
		Column1 string `db:"col_1"`
		Column2 string `db:"col_2"`
	}

	err := RunRollUp()
	suite.NoError(err)
	_, err = suite.pgClient.Exec(fmt.Sprintf("INSERT INTO %s (col_1, col_2) VALUES ('value 1', 'value 2');", suite.tableName))
	suite.NoError(err)

	row := suite.pgClient.QueryRow(fmt.Sprintf("SELECT * from %s where col_1='value 1';", suite.tableName))
	suite.NoError(row.Err())

	var dest testStruct
	err = row.Scan(&dest.ID, &dest.Column1, &dest.Column2)
	suite.NoError(err)
	suite.Equal(int64(1), dest.ID)
	suite.Equal("value 1", dest.Column1)
	suite.Equal("value 2", dest.Column2)

	// Test idempotency - RunRollUp shouldn't fail if there's no change
	err = RunRollUp()
	suite.NoError(err)

	err = RunRollback()
	suite.NoError(err)
	_, err = suite.pgClient.Exec(fmt.Sprintf("INSERT INTO %s (col_1, col_2) VALUES ('value 1', 'value 2');", suite.tableName))
	suite.Equal("pq: relation \"test_table\" does not exist", err.Error())

	// Test idempotency - RunRollback shouldn't fail if there's no change
	err = RunRollback()
	suite.NoError(err)
}

func (suite *PostgresMigrationSupportTestSuite) TestRunRollUpsReturnsErrorWhenMigratorCannotBeInitialized() {
	err := RunRollUp()
	suite.NotNil(err)
}

func (suite *PostgresMigrationSupportTestSuite) TestRunRollBacksReturnsErrorWhenMigratorCannotBeInitialized() {
	err := RunRollback()
	suite.NotNil(err)
}

func (suite *PostgresMigrationSupportTestSuite) prepareMigrationFiles() {
	firstUpQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\"id\" BIGSERIAL,\"col_1\" VARCHAR,\"col_2\" VARCHAR);", suite.tableName)
	firstDownQuery := fmt.Sprintf("DROP TABLE IF EXISTS %s;", suite.tableName)
	timestamp := time.Now().Format("20060102150405")
	firstMigrationUpFileName := fmt.Sprintf("%s_create_test_table.up.sql", timestamp)
	firstMigrationDownFileName := fmt.Sprintf("%s_create_test_table.down.sql", timestamp)

	_ = os.Mkdir(config.GetPostgresConfig().MigrationsDirectory(), os.ModePerm)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/%s", config.GetPostgresConfig().MigrationsDirectory(), firstMigrationUpFileName), []byte(firstUpQuery), 0644)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/%s", config.GetPostgresConfig().MigrationsDirectory(), firstMigrationDownFileName), []byte(firstDownQuery), 0644)
}

func (suite *PostgresMigrationSupportTestSuite) removeMigrationFilesDirectory() {
	_ = os.RemoveAll(config.GetPostgresConfig().MigrationsDirectory())
}

func TestPostgresMigrationSupportTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresMigrationSupportTestSuite))
}
