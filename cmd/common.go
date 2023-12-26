package cmd

import "errors"

var (
	errTooFewArgs  = errors.New("You haven't provided a single argument.")
	errTooManyArgs = errors.New("You haven provided too many arguments.")
)

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
