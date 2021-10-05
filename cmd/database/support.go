package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunCmd(args []string) error {
	switch args[0] {
	case "migrate":
		return dbMigrate()
	case "drop":
		return drop()
	case "export":
		return export()
	default:
		fmt.Println("Error")
	}
	return nil
}

func dbMigrate() error {
	m, err := createMigrate()
	if err != nil {
		return err
	}
	fmt.Println("Migrating...")

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func drop() error {
	m, err := createMigrate()
	if err != nil {
		return err
	}
	fmt.Println("Deleting...")

	return m.Drop()
}

func export() error {
	fmt.Println("export")
	return nil
}

func createMigrate() (*migrate.Migrate, error) {
	conn := postgres.Config{
		DSN: "host=" + "'postgres'" + " user=" + "postgres" + " password=" + "postgres" + " dbname=" + "id_service_db" + " port=" + "5432" + " sslmode=disable TimeZone=" + "Asia/Manila",
	}

	db, err := gorm.Open(postgres.New(conn))
	if err != nil {
		return nil, err
	}

	m, err := NewMigrate(db)
	if err != nil {
		return nil, err
	}

	return m, nil
}
