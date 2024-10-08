package utils

import (
	"os"
	"os/exec"
	"log"
)

func Log(message string) {
	log.Printf("- %s\n", message)
}

func LogFatal(message string) {
	log.Fatalf("- Fatal: %s\n", message)
}

func LogError(message string) {
	log.Printf("- Error: %s\n", message)
}

func ClearTerminal() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}