package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Hash(data string) string {
	hash := sha256.New()
	io.WriteString(hash, data)
	hashedBytes := hash.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}
