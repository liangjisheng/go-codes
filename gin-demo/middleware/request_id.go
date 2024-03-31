package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestID is a middleware that injects a 'X-Request-ID' into the context and request/response header of each request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "X-Request-ID"

		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get(key)

		if requestID == "" {
			requestID = uuid.NewV4().String()
		}

		// Expose it for use in the application
		c.Set(key, requestID)

		// Set X-Request-ID header
		c.Writer.Header().Set(key, requestID)
		c.Next()
	}
}
