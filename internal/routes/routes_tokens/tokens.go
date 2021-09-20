package routes_tokens

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddGroup(r *gin.Engine) {
	tokensGroup := r.Group("/tokens")

	tokensGroup.GET("/", func(c *gin.Context) {
		fmt.Println("tokens")
	})
}
