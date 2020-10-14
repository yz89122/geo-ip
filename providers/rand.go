package providers

import (
	r "math/rand"
	"time"

	"github.com/yz89120/geo-ip/providers/t"
)

func rand() t.Rand {
	return r.New(r.NewSource(time.Now().UnixNano()))
}
