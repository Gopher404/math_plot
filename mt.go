package mt

import (
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

var (
	blue  color.Color = color.RGBA{R: 30, G: 144, B: 255, A: 0}
	black color.Color = color.RGBA{R: 0, G: 0, B: 0, A: 0}
)

const (
	bottomPad  = 15
	charWidth  = 8
	charHeight = 9
)

type Plot struct {
	data []DataValue
	img  *image.RGBA
}
type DataValue struct {
	X uint
	Y uint
}

func NewPlot(data []DataValue) *Plot {
	return &Plot{
		data: data,
	}
}

func (p *Plot) AddData(data ...DataValue) {
	p.data = append(p.data, data...)
}

func (p *Plot) SaveDataInImage(width, height int) {
	p.img = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(p.img, p.img.Bounds(), &image.Uniform{C: color.RGBA{R: 255, G: 255, B: 255, A: 0}}, image.Point{}, draw.Src)

	maxValY := getMaxValY(p.data)
	maxValX := getMaxValX(p.data)
	minValY := getMinValY(p.data)
	minValX := getMinValX(p.data)

	//dataLen := len(p.data)

	numWidth := len(fmt.Sprintf("%d", maxValY.Y))*charWidth + 2

	drawLine(p.img, blue, numWidth, 0, numWidth, height-bottomPad)
	drawLine(p.img, blue, numWidth, height-bottomPad, width, height-bottomPad)

	d := &font.Drawer{
		Dst:  p.img,
		Src:  image.NewUniform(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.I(2), Y: fixed.I(10)},
	}

	// set plot

	var xLast, yLast int

	for i := range p.data {
		x := int(float32(width-numWidth-charWidth*2)*float32(p.data[i].X-minValX.X)/float32(maxValX.X-minValX.X) + float32(numWidth))
		y := int(float32(height) - float32(height-bottomPad-charHeight)*float32(p.data[i].Y-minValY.Y)/float32(maxValY.Y-minValY.Y) - float32(bottomPad))

		// num of Y
		ns := fmt.Sprintf("%d", p.data[i].Y)
		lns := len(ns)

		d.Dot.X = fixed.I(numWidth - lns*charWidth - 1)
		d.Dot.Y = fixed.I(y)

		d.DrawString(ns)
		drawLine(p.img, blue, numWidth-2, y, numWidth+2, y)

		//num of X
		ns = fmt.Sprintf("%d", p.data[i].X)
		lns = len(ns)

		d.Dot.X = fixed.I(x - 2)
		d.Dot.Y = fixed.I(height - charHeight + 7)

		d.DrawString(ns)
		drawLine(p.img, blue, x+lns-1, height-charHeight*2+5, x+lns-1, height-charHeight*2+1)

		// draw line
		if i != 0 {
			drawLine(p.img, black, xLast, yLast, x, y)
		}

		xLast, yLast = x, y

	}
}

func (p *Plot) SaveImageInFile(name, MIME string) error {
	MIME = strings.ToLower(MIME)

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	switch MIME {
	case "png":
		if err := png.Encode(f, p.img); err != nil {
			return err
		}
	case "jpeg":
		if err := jpeg.Encode(f, p.img, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown format: %s", MIME)
	}
	return nil

}

func (p *Plot) GetImage() {

}
