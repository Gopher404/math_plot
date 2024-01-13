package mt

import (
	"testing"
)

func TestDraw(t *testing.T) {
	data := []DataValue{
		{
			1, 10,
		},
		{
			2, 5,
		},
		{
			3, 2,
		},
		{
			4, 7,
		},
		{
			5, 5,
		},
		{
			6, 6,
		},
		{
			7, 2,
		},
		{
			8, 5,
		},
		{
			9, 10,
		},
	}
	p := NewPlot(data)
	p.SaveDataInImage(1000, 500)
	if err := p.SaveImageInFile("result.jpeg", "jpeg"); err != nil {
		t.Error(err)
	}
}
