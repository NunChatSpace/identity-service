package entities

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	gorm *gorm.DB
}

func (d Db) Ping() error {
	db, err := d.gorm.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}

type DB interface {
	User(ctx context.Context) UserInterface
}

func NewDB() (DB, error) {
	conn := postgres.Config{
		DSN: "host=" + "localhost" + " user=" + "postgres" + " password=" + "postgres" + " dbname=" + "id_service" + " port=" + "5432" + " sslmode=disable TimeZone=" + "Asia/Manila",
	}

	db, err := gorm.Open(postgres.New(conn))
	return Db{
		gorm: db,
	}, err
}
