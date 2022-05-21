/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"BitsOfAByte/proto/backend"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "proto",
	Short: "Install and manage Proton-GE installations ",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// Iniotialized protos configuration file
func initConfig() {
	// Set the configuration file location
	configDir, _ := os.UserConfigDir()
	viper.SetConfigName("config")
	viper.AddConfigPath(configDir + "/proto")
	viper.SetConfigType("json")

	// Set the app default settings
	viper.SetDefault("app.temp_storage", backend.UsePath("/tmp/proto/", "dir"))
	viper.SetDefault("app.install_directory", backend.UsePath("~/.steam/root/compatibilitytools.d/", "dir"))
	viper.SetDefault("app.proton_source", "GloriousEggroll/proton-ge-custom")
	viper.SetDefault("app.force_sum", "true")

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
