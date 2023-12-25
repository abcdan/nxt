package helper

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
                "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
                "0123456789-_"

func GenerateShortCode(length int) string {
    seed := rand.NewSource(time.Now().UnixNano())
    generator := rand.New(seed)

    b := make([]byte, length)
    for i := range b {
        b[i] = charset[generator.Intn(len(charset))]
    }
    return string(b)
}
