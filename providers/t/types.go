package t

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/yz89120/geo-ip/config"
	"go.uber.org/zap"
)

// Config type
type Config = config.Config

// GeoIPDB type
type GeoIPDB = *geoip2.Reader

// Logger type
type Logger = *zap.Logger

// Server type
type Server = *gin.Engine

// Rand type
type Rand = *rand.Rand

// RandomStringCharset type
type RandomStringCharset *string

// RandomString type
type RandomString func(length int, charset *string) string
