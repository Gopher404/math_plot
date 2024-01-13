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
	"strconv"
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
	draw.Draw(p.img, p.img.Bounds(), &image.Uniform{C: color.RGBA{R: 255, G: 255, B: 255, A: 0}}, image.ZP, draw.Src)
	maxValY := getMaxValY(p.data)
	maxValX := getMaxValX(p.data)

	dataLen := len(p.data)

	numWidth := len(strconv.Itoa(int(p.data[maxValY].Y)))*charWidth + 2

	drawLine(p.img, blue, numWidth, 0, numWidth, height-bottomPad)
	drawLine(p.img, blue, numWidth, height-bottomPad, width, height-bottomPad)

	d := &font.Drawer{
		Dst:  p.img,
		Src:  image.NewUniform(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.I(2), Y: fixed.I(10)},
	}

	// set of Y
	toSortY := make([]DataValue, dataLen)
	copy(toSortY, p.data)
	fmt.Println(toSortY)
	var minValY DataValue
	for i, v := range sortValues(toSortY, "Y") {
		if i == 0 {
			minValY = v
		}
		ns := fmt.Sprintf("%d", v.Y)
		lns := len(ns)
		x := numWidth - lns*charWidth
		y := int(float64(height) - float64(height-bottomPad-charHeight)*float64(v.Y-minValY.Y)/float64(p.data[maxValY].Y-minValY.Y) - float64(bottomPad))

		d.Dot = fixed.Point26_6{X: fixed.I(x - 1), Y: fixed.I(y + charHeight/2)}
		d.DrawString(ns)

		drawLine(p.img, blue, x+lns*charWidth-2, y, x+lns*charWidth+2, y)

	}

	// set of X
	toSortX := make([]DataValue, dataLen)
	copy(toSortX, p.data)

	var minValX DataValue

	for i, v := range sortValues(toSortX, "X") {
		if i == 0 {
			minValX = v
		}
		ns := fmt.Sprintf("%d", v.X)
		lns := len(ns)

		x := int(float64(width-numWidth-charWidth*2)*float64(v.X-minValX.X)/float64(p.data[maxValX].X-minValX.X) + float64(numWidth-1))
		y := height - charHeight + 5
		fmt.Println(v, x)
		d.Dot = fixed.Point26_6{X: fixed.I(x - 2), Y: fixed.I(y + 2)}
		d.DrawString(ns)

		drawLine(p.img, blue, x+lns, y-charHeight, x+lns, y-charHeight-4)

	}

	// set plot
	for i := 1; i < dataLen; i++ {
		drawLine(p.img, black,
			int(float64(width-numWidth-charWidth*2)*float64(p.data[i-1].X-minValX.X)/float64(p.data[maxValX].X-minValX.X)+float64(numWidth)),
			int(float64(height)-float64(height-bottomPad-charHeight)*float64(p.data[i-1].Y-minValY.Y)/float64(p.data[maxValY].Y-minValY.Y)-float64(bottomPad)),
			int(float64(width-numWidth-charWidth*2)*float64(p.data[i].X-minValX.X)/float64(p.data[maxValX].X-minValX.X)+float64(numWidth)),
			int(float64(height)-float64(height-bottomPad-charHeight)*float64(p.data[i].Y-minValY.Y)/float64(p.data[maxValY].Y-minValY.Y)-float64(bottomPad)),
		)

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
