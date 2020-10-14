package main

import (
	"log"

	"github.com/yz89120/geo-ip/container"
	"github.com/yz89120/geo-ip/providers"
	"github.com/yz89120/geo-ip/providers/t"
)

func main() {
	container := container.Get()
	providers.Provide(container)

	if err := container.Invoke(func(server t.Server) {
		server.Run(":8080")
	}); err != nil {
		log.Fatalln(err)
	}
}
