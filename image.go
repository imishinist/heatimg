package main

import (
	"image"
	"image/color"
)

var _ image.Image = (*HeatMapImage)(nil)

type HeatMapImage struct {
	heatmap HeatMap
	max     int
}

func (h *HeatMapImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (h *HeatMapImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: h.heatmap.Columns,
			Y: h.heatmap.Rows,
		},
	}
}

func (h *HeatMapImage) At(x, y int) color.Color {
	normed := float64(h.heatmap.CountsInts2D[y][x]) / float64(h.max)

	return color.RGBA{
		R: 0,
		G: 0,
		B: uint8(normed * (1<<8 - 1)),
		A: 255,
	}
}
