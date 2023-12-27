package img

import (
	"fmt"
	"image"
)

func encoder(stream []bytes) (image.Gray, error) {
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
	calcPixValToSet := func(x, y) (byte, error) {
		sIdx, err := imgXYToSlicePos(x, y, img)
		if err != nil {
			return intZeroVal, err
		}
		if sIdx < sLength {
			return stream[sIdx], nil
		}
		return paddingVal, nil
	}
	for x := 0; x <= size.Width; x++ {
		for y := 0; y <= size.Height; y++ {
			colToSet, err := calcPixValToSet(x, y)
			if err != nil {
				return emptyImage, err
			}
			img.SetColorIndex(x, y, colToSet)
		}
	}
	return img, nil
	/*
		f, _ := os.Create("image.png")
		defer f.Close()
		png.Encode(f, img)
	*/
}
