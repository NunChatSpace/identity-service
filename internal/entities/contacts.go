package entities

import "gorm.io/gorm"

func (d Db) Contact() ContactInterface {
	return contactDb(d)
}

type contactDb struct {
	gorm *gorm.DB
}

type ContactInterface interface {
	Add(info ContactModel) (ContactModel, error)
	Get(uid string) (ContactModel, error)
}

type ContactModel struct {
	Model
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func (ContactModel) TableName() string {
	return "contacts"
}

func (cdb contactDb) Add(info ContactModel) (ContactModel, error) {
	tx := cdb.gorm.Create(&info)

	return info, tx.Error
}

func (cdb contactDb) Get(uid string) (ContactModel, error) {
	contact := ContactModel{}
	cdb.gorm.First(&contact, "id = ?", uid)

	return contact, nil
}
