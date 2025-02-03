package migration

import (
	"database/sql"
	"log/slog"

	"github.com/Flikest/myMicroservices/pkg/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func CreateMigrations(db *sql.DB, migrationsPath string) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	errors.FailOnError(err, "error during migration")

	migration, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	errors.FailOnError(err, "error during migration")

	err = migration.Up()
	errors.FailOnError(err, "эта срань с миграциями!!!!")
	slog.Info("migrations completed")
}
