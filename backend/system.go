/*
Copyright © 2022 BitsOfAByte

*/
package backend

import (
	"fmt"
	"os"
	"strings"
)

// Restricts running the binary as a root user. (Identified by UID == 0)
func IsRoot() bool {
	return os.Geteuid() == 0

}

// Detects if the binary is a preview version. (Identified by Version string)
func IsPreview() bool {
	return Version == "0.0.1-next" && os.Getenv("PROTO_CONSENT") != "true"
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
