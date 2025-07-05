package colorutil

import (
	"fmt"
	"image"
	"image/color"
)

func BalanceColor(im image.Image, balance uint32) image.Image {
	bounds := im.Bounds()
	fmt.Printf("bounds: %v\n", bounds)

	minX, minY := bounds.Min.X, bounds.Min.Y
	maxX, maxY := bounds.Max.X, bounds.Max.Y

	processedImage := image.NewRGBA(im.Bounds())

	for i := minX; i < maxX; i++ {
		for j := minY; j < maxY; j++ {
			r, g, b, a := im.At(i, j).RGBA()
			newR := uint8(((balance * r) >> 8))
			newG := uint8(((balance * g) >> 8))
			newB := uint8(((balance * b) >> 8))
			newA := uint8(a >> 8)
			processedImage.SetRGBA(i, j, color.RGBA{R: newR, G: newG, B: newB, A: newA})
		}
	}
	return processedImage
}
