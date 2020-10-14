package providers

import (
	c "github.com/yz89120/geo-ip/config"
	"github.com/yz89120/geo-ip/providers/t"
)

func config() t.Config {
	return c.GetConfig()
}
