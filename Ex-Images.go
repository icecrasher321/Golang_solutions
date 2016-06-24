package main
// URL -> https://tour.golang.org/methods/25
import ("golang.org/x/tour/pic"
        "image"
        "image/color"
)

type Image struct{
	Width int
	Height int
	Colour uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.Colour, i.Colour, 255, 255}
}



func main() {
	m := Image{50, 50, 7}
	pic.ShowImage(m)
}
