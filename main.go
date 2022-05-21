/*
Copyright © 2022 BitsOfAByte

*/
package main

import (
	"BitsOfAByte/proto/backend"
	"BitsOfAByte/proto/cmd"
	"fmt"
	"os"
)

func main() {

	backend.PreventRoot()

	if backend.IsPreview() {
		fmt.Println("Detected a preview version of Proto! If you are a user, please download a stable version from GitHub directly.")
		fmt.Println("If you are certain you want to run this preview build, set environment variable 'PROTO_PREVIEW_CONSENT' to 'true'")
		os.Exit(1)
	}

	cmd.Execute()
}
