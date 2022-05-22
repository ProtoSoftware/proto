/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"BitsOfAByte/proto/backend"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of installed versions of Proton.",
	Run: func(cmd *cobra.Command, args []string) {

		// Read the install directory
		installDir := backend.UsePath(viper.GetString("app.install_directory"), true)
		dir, err := ioutil.ReadDir(installDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Version", "Size", "Installed"})

		// Get the sizes of all the installed Proton versions.
		var totalSize int64
		for _, d := range dir {

			size, err := backend.DirSize(installDir + d.Name())

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			hSize, hUnit := backend.HumanReadableSize(size)
			totalSize += size

			table.Append([]string{d.Name(), fmt.Sprintf("%v%s", hSize, hUnit), d.ModTime().Format("2006-01-02")})
		}

		tSize, tUnit := backend.HumanReadableSize(totalSize)

		table.SetFooter([]string{"Total", fmt.Sprintf("%v%s", tSize, tUnit), " "})
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
