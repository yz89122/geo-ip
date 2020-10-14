package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/yz89120/geo-ip/utils"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

// LogRequest add logger to context
func LogRequest(_ *dig.Container) (gin.HandlerFunc, error) {
	return func(c *gin.Context) {
		var logger *zap.Logger

		if l, ok := utils.ContextGetLogger(c); ok {
			logger = l
		} else {
			return
		}

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path += "?" + raw
		}

		logger.Info(
			"Request",
			zap.String("remote_addr", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
		)
	}, nil
}
