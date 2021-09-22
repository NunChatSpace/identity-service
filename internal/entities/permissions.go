package entities

import "gorm.io/gorm"

func (d Db) Permission() PermissionInterface {
	return permissionDb(d)
}

type permissionDb struct {
	gorm *gorm.DB
}

type PermissionInterface interface {
	Add(info PermissionModel) (PermissionModel, error)
	Get(uid []string) ([]PermissionModel, error)
}

type PermissionModel struct {
	Model
	Name string `json:"name"`
}

func (PermissionModel) TableName() string {
	return "permissions"
}

func (pdb permissionDb) Add(info PermissionModel) (PermissionModel, error) {
	tx := pdb.gorm.Create(&info)

	return info, tx.Error
}

func (pdb permissionDb) Get(uid []string) ([]PermissionModel, error) {
	permission := []PermissionModel{}
	pdb.gorm.Find(&permission, "id IN ?", uid)

	return permission, nil
}
