package ping

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yz89120/geo-ip/http/api/response"
	"go.uber.org/dig"
)

// Ping pong
func Ping(_ *dig.Container) (gin.HandlerFunc, error) {
	return func(c *gin.Context) {
		response := response.New()
		defer response.Respond(c)

		response.SetResult(time.Now().Format(time.RFC3339Nano))
	}, nil
}
