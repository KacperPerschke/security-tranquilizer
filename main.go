package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

var (
	myColForOne     = color.White
	myColForPadding = color.Gray16{0x7fff}
	myColForZero    = color.Black
	myPalette       = []color.Color{
		myColForOne,
		myColForPadding,
		myColForZero,
	}
)

func main() {
	size, err := resBySize(37000)
	if err != nil {
		panic(err)
	}
	height := size.Height
	width := size.Width
	img := image.NewPaletted(
		image.Rect(
			0, 0, // xMin, yMin
			width, height, // xMax, yMax
		),
		myPalette,
	)

	rand.Seed(time.Now().Unix())
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			img.SetColorIndex(x, y, uint8(rand.Intn(3)))
		}
	}
	f, _ := os.Create("image.png")
	defer f.Close()
	png.Encode(f, img)
}
