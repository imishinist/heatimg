# -*- coding: utf-8 -*-

import json
import math

"""
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
"""

def main():
    maps = {}
    cols, rows = 1001, 1001
    maps["columns"] = cols
    maps["rows"] = rows

    counts = [[0 for i in range(cols)] for j in range(rows)]

    # print(len(counts), len(counts[0]))

    for y in range(rows):
        for x in range(cols):
            xd = x - cols / 2 + 1
            yd = y - rows / 2 + 1
            dis = (math.sqrt(xd * xd + yd * yd))
            counts[y][x] = (int(dis) % 64) * 64

    maps["counts_ints2D"] = counts
    print(json.dumps(maps))


if __name__ == "__main__":
    main()
