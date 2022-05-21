/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"BitsOfAByte/proto/backend"
	"fmt"
	"os"

	"github.com/google/go-github/v44/github"
	cobra "github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [tag]",
	Short: "Download and install Proton to your system.",
	Long: `Download and install Proton to your system.
Run without arguments to install to the latest version or specify a tag to install.`,
	Run: func(cmd *cobra.Command, args []string) {

		/**
		----------------------
		|     Fetch Logic    |
		----------------------
		**/

		var tagData *github.RepositoryRelease

		switch len(args) {
		case 0: // Install latest tag.
			data, err := backend.GetReleases()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			tagData = data[0]
		default: // Install a specific tag.
			data, err := backend.GetReleaseData(args[0])

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			tagData = data
		}

		/**
		----------------------
		|   Download Logic   |
		----------------------
		**/

		installDir := backend.UsePath(viper.GetString("app.install_directory"), "dir")

		// Check if folder exists
		if _, err := os.Stat(installDir + tagData.GetName()); os.IsNotExist(err) {
			if err == nil {
				fmt.Printf("Looks like you already have %s installed.\n", tagData.GetName())
				os.Exit(1)
			}
		}

		// Fetch valid assets from the release.
		tar, sum, err := backend.GetValidAssets(tagData)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Handle the lack of a checksum depending on the user's preference.
		if sum == nil {
			forced := viper.GetBool("app.force_sum")
			if forced {
				fmt.Println("No checksum file found, aborting install.")
				os.Exit(1)
			}
			fmt.Println("No checksum file found, continuing without verification.")
		}

		// Download the assets to the temp directory.
		tmp := backend.UsePath(viper.GetString("app.temp_storage"), "file")

		// If it exists, download the checksum file.
		if sum != nil {
			_, err = backend.DownloadFile(tmp+sum.GetName(), sum.GetBrowserDownloadURL())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// Download the tarball.
		_, err = backend.DownloadFile(tmp+tar.GetName(), tar.GetBrowserDownloadURL())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		/**
		----------------------
		|   Checksum Logic   |
		----------------------
		**/

		// TODO: Verify checksums here.
		fmt.Println("Checksum verification not implemented, skipping regardless of setting.")

		/**
		----------------------
		|   Install Logic    |
		----------------------
		**/

		tarReader, err := os.Open(tmp + tar.GetName())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = backend.ExtractTar(installDir, tarReader)

		defer tarReader.Close()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		/**
		----------------------
		| Post-Install Logic |
		----------------------
		**/

		err = backend.ClearTemp()

		if err != nil {
			fmt.Println("Failed to perform cleanup on temp directory, please remove manually.")
		}

		fmt.Printf("%s has been successfully installed to %s\n", tagData.GetTagName(), viper.GetString("app.install_directory"))
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
