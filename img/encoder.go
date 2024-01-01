package img

import (
	"fmt"
	"image"
	"image/color"
)

func PackToImg(stream []byte) (*image.Gray, error) {
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
