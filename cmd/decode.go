package cmd

import (
	"fmt"
	"os"

	"github.com/KacperPerschke/security-tranquilizer/common"
	"github.com/KacperPerschke/security-tranquilizer/img"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode [flags] IMAGE_USED_AS_TRANSPORT",
	Short: "Your friendly decoder",
	Long:  "Decodes the contents of the given image as described in the Readme.\nThe name of a file to be decoded is given as an argument to the call.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := checkArgsCountMin(args); err != nil {
			return err
		}
		if err := checkArgsCountMax(args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		iFName := args[0]
		if err := common.FSExistsAsFile(iFName); err != nil {
			fmt.Printf("\nThere was an error with input file: %q\n\n", err)
			os.Exit(1)
		}
		if err := img.DecodeFromPNG(iFName); err != nil {
			fmt.Printf("\nThere was an error while decoding: %q\n\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	// Reminder — here you can define your flags and configuration settings.
}
