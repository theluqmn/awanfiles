package utils

import (
    "crypto/sha256"
    "encoding/hex"
)

func Hash(str string) (hash string) {
    h := sha256.New()
    h.Write([]byte(str))
    hashValue := h.Sum(nil)
	hash = hex.EncodeToString(hashValue)
	return hash
}