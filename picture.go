package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)


type Image struct{}


func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (i Image) At(x, y int) color.Color {

	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}

func main() {
	// 可以自己设置宽高,传递进去
	m := Image{}
	// 3 调用
	pic.ShowImage(m)
}
