package img

import (
	"fmt"
	"image"
	"image/color"
)

func PackToImg(stream []byte) (*image.Gray, error) {
	sLength := len(stream)

	rect, err := resBySize(sLength)
	if err != nil {
		return emptyImage, fmt.Errorf("problem while guessing suitable image size: %w", err)
	}
	img := image.NewGray(rect)

	// fmt.Printf("Stride=%4d\n", img.Stride)
	// fmt.Printf("BB=%#q\n", img.Bounds())

	paddingVal := fillerContent(stream)
	calcPixColToSet := func(x, y int) color.Gray {
		sIdx := imgXYToSlicePos(x, y, img)
		if sIdx < sLength {
			return newGray(stream[sIdx])
		}
		return newGray(paddingVal)
	}
	b := img.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			colToSet := calcPixColToSet(x, y)
			// fmt.Printf("enc: x=%4d, y=%4d, c=%3d i.e. %c\n", x, y, colToSet.Y, colToSet.Y)
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
