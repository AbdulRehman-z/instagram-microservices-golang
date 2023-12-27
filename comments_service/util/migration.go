package util

import (
	"log"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration(migrationUrl string, dbUrl string) {
	migration, err := migrate.New(migrationUrl, dbUrl)
	if err != nil {
		log.Fatalln("migrate: failed to initialize migration", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalln("migrate: failed to run migration up", err)
	}

	slog.Info("migration completed successfully")
}
