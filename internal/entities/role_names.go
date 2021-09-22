package entities

import "gorm.io/gorm"

func (d Db) RoleName() RoleNameInterface {
	return rolenameDb(d)
}

type rolenameDb struct {
	gorm *gorm.DB
}

type RoleNameInterface interface {
	Add(info RoleNameModel) (RoleNameModel, error)
	Get(uid string) (RoleNameModel, error)
}

type RoleNameModel struct {
	Model
	Name string `json:"name"`
}

func (RoleNameModel) TableName() string {
	return "role_names"
}

func (rndb rolenameDb) Add(info RoleNameModel) (RoleNameModel, error) {
	tx := rndb.gorm.Create(&info)

	return info, tx.Error
}

func (rndb rolenameDb) Get(name string) (RoleNameModel, error) {
	rolename := RoleNameModel{}
	rndb.gorm.First(&rolename, "name = ?", name)

	return rolename, nil
}
