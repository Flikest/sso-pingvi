package migration

import (
	"database/sql"
	"log/slog"

	"github.com/Flikest/myMicroservices/pkg/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func CreateMigrations(db *sql.DB, migrationsPath string) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	errors.FailOnError(err, "error during migration")
	migration, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	migration.Up()
	slog.Info("migrations completed")
}
