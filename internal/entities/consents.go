package entities

import "gorm.io/gorm"

func (d Db) Consent() ConsentInterface {
	return consentDb(d)
}

type consentDb struct {
	gorm *gorm.DB
}

type ConsentInterface interface {
	Add(info ConsentModel) (ConsentModel, error)
	Get(uid string) (ConsentModel, error)
}

type ConsentModel struct {
	Model
	ActionAllowID string `json:"action_allow_id"`
	UserID        string `json:"user_id"`
}

func (ConsentModel) TableName() string {
	return "consents"
}

func (cdb consentDb) Add(info ConsentModel) (ConsentModel, error) {
	tx := cdb.gorm.Create(&info)

	return info, tx.Error
}

func (cdb consentDb) Get(uid string) (ConsentModel, error) {
	consent := ConsentModel{}
	cdb.gorm.First(&consent, "id = ?", uid)

	return consent, nil
}
