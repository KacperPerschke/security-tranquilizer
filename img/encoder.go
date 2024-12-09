package img

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/KacperPerschke/security-tranquilizer/archiver"
	"github.com/KacperPerschke/security-tranquilizer/common"
)

func EncodeToPNG(of string, il []common.FileInfo) error {
	b, err := archiver.PackToArchive(il)
	if err != nil {
		return err // Should we include more information?
	}
	wErr := os.WriteFile("check.tar.gz", b, 0644)
	if wErr != nil {
		return err
	}

	img, err := bytesToImg(b)
	if err != nil {
		return err
	}

	oHandle, err := os.OpenFile(string(of), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("Problem during attempt to open file '%s' for writing: %w", of, err)
	}

	errPNG := png.Encode(oHandle, img)
	if errPNG != nil {
		return fmt.Errorf("Problem during attempt to write image to file '%s': %w", of, errPNG)
	}

	errClose := oHandle.Close()
	if err != nil {
		return fmt.Errorf("Problem during releasing handle to file '%s' after write: %w", of, errClose)
	}

	return nil
}

func bytesToImg(stream []byte) (*image.Gray, error) {
	sLen := len(stream)

	rect, err := resBySize(sLen)
	if err != nil {
		return emptyImage, fmt.Errorf("problem while guessing suitable image size: %w", err)
	}
	img := image.NewGray(rect)

	paddingVal := fillerContent(stream)
	calcPixColToSet := func(x, y int) color.Gray {
		sIdx := imgXYToSlicePos(x, y, img)
		if sIdx < sLen {
			return newGray(stream[sIdx])
		}
		return newGray(paddingVal)
	}
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			colToSet := calcPixColToSet(x, y)
			img.SetGray(x, y, colToSet)
		}
	}
	return img, nil
}

func newGray(b byte) color.Gray {
	return color.Gray{
		Y: b,
	}
}
