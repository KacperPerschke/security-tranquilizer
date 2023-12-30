package img

import (
	"image"
)

const (
	emptyString = ``
	intZeroVal  = 0
	paddingEOT  = 4 // https://en.wikipedia.org/wiki/End-of-Transmission_character
	paddingETX  = 3 // https://en.wikipedia.org/wiki/End-of-Text_character
)

var (
	emptyImage     = image.NewGray(emptyRectangle)
	emptyRectangle = image.Rect(0, 0, 0, 0)
)

func imgXYToSlicePos(x, y int, img *image.Gray) int {
	return x + y*img.Stride
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
