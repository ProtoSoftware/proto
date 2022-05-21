package backend

import (
	"fmt"
	"os"
)

func PreventRoot() {
	if os.Geteuid() == 0 {
		fmt.Println("Proto cannot not be run as root, please try again as a regular user without sudo.")
		os.Exit(1)
	}
}
