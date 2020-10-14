package providers

import (
	"github.com/yz89120/geo-ip/providers/t"
)

func randomStringCharset() t.RandomStringCharset {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"
	return &charset
}
