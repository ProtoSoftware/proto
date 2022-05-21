/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"BitsOfAByte/proto/backend"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of installed versions of Proton.",
	Run: func(cmd *cobra.Command, args []string) {
		installDir := backend.UsePath(viper.GetString("app.install_directory"), "dir")

		// for every directory inside of the install directory,
		// print the name of the directory.
		dir, err := ioutil.ReadDir(installDir)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, d := range dir {

			size, err := backend.DirSize(installDir + d.Name())

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			hSize, hUnit := backend.HumanReadableSize(size)

			fmt.Printf("Version: %s | Size: %v%s | Modified: %s\n", d.Name(), hSize, hUnit, d.ModTime().Format("2006-01-02 @ 15:04"))

		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
