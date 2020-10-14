package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	// error is ignored
	godotenv.Load()

}

// Config of the app
var config = Config{
	Debug: false,
	GeoIP: geoIP{
		DBPath: "",
	},
	RequestID: requestID{
		Length: 11,
	},
}

// GetConfig return the config
func GetConfig() Config {
	if value, ok := os.LookupEnv("APP_DEBUG"); ok {
		value = strings.Trim(value, " \r\n")
		value = strings.ToLower(value)
		if value == "true" {
			config.Debug = true
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
	}

	if value, ok := os.LookupEnv("GEO_IP_DB_PATH"); ok {
		value = strings.Trim(value, " \r\n")
		config.GeoIP.DBPath = value
	} else {
		// log.Fatalln("GEO_IP_DB_PATH is not set")
	}

	if value, ok := os.LookupEnv("REQUEST_ID_LENGTH"); ok {
		if length, err := strconv.Atoi(value); err == nil {
			config.RequestID.Length = length
		}
	}

	return config
}

// Config type
// The configuration of the app
type Config struct {
	Debug     bool
	GeoIP     geoIP
	RequestID requestID
}

type requestID struct {
	Length int
}

type geoIP struct {
	DBPath string
}
