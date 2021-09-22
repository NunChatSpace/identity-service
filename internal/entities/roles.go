package entities

import "gorm.io/gorm"

func (d Db) Role() RoleInterface {
	return roleDb(d)
}

type roleDb struct {
	gorm *gorm.DB
}

type RoleInterface interface {
	Add(info RoleModel) (RoleModel, error)
	Get(uid string) (RoleModel, error)
}

type RoleModel struct {
	Model
	RoleNameID   string `json:"role_name_id"`
	PermissionID string `json:"permission_id"`
}

func (RoleModel) TableName() string {
	return "roles"
}

func (rdb roleDb) Add(info RoleModel) (RoleModel, error) {
	tx := rdb.gorm.Create(&info)

	return info, tx.Error
}

func (rdb roleDb) Get(uid string) (RoleModel, error) {
	role := RoleModel{}
	rdb.gorm.First(&role, "id = ?", uid)

	return role, nil
}
