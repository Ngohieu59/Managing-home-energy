package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenRequestId() gin.HandlerFunc { //Assign a unique Request ID to each request.
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.New().String()
			c.Request.Header.Set("X-Request-Id", requestId)
		}
		c.Set("X-Request-Id", requestId)
		c.Next()
	}
}
