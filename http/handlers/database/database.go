package database

import (
	"context"

	"github.com/NunChatSpace/identity-service/internal/entities"
	"github.com/gin-gonic/gin"
)

type dbContext struct{}

func FromContext(c context.Context) entities.DB {
	val := c.Value(&dbContext{})
	if val == nil {
		return nil
	}

	return val.(entities.DB)
}

func Handler(db entities.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), &dbContext{}, db)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
