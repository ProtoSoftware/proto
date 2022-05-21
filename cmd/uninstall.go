/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"BitsOfAByte/proto/backend"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:     "uninstall <version>",
	Short:   "Uninstall a version of Proton from your system.",
	Example: "proto uninstall ProtonGE7-18",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		installDir := backend.UsePath(viper.GetString("app.install_directory"), "dir") + args[0]

		// If the directory doesn't exist, we can't uninstall.
		if _, err := os.Stat(installDir); os.IsNotExist(err) {
			fmt.Println("The specified version of Proton was not found.")
			os.Exit(1)
		}

		err := os.RemoveAll(installDir)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Uninstalled:" + args[0])
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
