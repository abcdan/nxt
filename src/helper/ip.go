package helper

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashIP(ip string) string {
	hasher := sha512.New()
	hasher.Write([]byte(ip))
	return hex.EncodeToString(hasher.Sum(nil))
}
