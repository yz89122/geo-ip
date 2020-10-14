package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/yz89120/geo-ip/providers/t"
	"github.com/yz89120/geo-ip/routes"
	"go.uber.org/dig"
)

func server(container *dig.Container) t.Server {
	engine := gin.New()
	routes.SetUpRoutes(container, engine)
	return engine
}
