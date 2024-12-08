package common

import "os"

type FileInfo struct {
	Info os.FileInfo
	Path string
}
