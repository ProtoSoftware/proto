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
			// The directory doesnt exist, meaning there are no installed versions.
			if os.IsNotExist(err) {
				fmt.Println("No installed versions of proton found at " + installDir)
				os.Exit(0)
			}

			// Something else went wrong, eg. permissions.
			fmt.Println(err)
			os.Exit(1)
		}

		// Get all of the installed versions and their sizes and create a table to display them.
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Version", "Size", "Installed"})
		var totalSize int64
		for _, d := range dir {
			size, err := backend.DirSize(installDir + d.Name())

			// Something went wrong getting the size of the directory.
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Get the size of the directory, and add to the total size, then append to the table.
			hSize, hUnit := backend.HumanReadableSize(size)
			totalSize += size
			table.Append([]string{d.Name(), fmt.Sprintf("%v%s", hSize, hUnit), d.ModTime().Format("2006-01-02")})
		}

		// No installed versions found in the install directory.
		if table.NumLines() == 0 {
			fmt.Println("No installed versions of proton found at " + installDir)
			os.Exit(0)
		}

		// Format the total size and render the table.
		tSize, tUnit := backend.HumanReadableSize(totalSize)
		table.SetFooter([]string{"Total", fmt.Sprintf("%v%s", tSize, tUnit), " "})
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
