//go:generate go run github.com/kyleconroy/sqlc/cmd/sqlc generate

package db

import (
	"embed"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed sql/migrations/*.sql
var fs embed.FS

func UpgradeDatabase(databaseUrl string) error {
	d, err := iofs.New(fs, "sql/migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, databaseUrl)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
