/*
Copyright © 2022 BitsOfAByte

*/
package shared

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

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
