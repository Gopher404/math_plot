package mt

import (
	"image"
	"image/color"
	"math"
)

func drawLine(img *image.RGBA, c color.Color, x0, y0, x1, y1 int) {
	dx := int(math.Abs(float64(x1 - x0)))
	dy := int(math.Abs(float64(y1 - y0)))
	var (
		sx int
		sy int
	)
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	e := dx - dy
	img.Set(x1, y1, c)
	for x0 != x1 || y0 != y1 {
		img.Set(x0, y0, c)
		e2 := e * 2
		if e2 > -dy {
			e -= dy
			x0 += sx
		}
		if e2 < dx {
			e += dx
			y0 += sy
		}
	}
}
