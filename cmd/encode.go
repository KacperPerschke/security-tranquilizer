package cmd

import (
	"fmt"
	"image/png"
	"os"

	"github.com/spf13/cobra"

	"github.com/KacperPerschke/security-tranquilizer/img"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Your friendly encoder",
	Long:  "Encodes the contents of the given file as described in the Readme.\nThe name of a file to be encoded is given as an argument to the call.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkArgsCount(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := encodeFileToPNG(cmd, args)
		if err != nil {
			fmt.Printf("\nThere was an error while encoding: %q\n\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	addOutputFlag(encodeCmd)
	// Reminder — here you can define your flags and configuration settings.
}

func encodeFileToPNG(c *cobra.Command, args []string) error {
	iFName := args[0]
	oFName, err := getOutFileName(c)
	if err != nil {
		return fmt.Errorf("Problem during attempt to get value of `output` flag: %w", err)
	}

	b, err := os.ReadFile(iFName)
	if err != nil {
		return fmt.Errorf("Problem during attempt to read file '%s': %w", iFName, err)
	}

	img, err := img.PackToImg(b)
	if err != nil {
		return err
	}

	oHandle, err := os.OpenFile(oFName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("Problem during attempt to open file '%s' for writing: %w", oFName, err)
	}

	errPNG := png.Encode(oHandle, img)
	if errPNG != nil {
		return fmt.Errorf("Problem during attempt to write image to file '%s': %w", oFName, errPNG)
	}

	errClose := oHandle.Close()
	if err != nil {
		return fmt.Errorf("Problem during releasing handle to file '%s' after write: %w", oFName, errClose)
	}

	return nil
}
