package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RequestIDHeader set request ID to header `X-Request-ID`
func RequestIDHeader(_ *dig.Container) (gin.HandlerFunc, error) {
	return func(c *gin.Context) {
		if requestID, ok := c.Get("request_id"); ok {
			c.Header("X-Request-ID", requestID.(string))
		}
	}, nil
}
