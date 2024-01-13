package mt

import (
	"math"
	"math/rand"
	"testing"
)

func TestDraw1(t *testing.T) {
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
	if err := p.SaveImageInFile("result/result.jpeg", "jpeg"); err != nil {
		t.Error(err)
	}
}

func TestDraw2(t *testing.T) {
	var data []DataValue
	for i := 0; i < 8; i++ {
		data = append(data, DataValue{
			X: uint(math.Abs(rand.Float64()) * 10),
			Y: uint(math.Abs(rand.Float64()) * 10),
		})
	}

	p2 := NewPlot(data)
	p2.SaveDataInImage(1000, 500)
	if err := p2.SaveImageInFile("result/result2.jpeg", "jpeg"); err != nil {
		t.Error(err)
	}
}
