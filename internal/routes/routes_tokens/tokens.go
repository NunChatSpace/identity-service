package routes_tokens

import (
	"encoding/json"
	"net/http"

	"github.com/NunChatSpace/identity-service/http/handlers/database"
	dt "github.com/NunChatSpace/identity-service/internal/deliveries/deliveries_tokens"
	"github.com/NunChatSpace/identity-service/internal/response_wrapper"
	"github.com/gin-gonic/gin"
)

func AddGroup(r *gin.Engine) {
	tokensGroup := r.Group("/tokens")

	tokensGroup.POST("/", getToken)
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

	m, err := dt.GetToken(db, signinModel)
	if err != nil {
		response_wrapper.Resp(m.ErrorCode, m.Data, err.Error(), c)
		return
	}
	response_wrapper.Resp(m.ErrorCode, m.Data, "", c)
}
