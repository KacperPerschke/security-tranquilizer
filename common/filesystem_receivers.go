package common

func (el FileInfo) AppearsToBeFile() bool {
	if el.Info.Mode().IsDir() {
		return false
	}
	return true
}
