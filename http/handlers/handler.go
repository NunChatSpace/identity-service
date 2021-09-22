package handlers

import (
	"github.com/NunChatSpace/identity-service/http/handlers/cors"
	"github.com/NunChatSpace/identity-service/http/handlers/database"
	"github.com/NunChatSpace/identity-service/internal/entities"
	"github.com/gin-gonic/gin"
)

func SetHandlers(db entities.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Handler())
	router.Use(database.Handler(db))

	return router
}
