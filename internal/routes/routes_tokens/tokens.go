package routes_tokens

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/NunChatSpace/identity-service/http/handlers/database"
	dt "github.com/NunChatSpace/identity-service/internal/deliveries/deliveries_tokens"
	"github.com/NunChatSpace/identity-service/internal/response_wrapper"
	"github.com/gin-gonic/gin"
)

func AddGroup(r *gin.Engine) {
	tokensGroup := r.Group("/tokens")

	tokensGroup.POST("/", getToken)
	tokensGroup.POST("/refresh", refreshToken)
	tokensGroup.GET("/intrspct", introspection)
}

func getToken(c *gin.Context) {
	var signinModel dt.SignInModel

	err := json.NewDecoder(c.Request.Body).Decode(&signinModel)
	if err != nil {
		response_wrapper.Resp(http.StatusInternalServerError, "", err.Error(), c)
		return
	}

	db := database.FromContext(c.Request.Context())
	if db == nil {
		response_wrapper.Resp(http.StatusInternalServerError, "", "db does not exist", c)
		return
	}

	m := dt.GetToken(db, signinModel)
	response_wrapper.Resp(m.StatusCode, m.Data, "", c)
}

func refreshToken(c *gin.Context) {
	auth := c.Request.Header["Authorization"]

	if len(auth) == 0 {
		response_wrapper.Resp(http.StatusForbidden, "", "unauthenticate", c)
		return
	}

	db := database.FromContext(c.Request.Context())
	if db == nil {
		response_wrapper.Resp(http.StatusInternalServerError, "", "db does not exist", c)
		return
	}

	token := strings.Split(auth[0], " ")[1]
	m := dt.RefreshToken(db, token)
	response_wrapper.Resp(m.StatusCode, m.Data, "", c)
}

func introspection(c *gin.Context) {
	auth := c.Request.Header["Authorization"]

	if len(auth) == 0 {
		response_wrapper.Resp(http.StatusForbidden, "", "unauthenticate", c)
		return
	}

	db := database.FromContext(c.Request.Context())
	if db == nil {
		response_wrapper.Resp(http.StatusInternalServerError, "", "db does not exist", c)
		return
	}

	token := strings.Split(auth[0], " ")[1]
	m := dt.IntrospectionToken(db, token)

	response_wrapper.Resp(m.StatusCode, m.Data, "", c)
}
