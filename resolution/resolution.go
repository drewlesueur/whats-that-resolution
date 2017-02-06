package resolution

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

func SameColor(a, b color.Color) bool {
	if a == nil {
		if b == nil {
			return true
		}
		return false
	}

	if b == nil {
		return false
	}

	ra, ga, ba, aa := a.RGBA()
	rb, gb, bb, ab := b.RGBA()

	return WithinThreshold(ra, rb) && WithinThreshold(ga, gb) && WithinThreshold(ba, bb) && WithinThreshold(aa, ab)
}

var Threshold = math.Floor(0.1 * 0xFFFF)

func WithinThreshold(a, b uint32) bool {
	aInt := int(a)
	bInt := int(b)

	diff := aInt - bInt
	if diff < 0 {
		diff = -diff
	}
	return float64(diff) <= Threshold
}

func CheckResolution(image image.Image) (w, h, ow, oh int, err error) {
	rectBounds := image.Bounds()
	originalWidth := rectBounds.Max.X - rectBounds.Min.X
	originalHeight := rectBounds.Max.Y - rectBounds.Min.Y
	minPixelWidth := originalWidth
	minPixelHeight := originalHeight

	fmt.Println("orgigWidth:", originalWidth)
	fmt.Println("origHeight:", originalHeight)

	// loop in rows
	var lastColor color.Color
	currentPixelWidth := 0
	currentPixelHeight := 0

	for y := rectBounds.Min.X; y < rectBounds.Max.Y; y++ {
		currentPixelWidth = 0
		lastColor = nil
		for x := rectBounds.Min.X; x < rectBounds.Max.X; x++ {
			color := image.At(x, y)
			//r, g, b, a := color.RGBA()
			if lastColor == nil {
				//fmt.Println("starting", x, y, r, g, b, a)
				currentPixelWidth = 1
			} else if SameColor(color, lastColor) {
				//fmt.Println("same", x, y, r, g, b, a)
				currentPixelWidth += 1
			} else {
				//lr, lg, lb, la := lastColor.RGBA()
				//fmt.Println("different", x, y, r, g, b, a)
				//fmt.Printf("%d-%d=%d\n", r, lr, r-lr)
				//fmt.Printf("%d-%d=%d\n", g, lg, g-lg)
				//fmt.Printf("%d-%d=%d\n", b, lb, b-lb)
				//fmt.Printf("%d-%d=%d\n", a, la, a-la)

				if currentPixelWidth < minPixelWidth {
					minPixelWidth = currentPixelWidth
					if minPixelWidth == 1 {
						fmt.Println("exiting early with a 1 (width)")
						return originalWidth, originalHeight, originalWidth, originalHeight, nil
					}
				}
				currentPixelWidth = 0
			}
			lastColor = color
		}
	}

	// now do columns
	for x := rectBounds.Min.X; x < rectBounds.Max.X; x++ {
		currentPixelHeight = 0
		lastColor = nil
		for y := rectBounds.Min.X; y < rectBounds.Max.Y; y++ {
			color := image.At(x, y)
			if lastColor == nil {
				currentPixelHeight = 1
			} else if SameColor(color, lastColor) {
				currentPixelHeight += 1
			} else {
				if currentPixelHeight < minPixelHeight {
					minPixelHeight = currentPixelHeight
					if minPixelHeight == 1 {
						return originalWidth, originalHeight, originalWidth, originalHeight, nil
					}
				}
				currentPixelHeight = 0
			}
			lastColor = color
		}
	}

	if minPixelWidth < minPixelHeight {
		return originalWidth / minPixelWidth, originalHeight / minPixelWidth, originalWidth, originalHeight, nil
	} else {
		return originalWidth / minPixelHeight, originalHeight / minPixelHeight, originalWidth, originalHeight, nil
	}
}
