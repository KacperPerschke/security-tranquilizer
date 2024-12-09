package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	stringZeroVal = ``
)

var (
	errTooFewArgs  = errors.New("You have provided too few arguments.")
	errTooManyArgs = errors.New("You have provided too many arguments.")
)

func addOutputFlag(c *cobra.Command) {
	c.Flags().StringVar(&outFileName, "output", "", "name of output file (there is no default value; you must supply sth)")
	c.MarkFlagRequired("output")
}

func checkArgsCountMax(a []string) error {
	if len(a) > 1 {
		return errTooManyArgs
	}
	return nil
}

func checkArgsCountMin(a []string) error {
	if len(a) == 0 {
		return errTooFewArgs
	}
	return nil
}

func getOutFileName(c *cobra.Command) (string, error) {
	outFileName, err := c.Flags().GetString("output")
	if err != nil {
		return stringZeroVal, err
	}
	return outFileName, nil
}
