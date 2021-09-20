package http

import (
	"github.com/NunChatSpace/identity-service/http/handlers"
	"github.com/NunChatSpace/identity-service/internal/routes/routes_tokens"
	"github.com/NunChatSpace/identity-service/internal/routes/routes_users"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	router := gin.Default()
	routes_users.AddGroup(router)
	routes_tokens.AddGroup(router)

	handlers.SetHandlers(router)
	return router
}
