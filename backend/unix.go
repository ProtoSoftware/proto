package backend

import (
	"fmt"
	"os"
	"strings"
)

// Restricts running the binary as a root user. (Identified by UID == 0)
func PreventRoot() {
	if os.Geteuid() == 0 {
		fmt.Println("Proto cannot not be run as root, please try again as a regular user without sudo.")
		os.Exit(1)
	}
}

// Prompt a user to confirm something
func Prompt(message string, defaultValue bool) bool {
	var response string

	Debug("Prompt: Asking for user input")

	fmt.Print(message)
	fmt.Scanln(&response)

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		return defaultValue
	}
}
