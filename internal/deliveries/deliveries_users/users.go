package deliveries_users

import (
	"fmt"
	"net/http"

	"github.com/NunChatSpace/identity-service/internal/cryptography"
	"github.com/NunChatSpace/identity-service/internal/entities"
	"github.com/NunChatSpace/identity-service/internal/response_wrapper"
)

type UserRegisterModel struct {
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

func Register(db entities.DB, model UserRegisterModel) (response_wrapper.Model, error) {
	rolename, err := db.RoleName().Get("user")
	if err != nil {
		return response_wrapper.Model{
			ErrorCode: http.StatusInternalServerError,
			Data:      UserRegisterModel{},
			Message:   "",
		}, err
	}

	cinfo := entities.ContactModel{
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
		Address:     model.Address,
	}
	contact, err := db.Contact().Add(cinfo)
	if err != nil {
		return response_wrapper.Model{
			ErrorCode: http.StatusInternalServerError,
			Data:      UserRegisterModel{},
			Message:   "",
		}, err
	}

	ep, err := cryptography.Encrypt(model.Password)
	if err != nil {
		return response_wrapper.Model{
			ErrorCode: http.StatusInternalServerError,
			Data:      UserRegisterModel{},
			Message:   "",
		}, err
	}

	fmt.Println("encrypted password: ", ep)

	uinfo := entities.UserModel{
		FirstName:  model.FirstName,
		MiddleName: model.MiddleName,
		LastName:   model.LastName,
		ContactID:  contact.ID,
		Password:   ep,
		RoleNameID: rolename.ID,
	}

	_, err = db.User().Add(uinfo)
	if err != nil {
		return response_wrapper.Model{
			ErrorCode: http.StatusInternalServerError,
			Data:      UserRegisterModel{},
			Message:   "",
		}, err
	}

	return response_wrapper.Model{
		ErrorCode: http.StatusOK,
		Data:      model,
		Message:   "",
	}, nil
}
