package providers

import (
	"github.com/oschwald/geoip2-golang"
	"github.com/yz89120/geo-ip/providers/t"
)

func geoIPDB(c t.Config) (t.GeoIPDB, error) {
	return geoip2.Open(c.GeoIP.DBPath)
}
