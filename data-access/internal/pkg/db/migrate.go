package db

import (
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate will run the migrations within the
// ./migrations folder. Taken directly from
// the golang-migrate documentation:
// https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
func Migrate(uri string) error {
	m, err := migrate.New(
		"file://migrations",
		uri)
	if err != nil {
		return err
	}
	direction := os.Getenv("MIGRATION_DIRECTION")
	if direction == "down" {
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		return nil
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
