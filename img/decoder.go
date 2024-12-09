package img

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/KacperPerschke/security-tranquilizer/archiver"
)

func DecodeFromPNG(path string) error {
	imgRead, err := ReadFromPNG(path)
	if err != nil {
		return err
	}
	imgGray := image.NewGray(imgRead.Bounds())
	draw.Draw(imgGray, imgGray.Bounds(), imgRead, imgRead.Bounds().Min, draw.Src)
	bOut, err := bytesFromImg(imgGray)
	if err != nil {
		return err
	}

	if err := archiver.UnpackFrom(bytes.NewReader(bOut)); err != nil {
		return err
	}
	return nil
}

func bytesFromImg(i *image.Gray) ([]byte, error) {
	emptyStream := make([]byte, 0, 0)

	xLast, yLast, err := calcLastPos(i)
	if err != nil {
		return emptyStream, fmt.Errorf("problem during calclulating padding: %w", err)
	}
	haveReachedLastPix := func(x, y int) bool {
		return x == xLast && y == yLast
	}

	lenExpected := imgXYToSlicePos(xLast, yLast, i)
	stream := make(
		[]byte,
		lenExpected, // It allows inserting by index rather than using 'append'.
		lenExpected, // Let's allocate the exact number of memory cells needed.
	)

	b := i.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			val := grayToByte(i.At(x, y))
			sIdx := imgXYToSlicePos(x, y, i)
			if haveReachedLastPix(x, y) {
				return stream, nil
			}
			if sIdx < lenExpected {
				stream[sIdx] = val
			}
		}
	}

	return emptyStream, nil
}

func calcLastPos(i *image.Gray) (int, int, error) {
	b := i.Bounds()
	xMin := b.Min.X
	xMax := b.Max.X - 1
	yMin := b.Min.Y
	yMax := b.Max.Y - 1

	xPrev, yPrev, fillValue := xMax, yMax, grayToByte(i.At(xMax, yMax))
	if fillValue != paddingEOT && fillValue != paddingETX {
		err := fmt.Errorf("the last pixel is wild: %X", fillValue)
		return intZeroVal, intZeroVal, err
	}

	for y := yMax; y >= yMin; y-- {
		for x := xMax; x >= xMin; x-- {
			colorAt := grayToByte(i.At(x, y))
			if colorAt != fillValue {
				return xPrev, yPrev, nil
			}
			xPrev, yPrev = x, y
		}
	}

	return intZeroVal, intZeroVal, errors.New("I didn't find the end") // Should not happen.
}

func calcStreamLen(i image.Image, xLast, yLast int) int {
	numOfFullRows := yLast
	numOfPixsInFullRow := i.Bounds().Max.X
	numOfPixsInNotFullRow := xLast
	return numOfFullRows*numOfPixsInFullRow + numOfPixsInNotFullRow
}

func grayToByte(c color.Color) byte {
	cg, isGray := c.(color.Gray)
	if !isGray {
		err := fmt.Sprintf("Unexpected color model '%T'!", c)
		panic(err)
	}
	return cg.Y
}
