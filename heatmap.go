package main

type HeatMap struct {
	GridLevel    int     `json:"gridLevel"`
	Columns      int     `json:"columns"`
	Rows         int     `json:"rows"`
	MinX         float64 `json:"minX"`
	MinY         float64 `json:"minY"`
	MaxX         float64 `json:"maxX"`
	MaxY         float64 `json:"maxY"`
	CountsInts2D [][]int `json:"counts_ints2D"`
}

func (h *HeatMap) max() int {
	max := 0
	for _, row := range h.CountsInts2D {
		for _, col := range row {
			if max < col {
				max = col
			}
		}
	}
	return max
}
