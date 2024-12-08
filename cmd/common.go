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

func checkArgsCount(a []string) error {
	l := len(a)
	switch {
	case l == 0:
		return errTooFewArgs
	case l > 1:
		return errTooManyArgs
	default:
		return nil
	}
}

func getOutFileName(c *cobra.Command) (string, error) {
	outFileName, err := c.Flags().GetString("output")
	if err != nil {
		return stringZeroVal, err
	}
	return outFileName, nil
}
