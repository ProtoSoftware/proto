package backend

import (
	"fmt"

	"github.com/spf13/viper"
)

// Log a message if the logger is enabled.
func Debug(msg string) {
	if viper.GetBool("app.verbose") {
		fmt.Printf("[DEBUG] %s\n", msg)
	}
}
