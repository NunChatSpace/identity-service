package database

import (
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func NewMigrate(db *gorm.DB) (*migrate.Migrate, error) {
	database, _ := db.DB()
	dbDriver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	_, filename, _, _ := runtime.Caller(0)
	migrationsPath := path.Join(path.Dir(filename), "../../internal/migrations")

	m, err := migrate.NewWithDatabaseInstance("file:///"+migrationsPath, "postgres", dbDriver)
	if err != nil {
		return nil, err
	}

	return m, nil
}
