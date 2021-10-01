package deliveries_tokens

import (
	"net/http"
	"strings"
	"time"

	"github.com/NunChatSpace/identity-service/internal/cryptography"
	"github.com/NunChatSpace/identity-service/internal/entities"
	"github.com/NunChatSpace/identity-service/internal/jwt_token"
	"github.com/NunChatSpace/identity-service/internal/response_wrapper"
	"github.com/google/uuid"
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

func GetToken(db entities.DB, model SignInModel) response_wrapper.Model {
	result := getUser(db, model)
	if result.StatusCode != http.StatusOK {
		return result
	}

	user := result.Data.(entities.UserModel)
	roles, err := db.Role().Get(user.RoleNameID)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	permission, err := db.Permission().Get(getPermissionIDs(roles))
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	perms := getPermissionNames(permission)
	verifyCode := strings.ToUpper(uuid.New().String())
	accessToken, err := jwt_token.CreateJWToken(user, perms, "user_access", time.Now().Add(time.Minute*15).Unix(), verifyCode)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}
	refreshToken, err := jwt_token.CreateJWToken(user, nil, "refresh_access", time.Now().Add(time.Minute*60).Unix(), verifyCode)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	vcInfo := entities.VerifyCodeModel{
		VerifyCode: verifyCode,
		UserID:     user.ID,
	}

	_, err = db.VerifyCode().Add(vcInfo)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	return response_wrapper.Model{
		StatusCode: http.StatusOK,
		Data: TokenModel{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		Message: "",
	}
}

func RefreshToken(db entities.DB, token string) response_wrapper.Model {
	rftoken, err := jwt_token.Decode(token)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusForbidden,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	if rftoken.Type != "refresh_access" {
		return response_wrapper.Model{
			StatusCode: http.StatusForbidden,
			Data:       TokenModel{},
			Message:    "invalid token type",
		}
	}

	um := entities.UserModel{
		Model: entities.Model{
			ID: rftoken.UserID,
		},
	}

	user, err := db.User().Get(um)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	roles, err := db.Role().Get(user.RoleNameID)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	permission, err := db.Permission().Get(getPermissionIDs(roles))
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	perms := getPermissionNames(permission)
	verifyCode := strings.ToUpper(uuid.New().String())
	accessToken, err := jwt_token.CreateJWToken(user, perms, "user_access", time.Now().Add(time.Minute*15).Unix(), verifyCode)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	refreshToken, err := jwt_token.CreateJWToken(user, nil, "refresh_access", time.Now().Add(time.Minute*60).Unix(), verifyCode)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	vcInfo := entities.VerifyCodeModel{
		VerifyCode: verifyCode,
		UserID:     user.ID,
	}

	_, err = db.VerifyCode().Add(vcInfo)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Data:       TokenModel{},
			Message:    err.Error(),
		}
	}

	return response_wrapper.Model{
		StatusCode: http.StatusOK,
		Data: TokenModel{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		Message: "",
	}
}

func IntrospectionToken(db entities.DB, token string) response_wrapper.Model {
	accessToken, err := jwt_token.Decode(token)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
		}
	}

	if accessToken.Type != "user_access" {
		return response_wrapper.Model{
			StatusCode: http.StatusForbidden,
			Message:    "invalid token type",
		}
	}

	isValid, err := db.VerifyCode().IsValid(accessToken.UserID, accessToken.VerifyCode)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if !isValid {
		return response_wrapper.Model{
			StatusCode: http.StatusForbidden,
			Message:    "unauthenticate",
		}
	}

	return response_wrapper.Model{
		StatusCode: http.StatusAccepted,
		Data:       accessToken,
	}
}

func getUser(db entities.DB, model SignInModel) response_wrapper.Model {
	contact, err := db.Contact().Get(entities.ContactModel{
		Email: model.Email,
	})
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	if contact.ID == "" {
		return response_wrapper.Model{
			StatusCode: http.StatusForbidden,
			Message:    "user does not exist",
		}
	}

	user, err := db.User().Get(entities.UserModel{
		ContactID: contact.ID,
	})
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	if user.ID == "" {
		return response_wrapper.Model{
			StatusCode: http.StatusForbidden,
			Message:    "user does not exist",
		}
	}

	dp, err := cryptography.Decrypt(user.Password)
	if err != nil {
		return response_wrapper.Model{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	if dp != model.Password {
		return response_wrapper.Model{
			StatusCode: http.StatusForbidden,
			Message:    "invalid email or password",
		}
	}

	return response_wrapper.Model{
		StatusCode: http.StatusOK,
		Data:       user,
		Message:    "",
	}
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
