package response_wrapper

import "github.com/gin-gonic/gin"

type Model struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}

func Resp(genericError int, data interface{}, message string, c *gin.Context) {
	c.JSON(genericError, gin.H{
		"data":    data,
		"message": message,
	})
}
