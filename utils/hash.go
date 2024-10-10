package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func generateSalt() string {
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    if err != nil {
        return ""
    }
    return hex.EncodeToString(salt)
}

func Hash(str string) string {
    h := sha256.New()

    h.Write([]byte(str + generateSalt()))
    hashValue := h.Sum(nil)
    hash := hex.EncodeToString(hashValue)
    return hash
}