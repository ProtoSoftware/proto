/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"BitsOfAByte/proto/backend"

	"github.com/spf13/cobra"
)

// appUpdateCmd represents the update command
var appUpdateCmd = &cobra.Command{
	Use:   "app-update",
	Short: "Update to the latest version of Proto",
	Run: func(cmd *cobra.Command, args []string) {
		// Call the updater from the backend.

		// TODO: Add progress bar, confirmation.
		backend.AppUpdate(backend.Version)
	},
}

func init() {
	rootCmd.AddCommand(appUpdateCmd)
}
