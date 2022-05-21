/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"BitsOfAByte/proto/backend"
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
		releases, err := backend.GetReleases()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Tag", "Released On", "Install Command"})

		limit, _ := cmd.Flags().GetInt("limit")

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

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(releasesCmd)
	releasesCmd.Flags().IntP("limit", "l", 5, "Limit the number of releases to show.")
}
