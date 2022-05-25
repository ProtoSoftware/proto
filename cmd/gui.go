/*
Copyright © 2022 BitsOfAByte

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// guiCmd represents the gui command
var guiCmd = &cobra.Command{
	Use:     "gui",
	Short:   "Opens the Proto GUI (Coming Soon)",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The GUI has not been implemented yet! For now you'll have to use the command line.")
		fmt.Println("For help with the command line, use the 'proto help' command.")
		fmt.Println("To exit this menu, press ENTER.")
		fmt.Scanln()
	},
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
