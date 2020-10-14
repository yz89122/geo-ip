package ipaddresslookup

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/yz89120/geo-ip/http/api/response"
	"github.com/yz89120/geo-ip/providers/t"
	"go.uber.org/dig"
)

// Query for retriving input
type Query struct {
	IPAddress string `uri:"ip_addr" bind:"required"`
}

type apiResponse struct {
	*geoip2.City
}

// IPAddressLookup respond the info of the given IP Address
func IPAddressLookup(container *dig.Container) (gin.HandlerFunc, error) {
	var db *geoip2.Reader
	if err := container.Invoke(func(d t.GeoIPDB) {
		db = d
	}); err != nil {
		return nil, err
	}

	return func(c *gin.Context) {
		r := response.New()
		defer r.Respond(c)

		query := Query{}
		c.BindUri(&query)

		if city, err := db.City(net.ParseIP(query.IPAddress)); err != nil {
			r.AppendError(response.NotFoundError)
		} else {
			r.SetResult(&apiResponse{
				City: city,
			})
		}
	}, nil
}
