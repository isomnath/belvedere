package console

import (
	"context"
	"fmt"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Needed for database driver for migration
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Needed for file source for migration
	_ "github.com/lib/pq"                                      // Needed for driver source
)

func getMigrator() (*migrate.Migrate, error) {
	ctx := context.Background()
	migrationsPath := fmt.Sprintf("file://%s", config.GetPostgresConfig().MigrationsDirectory())
	migrator, err := migrate.New(migrationsPath, config.GetPostgresConfig().ConnectionURL())

	if err != nil {
		log.Log.PostgresErrorf(ctx, "failed to initialize migrator with error: %v", err)
		return nil, err
	}
	return migrator, nil
}

func handleMigratorError(ctx context.Context, err error, errorMsgTemplate string) error {
	if err != nil {
		//TODO: Add tests to cover this error scenario
		if err.Error() != "no change" {
			log.Log.PostgresErrorf(ctx, errorMsgTemplate, err)
			return err
		}
	}
	return nil
}

// RunRollUp - Allow running migrations from configured directory
// Ideally to be used as console command
func RunRollUp() error {
	ctx := context.Background()
	migrator, err := getMigrator()
	if err != nil {
		return err
	}

	err = migrator.Up()

	return handleMigratorError(ctx, err, "failed to run roll ups with error: %v")
}

// RunRollback - Allow running rollbacks from configured directory
// Ideally to be used as console command
func RunRollback() error {
	ctx := context.Background()
	migrator, err := getMigrator()
	if err != nil {
		return err
	}

	err = migrator.Down()

	return handleMigratorError(ctx, err, "failed to run roll backs with error: %v")
}
