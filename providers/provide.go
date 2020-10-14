package providers

import "go.uber.org/dig"

// Provide provides services for container
func Provide(container *dig.Container) {
	container.Provide(config)
	container.Provide(geoIPDB)
	container.Provide(logger)
	container.Provide(rand)
	container.Provide(randomString)
	container.Provide(randomStringCharset)
	container.Provide(server)
}
