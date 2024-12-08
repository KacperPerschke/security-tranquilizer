package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "security-tranquilizer",
		Short: "Your friend",
		Long:  `Plase read the Readme ( https://github.com/KacperPerschke/security-tranquilizer/blob/main/README.md )`,
		// Run: Intentioanlly empty
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Reminder — here you can define your flags and configuration settings.
}
