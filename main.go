package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

const (
	height = 120
	width  = 160
)

var (
	myColBlack = color.Black
	myColGray  = color.Gray16{0x7fff}
	myColWhite = color.White
	myPalette  = []color.Color{
		myColBlack,
		myColGray,
		myColWhite,
	}
	/*
		myPalIdxBlack = 0
		myPalIdxGray  = 1
		myPalIdxWhite = 2
	*/
)

func main() {
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
