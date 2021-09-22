package http

import (
	"github.com/NunChatSpace/identity-service/http/handlers"
	"github.com/NunChatSpace/identity-service/internal/entities"
	"github.com/NunChatSpace/identity-service/internal/routes/routes_tokens"
	"github.com/NunChatSpace/identity-service/internal/routes/routes_users"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	db, err := entities.NewDB()
	if err != nil {
		panic(err)
	}

	router := handlers.SetHandlers(db)
	routes_users.AddGroup(router)
	routes_tokens.AddGroup(router)

	return router
}
