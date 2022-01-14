package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (Image) ColorModel() color.Model { return color.RGBAModel }
func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 400)
}
func (Image) At(x, y int) color.Color {

	if y%2 == 1 && x%2 == 1 {
		return color.RGBA{uint8((x + y/2) % 255), 40, 145, 255}
	} else if y%2 == 0 && x%2 == 0 {
		return color.RGBA{245, uint8((x * y) % 255), 145, 255}
	} else if y%2 == 1 && x%2 == 0 {
		return color.RGBA{245, 40, uint8((x * y) % 255), 255}
	} else if y%2 == 0 && x%2 == 1 {
		return color.RGBA{82, 234, 145, uint8((x + y/2) % 255)}
	}
	return color.Black
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
