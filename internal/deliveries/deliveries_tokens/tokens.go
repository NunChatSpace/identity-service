package deliveries_tokens

import (
	"errors"
	"net/http"
	"time"

	"github.com/NunChatSpace/identity-service/internal/cryptography"
	"github.com/NunChatSpace/identity-service/internal/deliveries"
	"github.com/NunChatSpace/identity-service/internal/entities"
	"github.com/NunChatSpace/identity-service/internal/jwt_token"
	"github.com/NunChatSpace/identity-service/internal/response_wrapper"
)

type TokenModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignInModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenModel struct {
	Token string `json:"token"`
}

func GetToken(db entities.DB, model SignInModel) (response_wrapper.Model, error) {
	result, err := getUser(db, model)
	if err != nil {
		return result, err
	}

	user := result.Data.(entities.UserModel)
	roles, err := db.Role().Get(user.RoleNameID)
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}

	permission, err := db.Permission().Get(getPermissionIDs(roles))
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}

	perms := getPermissionNames(permission)
	accessToken, err := jwt_token.CreateJWToken(user, perms, "user_access", time.Now().Add(time.Minute*15).Unix())
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}
	refreshToken, err := jwt_token.CreateJWToken(user, nil, "refresh_access", time.Now().Add(time.Minute*60).Unix())
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}

	return response_wrapper.Model{
		ErrorCode: http.StatusOK,
		Data: TokenModel{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		Message: "",
	}, nil
}

func RefreshToken(db entities.DB, token string) (response_wrapper.Model, error) {
	rftoken, err := jwt_token.Decode(token)
	if err != nil {
		return deliveries.Forbidden(nil, err)
	}

	if rftoken.Type != "refresh_access" {
		return deliveries.Forbidden(nil, errors.New("invalid token type"))
	}

	um := entities.UserModel{
		Model: entities.Model{
			ID: rftoken.UserID,
		},
	}

	user, err := db.User().Get(um)
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}

	roles, err := db.Role().Get(user.RoleNameID)
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}

	permission, err := db.Permission().Get(getPermissionIDs(roles))
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}

	perms := getPermissionNames(permission)
	accessToken, err := jwt_token.CreateJWToken(user, perms, "user_access", time.Now().Add(time.Minute*15).Unix())
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}

	refreshToken, err := jwt_token.CreateJWToken(user, nil, "refresh_access", time.Now().Add(time.Minute*60).Unix())
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}

	return response_wrapper.Model{
		ErrorCode: http.StatusOK,
		Data: TokenModel{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		Message: "",
	}, nil
}

func getUser(db entities.DB, model SignInModel) (response_wrapper.Model, error) {
	contact, err := db.Contact().Get(entities.ContactModel{
		Email: model.Email,
	})
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}
	if contact.ID == "" {
		return deliveries.Forbidden(TokenModel{}, errors.New("user does not exist"))
	}

	user, err := db.User().Get(entities.UserModel{
		ContactID: contact.ID,
	})
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}
	if user.ID == "" {
		return deliveries.Forbidden(TokenModel{}, errors.New("user does not exist"))
	}

	dp, err := cryptography.Decrypt(user.Password)
	if err != nil {
		return deliveries.InternalError(TokenModel{}, err)
	}
	if dp != model.Password {
		return deliveries.Forbidden(TokenModel{}, errors.New("invalid email or password"))
	}

	return response_wrapper.Model{
		ErrorCode: http.StatusOK,
		Data:      user,
		Message:   "",
	}, nil
}

func getPermissionIDs(rs []entities.RoleModel) []string {
	var result []string

	for _, v := range rs {
		result = append(result, v.PermissionID)
	}

	return result
}

func getPermissionNames(pm []entities.PermissionModel) []string {
	var result []string

	for _, v := range pm {
		result = append(result, v.Name)
	}

	return result
}
