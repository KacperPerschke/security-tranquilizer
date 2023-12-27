package img

import (
	"fmt"
	"image"
	"image/color"
)

func Encoder(stream []byte) (*image.Gray, error) {
	emptyImage := image.NewGray(image.Rect(0, 0, 0, 0))
	sLength := len(stream)

	size, err := resBySize(sLength)
	if err != nil {
		return emptyImage, fmt.Errorf("problem while guessing suitable image size: %w", err)
	}

	img := image.NewGray(
		image.Rect(
			0, 0, // xMin, yMin
			size.Width, size.Height, // xMax, yMax
		),
	)

	paddingVal := fillerContent(stream)
	calcPixColToSet := func(x, y int) (color.Gray, error) {
		sIdx, err := imgXYToSlicePos(x, y, img)
		if err != nil {
			return castByteToGray(0), err
		}
		if sIdx < sLength {
			return castByteToGray(stream[sIdx]), nil
		}
		return castByteToGray(paddingVal), nil
	}
	for x := 0; x <= size.Width; x++ {
		for y := 0; y <= size.Height; y++ {
			colToSet, err := calcPixColToSet(x, y)
			if err != nil {
				return emptyImage, err
			}
			img.SetGray(x, y, colToSet)
		}
	}
	return img, nil
}

func castByteToGray(b byte) color.Gray {
	return color.Gray{
		Y: b,
	}
}
