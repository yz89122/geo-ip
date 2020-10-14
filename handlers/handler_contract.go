package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// HandlerFactoryContract is the form the handler factory
type HandlerFactoryContract func(*dig.Container) (gin.HandlerFunc, error)
