package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// MiddlewareFactoryContract is the form of the middleware factory
type MiddlewareFactoryContract func(container *dig.Container) (gin.HandlerFunc, error)
