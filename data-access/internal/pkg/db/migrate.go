package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate will run the migrations within the
// ./migrations folder. Taken directly from
// the golang-migrate documentation:
// https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
func Migrate() error {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@postgres:5432/pets?sslmode=disable")
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}
