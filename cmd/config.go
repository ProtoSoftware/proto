/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"BitsOfAByte/proto/shared"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Change the configuration of Proto",
}

var showConfCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		formatOut, err := json.MarshalIndent(viper.AllSettings(), "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(formatOut))
		fmt.Println("Located at: " + viper.ConfigFileUsed())
	},
}

// verboseCmd represents the verbose command
var verboseCmd = &cobra.Command{
	Use:       "verbose <bool>",
	Short:     "Toggle verbose mode",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"true", "false"},
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("app.verbose", args[0])
		viper.WriteConfig()
	},
}

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:     "source <owner/repo>",
	Short:   "Change the source of Proton downloads",
	Example: "proto config source GloriousEggroll/proton-ge-custom",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("app.proton_source", args[0])
		viper.WriteConfig()
	},
}

// tempCmd represents the temp command
var tempCmd = &cobra.Command{
	Use:     "temp <dir>",
	Short:   "Change the temporary storage location",
	Example: "proto config temp /tmp/proto/",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Temp location is now: ", shared.UsePath(args[0], true))
		viper.Set("app.temp_storage", shared.UsePath(args[0], true))
		viper.WriteConfig()
	},
}

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:       "force-sum <bool>",
	Short:     "Enable or disable mandatory checksum passing for all downloads",
	Example:   "proto config force-sum true",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"true", "false"},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Force sum has been set to", args[0])
		viper.Set("app.force_sum", args[0])
		viper.WriteConfig()
	},
}

// installDirCmd represents the install command
var installDirCmd = &cobra.Command{
	Use:     "install <dir>",
	Short:   "Change the proton install directory",
	Example: "proto config install ~/.steam/root/compatibilitytools.d/",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Install location is now: ", shared.UsePath(args[0], true))
		viper.Set("app.install_directory", shared.UsePath(args[0], true))
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(sourceCmd)
	configCmd.AddCommand(tempCmd)
	configCmd.AddCommand(installDirCmd)
	configCmd.AddCommand(checksumCmd)
	configCmd.AddCommand(verboseCmd)
	configCmd.AddCommand(showConfCmd)
}
