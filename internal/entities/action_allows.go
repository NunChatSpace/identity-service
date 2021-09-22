package entities

import "gorm.io/gorm"

func (d Db) ActionAllow() ActionAllowInterface {
	return actionAllowDb(d)
}

type actionAllowDb struct {
	gorm *gorm.DB
}

type ActionAllowInterface interface {
	Add(info ActionAllowModel) (ActionAllowModel, error)
	Get(uid string) (ActionAllowModel, error)
}

type ActionAllowModel struct {
	Model
	Name string `json:"name"`
}

func (ActionAllowModel) TableName() string {
	return "action_allows"
}

func (aadb actionAllowDb) Add(info ActionAllowModel) (ActionAllowModel, error) {
	tx := aadb.gorm.Create(&info)

	return info, tx.Error
}

func (aadb actionAllowDb) Get(uid string) (ActionAllowModel, error) {
	actionAllow := ActionAllowModel{}
	aadb.gorm.First(&actionAllow, "id = ?", uid)

	return actionAllow, nil
}
