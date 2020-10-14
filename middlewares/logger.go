package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yz89120/geo-ip/providers/t"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

// Logger add logger to context
func Logger(container *dig.Container) (gin.HandlerFunc, error) {
	var logger *zap.Logger

	// Get base logger from container
	if err := container.Invoke(func(l t.Logger) {
		logger = l
	}); err != nil {
		return nil, err
	}

	if logger == nil {
		return nil, errors.New("Logger")
	}

	// the handler
	return func(c *gin.Context) {
		// If request ID is set
		if requestID, ok := c.Get("request_id"); ok {
			if r, ok := requestID.(string); ok {
				// log with request ID
				logger = logger.With(zap.String("request_id", r))
			}
		}

		c.Set("logger", logger)
	}, nil
}
