package crypto

import (
	"crypto/sha1"
	"fmt"
)

func ToSHA1(data []byte) string {
	hash := sha1.New()
	hash.Write(data)

	return fmt.Sprintf("%x", hash.Sum(nil))
}
