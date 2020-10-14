package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/yz89120/geo-ip/config"
	"github.com/yz89120/geo-ip/providers/t"
	"go.uber.org/dig"
)

// RequestID create a handler for setting request ID
func RequestID(container *dig.Container) (gin.HandlerFunc, error) {
	var config config.Config
	if err := container.Invoke(func(c t.Config) {
		config = c
	}); err != nil {
		return nil, err
	}

	var randomString t.RandomString
	if err := container.Invoke(func(r t.RandomString) {
		randomString = r
	}); err != nil {
		return nil, err
	}

	return func(context *gin.Context) {
		requestID := randomString(config.RequestID.Length, nil)
		context.Set("request_id", requestID)
	}, nil
}
