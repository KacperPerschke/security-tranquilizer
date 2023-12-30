package img

import (
	"errors"
	"fmt"
	"image"
	"os"
)

func ReadFromPNG(fName string) (image.Image, error) {
	fHandle, err := os.Open(fName)
	if err != nil {
		errRet := fmt.Errorf(
			"Problem during attempt to open file '%s' for reading: %w",
			fName, err,
		)
		return emptyImage, errRet
	}
	img, formatName, err := image.Decode(fHandle)
	if err != nil {
		return emptyImage, err
	}
	if formatName != `png` {
		err := errors.New("file has another format than 'png'")
		return emptyImage, err
	}
	return img, nil
}

func WriteToPNG(i image.Image, fName string) error {
	return nil
}
