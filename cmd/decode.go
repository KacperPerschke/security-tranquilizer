package cmd

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/KacperPerschke/security-tranquilizer/img"
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
		err := decodeFileFromPNG(cmd, args)
		if err != nil {
			fmt.Printf("\nThere was an error while encoding: %q\n\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	addOutputFlag(decodeCmd)
	// Reminder — here you can define your flags and configuration settings.
}

func decodeFileFromPNG(c *cobra.Command, args []string) error {
	iFName := args[0]
	oFName, err := getOutFileName(c)
	if err != nil {
		return fmt.Errorf("Problem during attempt to get value of `output` flag: %w", err)
	}

	imgRead, err := img.ReadFromPNG(iFName)
	if err != nil {
		return err
	}
	imgGray := image.NewGray(imgRead.Bounds())
	draw.Draw(imgGray, imgGray.Bounds(), imgRead, imgRead.Bounds().Min, draw.Src)
	bOut, err := img.UnpackFromImg(imgGray)
	if err != nil {
		return err
	}

	errWrite := os.WriteFile(oFName, bOut, 0644)
	if errWrite != nil {
		return fmt.Errorf("Problem during attempt to write image to file '%s': %w", oFName, errWrite)
	}

	return nil
}
