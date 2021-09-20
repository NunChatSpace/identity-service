package entities

import (
	"context"

	"gorm.io/gorm"
)

func (d Db) User(ctx context.Context) UserInterface {
	d.gorm = d.gorm.WithContext(ctx)

	return userDb(d)
}

type userDb struct {
	gorm *gorm.DB
}

type UserInterface interface {
	GetUser(uid string) (UserModel, error)
}

type UserModel struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	ContactID  string `json:"contact_id"`
	Password   string `json:"password"`
	RoleID     string `json:"role_id"`
	ConsentID  string `json:"consent_id"`
}

func (udb userDb) GetUser(uid string) (UserModel, error) {
	user := UserModel{}
	udb.gorm.First(&user, "id = ?", uid)

	return user, nil
}
