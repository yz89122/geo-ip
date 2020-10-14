package providers

import (
	"github.com/yz89120/geo-ip/providers/t"
)

func randomString(seededRand t.Rand, defaultCharset t.RandomStringCharset) t.RandomString {
	return func(length int, charset *string) string {
		if length <= 0 {
			return ""
		}

		if charset == nil {
			charset = defaultCharset
		}

		b := make([]byte, length)

		for i := range b {
			b[i] = (*charset)[seededRand.Intn(len(*charset))]
		}

		return string(b)
	}
}
