package routes_users

import (
	"github.com/gin-gonic/gin"
)

func AddGroup(r *gin.Engine) {
	usersGroup := r.Group("/users")

	usersGroup.GET("/", getUser)
}

func getUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": "nun",
	})
}
