package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ianprogrammer/go-api-integration-test/config"
	"github.com/labstack/gommon/log"
)

func MigrateDB(database config.DatabaseConfig, path string) error {

	ds := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", database.UserName, database.Password, database.Host, database.DatabasePort, database.DatabaseName)
	m, err := migrate.New(
		path,
		ds)

	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Info("nothing change in migration")
			return nil
		}
		return err
	}

	return nil

}
