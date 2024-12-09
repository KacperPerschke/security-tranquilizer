package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wneessen/go-fileperm"
)

func FSCanCreate(fp string) error {
	dp := filepath.Dir(fp)
	dpi, err := os.Stat(dp)
	if err != nil {
		return err
	}
	if !dpi.IsDir() {
		return fmt.Errorf(
			"‘%s’ is not a directory",
			dp,
		)
	}
	p, err := fileperm.New(dp)
	if err != nil {
		return fmt.Errorf(
			"failed to read permitions of ‘%s’: %w",
			dp,
			err,
		)
	}
	if !p.UserWriteExecutable() {
		return fmt.Errorf(
			"cannot create a file ‘%s’",
			fp,
		)
	}
	return nil
}

func FSExistsAsFile(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf(
			"‘%s’ is a directory, not a file",
			path,
		)
	}
	return nil
}

const pathUP = ".."

func FSPrepElI(el string) (FileInfo, error) {
	eli, err := os.Stat(el)
	if err != nil {
		return FileInfo{}, err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return FileInfo{}, err
	}
	abs, err := filepath.Abs(el)
	if err != nil {
		return FileInfo{}, err
	}
	rel, err := filepath.Rel(cwd, abs)
	if err != nil {
		return FileInfo{}, err
	}
	relPathGoesUp := func() bool {
		return strings.Contains(rel, pathUP)
	}
	if relPathGoesUp() {
		return FileInfo{}, fmt.Errorf(
			"path ‘%s’ is outside of current working dir",
			el,
		)
	}
	prepEli := FileInfo{
		Info: eli,
		Path: rel,
	}
	return prepEli, nil
}
