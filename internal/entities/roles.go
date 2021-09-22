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
	Get(roleNameID string) ([]RoleModel, error)
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

func (rdb roleDb) Get(roleNameID string) ([]RoleModel, error) {
	roles := []RoleModel{}
	rdb.gorm.Find(&roles, "role_name_id = ?", roleNameID)

	return roles, nil
}
