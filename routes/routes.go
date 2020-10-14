package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yz89120/geo-ip/handlers"
	"github.com/yz89120/geo-ip/handlers/ipaddresslookup"
	"github.com/yz89120/geo-ip/handlers/ping"
	"github.com/yz89120/geo-ip/middlewares"
	"go.uber.org/dig"
)

var globalMiddlewares []middlewares.MiddlewareFactoryContract = []middlewares.MiddlewareFactoryContract{
	middlewares.Recovery,
	middlewares.RequestID,
	middlewares.Logger,
	middlewares.LogRequest,
	middlewares.RequestIDHeader,
}

type route struct {
	method         string
	path           string
	handlerFactory handlers.HandlerFactoryContract
}

var rootRoutes []route = []route{
	route{
		method:         "any",
		path:           "ping",
		handlerFactory: ping.Ping,
	},
	route{
		method:         "get",
		path:           "ip_addr/:ip_addr",
		handlerFactory: ipaddresslookup.IPAddressLookup,
	},
}

// SetUpRoutes set up routes
func SetUpRoutes(container *dig.Container, e *gin.Engine) error {
	for _, middleware := range globalMiddlewares {
		if handler, err := middleware(container); err == nil {
			e.Use(handler)
		} else {
			return err
		}
	}

	for _, route := range rootRoutes {
		if handler, err := route.handlerFactory(container); err == nil {
			switch route.method {
			case "get":
				e.GET(route.path, handler)
			case "any":
				e.Any(route.path, handler)
			}
		} else {
			return err
		}
	}

	return nil
}
