package utils

import (
	"fmt"
	"strings"
    "crypto/sha256"
    "encoding/hex"
)

func Hash(str string) (string) {
    parts := strings.Split(str, ".")
    strWithoutExt := parts[0]

    h := sha256.New()
    h.Write([]byte(strWithoutExt))
    hashValue := h.Sum(nil)
    hash := hex.EncodeToString(hashValue)
    return fmt.Sprintf("%s.%s", hash, parts[1])
}