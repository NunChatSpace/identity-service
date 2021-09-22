package cors

import "github.com/gin-gonic/gin"

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
