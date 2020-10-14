package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ContextGetLogger return logger from the context
func ContextGetLogger(c *gin.Context) (logger *zap.Logger, ok bool) {
	if l, exists := c.Get("logger"); exists {
		logger, ok = l.(*zap.Logger)
	}
	return
}
