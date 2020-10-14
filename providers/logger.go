package providers

import (
	"github.com/yz89120/geo-ip/providers/t"
	"go.uber.org/zap"
)

func logger(c t.Config) (t.Logger, error) {
	if c.Debug {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}
