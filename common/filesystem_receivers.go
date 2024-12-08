package common

import "io/fs"

func (el FileInfo) IsFileOrSymlink() bool {
	mode := el.Info.Mode()
	isSymlink := func() bool {
		return mode&fs.ModeSymlink == 0
	}
	if mode.IsRegular() || isSymlink() {
		return true
	}
	return false
}
