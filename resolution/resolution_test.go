package resolution

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"testing"
)

func drawRect(image *image.RGBA, _x, _y, w, h int, color color.Color) {
	for x := _x; x < w; x++ {
		for y := _y; y < h; y++ {
			image.Set(x, y, color)
			//fmt.Println(x, y)
		}
	}
}

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

func TestSameColor(t *testing.T) {
	same := SameColor(red, red)
	if !same {
		t.Fatal("should be the same color")
	}

	same = SameColor(color.RGBA{1, 0, 0, 255}, color.RGBA{1, 0, 0, 255})
	if !same {
		t.Fatal("should be the same color")
	}

	same = SameColor(color.RGBA{1, 0, 0, 255}, color.RGBA{99, 0, 0, 255})
	if same {
		t.Fatal("should be the different colors")
	}
}

func TestResolution(t *testing.T) {
	//palette := []color.Color{}
	// img := image.NewPaletted(image.Rect(0, 0, 50, 100), palette)

	img := image.NewRGBA(image.Rect(0, 0, 10, 15))
	drawRect(img, 0, 0, 5, 5, green)
	//img.Set(5, 5, green)
	file, err := os.Create("test2.gif")

	if err != nil {
		t.Fatal("could not save image")
	}

	gif.Encode(file, img, &gif.Options{NumColors: 256})

	w, h, _, _, err := CheckResolution(img)
	fmt.Println("the wh", w, h)

	if err != nil {
		t.Fatal("there was an error trying to get the resolution")
	}

	if w != 2 {
		t.Fatal("Incorrect width")
	}

	if h != 3 {
		t.Fatal("Incorrect height")
	}
}

func TestResolution2(t *testing.T) {
	//palette := []color.Color{}
	// img := image.NewPaletted(image.Rect(0, 0, 50, 100), palette)

	img := image.NewRGBA(image.Rect(0, 0, 10, 16))
	drawRect(img, 0, 0, 5, 2, green)
	//img.Set(5, 5, green)
	file, err := os.Create("test2.gif")

	if err != nil {
		t.Fatal("could not save image")
	}

	gif.Encode(file, img, &gif.Options{NumColors: 256})

	w, h, _, _, err := CheckResolution(img)
	fmt.Println("the wh", w, h)

	if err != nil {
		t.Fatal("there was an error trying to get the resolution")
	}

	if w != 5 {
		t.Fatal("Incorrect width")
	}

	if h != 8 {
		t.Fatal("Incorrect height")
	}
}
