package jwt_token

import (
	"errors"

	"github.com/NunChatSpace/identity-service/internal/entities"
	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	jwt.StandardClaims
	Permission []string `json:"permission"`
	Type       string   `json:"type"`
	UserID     string   `json:"user_id"`
	VerifyCode string   `json:"verify_code"`
}

func CreateJWToken(um entities.UserModel, pm []string, typeStr string, exp int64, verifyCode string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = um.ID
	atClaims["permission"] = pm
	atClaims["type"] = typeStr
	atClaims["exp"] = exp
	atClaims["verify_code"] = verifyCode
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ThisIsSecretKey@ForJWTToken"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func Decode(refreshToken string) (TokenClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("ThisIsSecretKey@ForJWTToken"), nil
	})
	if err != nil {
		return TokenClaims{}, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return TokenClaims{}, errors.New("jwt parsing error")
	}

	return *claims, nil
}

func RefreshToken(refreshToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("ThisIsSecretKey@ForJWTToken"), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", "", errors.New("jwt parsing error")
	}
	if claims.Type != "refresh_access" {
		return "", "", errors.New("invalid token type")
	}

	return "", "", nil
}
