package entities

import (
	"gorm.io/gorm"
)

func (d Db) VerifyCode() VerifyCodeInterface {
	return verifyCodeDb(d)
}

type verifyCodeDb struct {
	gorm *gorm.DB
}

type VerifyCodeInterface interface {
	IsValid(userID string, verifyCode string) (bool, error)
	Add(info VerifyCodeModel) (VerifyCodeModel, error)
}

type VerifyCodeModel struct {
	Model
	VerifyCode string `json:"verify_code"`
	UserID     string `json:"user_id"`
}

func (VerifyCodeModel) TableName() string {
	return "verify_codes"
}

func (vcdb verifyCodeDb) IsValid(userID string, verifyCode string) (bool, error) {
	model := VerifyCodeModel{
		VerifyCode: verifyCode,
		UserID:     userID,
	}
	var result VerifyCodeModel
	tx := vcdb.gorm.Model(&model).First(&result)

	return (result.ID != ""), tx.Error
}

func (vcdb verifyCodeDb) Add(info VerifyCodeModel) (VerifyCodeModel, error) {
	tx := vcdb.gorm.Create(&info)

	return info, tx.Error
}
