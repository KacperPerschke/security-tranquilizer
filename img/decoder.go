package img

import (
	"errors"
	"fmt"
	"image"
	"image/color"
)

func UnpackFromImg(i *image.Gray) ([]byte, error) {
	emptyStream := make([]byte, 0, 0)

	xLast, yLast, err := calcLastPos(i)
	if err != nil {
		return emptyStream, fmt.Errorf("problem during calclulating padding: %w", err)
	}
	// fmt.Printf("xLast=%4d, yLast=%4d\n", xLast, yLast)
	haveReachedLastPix := func(x, y int) bool {
		return x == xLast && y == yLast
	}

	desCap := imgXYToSlicePos(xLast, yLast, i)
	// fmt.Printf("dec: desCap=%3d\n", desCap)
	stream := make(
		[]byte,
		desCap, // initial len
		desCap, //capacity
	)

	b := i.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			val := grayToByte(i.At(x, y))
			sIdx := imgXYToSlicePos(x, y, i)
			// fmt.Printf("dec: x=%3d, y=%3d, sIdx=%3d\n", x, y, sIdx)
			if haveReachedLastPix(x, y) {
				return stream, nil
			}
			if sIdx < desCap {
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
	fillValue := grayToByte(i.At(xMax, yMax))
	// fmt.Printf("fillValue=%4d, %T\n", fillValue, fillValue)
	if fillValue != paddingEOT && fillValue != paddingETX {
		err := fmt.Errorf("the last pixel is wild: %X", fillValue)
		return intZeroVal, intZeroVal, err
	}
	for y := yMax; y >= yMin; y-- {
		for x := xMax; x >= xMin; x-- {
			colorAt := grayToByte(i.At(x, y))
			// fmt.Printf("cLP: x=%5d, y=%5d, c=%3d, %T\n", x, y, colorAt, colorAt)
			if colorAt != fillValue {
				return x + 1, y, nil
			}
		}
	}

	return intZeroVal, intZeroVal, errors.New("I didn't find the end")
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
