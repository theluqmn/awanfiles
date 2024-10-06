package utils

import (
	"log"
)

func Log(message string) {
	log.Printf("%s\n", message)
}

func LogFatal(message string) {
	log.Fatalf("Fatal: %s\n", message)
}