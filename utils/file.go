package utils

import (
    "path/filepath"
    "strings"
)

func GetFileName(fileName string) string {
    base := filepath.Base(fileName)
    return strings.TrimSuffix(base, filepath.Ext(base))
}

func GetFileFormat(fileName string) string {
	return filepath.Ext(fileName)
}