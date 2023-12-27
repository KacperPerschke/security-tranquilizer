package img

import (
	"fmt"
	"image"
)

const (
	intZeroVal = 0
	paddingEOT = 4 // https://en.wikipedia.org/wiki/End-of-Transmission_character
	paddingETX = 3 // https://en.wikipedia.org/wiki/End-of-Text_character
)

func imgXYToSlicePos(x, y int, img *image.Gray) (int, error) {
	bBox := img.Rect
	fmtBBoxForError := func() string {
		return fmt.Sprintf(
			"[xMin=%d, yMin=%d, xMax=%d, yMax=%d]",
			bBox.Min.X,
			bBox.Min.Y,
			bBox.Max.X,
			bBox.Max.Y,
		)
	}
	switch {
	case x < bBox.Min.X || x > bBox.Max.X:
		err := fmt.Errorf(
			"x value %d outside image bounding box %s",
			x,
			fmtBBoxForError(),
		)
		return intZeroVal, err
	case y < bBox.Min.Y || y > bBox.Max.Y:
		err := fmt.Errorf(
			"y value %d outside image bounding box %s",
			y,
			fmtBBoxForError(),
		)
		return intZeroVal, err
	default:
		return x + y*img.Stride, nil
	}
}

// fillerContent task is to provide a value that did not appear
// in the last byte of the stream.
func fillerContent(s []byte) byte {
	lastByte := s[len(s)-1]
	if lastByte != paddingEOT {
		return paddingEOT
	}
	return paddingETX
}
