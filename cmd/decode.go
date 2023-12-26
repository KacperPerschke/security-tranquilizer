package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Your friendly decoder",
	Long:  "Decodes the contents of the given image as described in the Readme.\nThe name of a file to be decoded is given as an argument to the call.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkArgsCount(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decode called")
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	// Reminder — here you can define your flags and configuration settings.
}
