/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"ProtoSoftware/proto/shared"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// releasesCmd represents the releases command
var releasesCmd = &cobra.Command{
	Use:     "releases",
	Short:   "Show all available releases from the proton source.",
	Example: `proto releases --limit 5`,
	Run: func(cmd *cobra.Command, args []string) {

		// If there are multiple sources, ask the user which one to use.
		source := shared.GetSourceIndex()

		// Get the releases from the backend.
		releases, err := shared.GetReleases(source)
		shared.Check(err)

		// Create a table to display the releases.
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Tag", "Released On", "Install Command"})
		limit, _ := cmd.Flags().GetInt("limit")

		// Loop through the releases and add them to the table up to the limit.
		for _, release := range releases {

			if limit > 0 {
				limit--
			} else {
				break
			}

			table.Append([]string{
				release.GetTagName(),
				release.GetPublishedAt().Format("2006-01-02"),
				fmt.Sprintf("proto install %s", release.GetTagName()),
			})
		}

		// Display the table.
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(releasesCmd)

	// Register command flags
	releasesCmd.Flags().IntP("limit", "l", 5, "Limit the number of releases to show.")
}
