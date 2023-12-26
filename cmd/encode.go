package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Your friendly encoder",
	Long:  "Encodes the contents of the given file as described in the Readme.\nThe name of a file to be encoded is given as an argument to the call.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkArgsCount(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encode called")
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	// Reminder — here you can define your flags and configuration settings.
}
