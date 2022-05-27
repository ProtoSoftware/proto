/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"BitsOfAByte/proto/shared"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "proto",
	Short:   "Install and manage Proton-GE installations ",
	Version: shared.Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Register persistent flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Skip all confirmation prompts")

	// Register flags to config
	viper.BindPFlag("app.verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// Iniotialized protos configuration file
func initConfig() {
	// Set the configuration file location
	configDir, _ := os.UserConfigDir()
	viper.SetConfigName("config")
	viper.AddConfigPath(configDir + "/proto")
	viper.SetConfigType("json")

	// Set the app default settings
	viper.SetDefault("app.temp_storage", shared.UsePath("/tmp/proto/", true))
	viper.SetDefault("app.install_directory", shared.UsePath("~/.steam/root/compatibilitytools.d/", true))
	viper.SetDefault("app.proton_source", "GloriousEggroll/proton-ge-custom")
	viper.SetDefault("app.force_sum", "true")
	viper.SetDefault("app.verbose", "false")

	// Write a configuration file if it doesnt exist, or throw an error if something goes wrong
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			os.MkdirAll(configDir+"/proto", os.ModePerm)
			viper.SafeWriteConfig()
		} else {
			os.Exit(1)
		}
	}
}
