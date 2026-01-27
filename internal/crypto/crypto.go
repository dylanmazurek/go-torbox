package crypto

import (
	"crypto/sha1"
	"fmt"
)

// ToSHA1 computes the SHA-1 hash of the given byte slice and returns it as a hexadecimal string.
// p is the input byte slice to be hashed.
// b is an optional byte slice to which the hash will be appended.
func ToSHA1(p []byte, b []byte) string {
	hash := sha1.New()
	hash.Write(p)

	newHashHexStr := fmt.Sprintf("%x", hash.Sum(b))

	return newHashHexStr
}
