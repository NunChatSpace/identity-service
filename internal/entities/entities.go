package entities

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Model struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ToUpper(uuid.New().String())
	}

	return nil
}

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
	ActionAllow() ActionAllowInterface
	Consent() ConsentInterface
	Contact() ContactInterface
	Permission() PermissionInterface
	Role() RoleInterface
	RoleName() RoleNameInterface
	User() UserInterface
	VerifyCode() VerifyCodeInterface
}

func NewDB() (DB, error) {
	conn := postgres.Config{
		// DSN: "postgresql://postgres:postgres@postgres/id_service?sslmode=disable",
		DSN: "host=" + "'postgres'" + " user=" + "postgres" + " password=" + "postgres" + " dbname=" + "id_service_db" + " port=" + "5432" + " sslmode=disable TimeZone=" + "Asia/Manila",
	}

	db, err := gorm.Open(postgres.New(conn))
	return Db{
		gorm: db,
	}, err
}
