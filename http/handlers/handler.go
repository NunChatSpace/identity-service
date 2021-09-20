package handlers

import (
	"github.com/NunChatSpace/identity-service/http/handlers/cors"
	"github.com/gin-gonic/gin"
)

func SetHandlers(r *gin.Engine) {
	r.Use(cors.Handler())
}
