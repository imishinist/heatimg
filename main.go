package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
)

var (
	input  = flag.String("input", "-", "input file name")
	output = flag.String("output", "out.png", "output image file name")
)

func main() {
	flag.Parse()

	if _, err := os.Stat(*output); err == nil {
		fmt.Printf("file already exists: %v", *output)
		return
	}

	in := inputFile(*input)
	defer in.Close()

	var heatmap HeatMap
	if err := json.NewDecoder(in).Decode(&heatmap); err != nil {
		log.Fatal(err)
	}

	img := &HeatMapImage{
		heatmap, heatmap.max(),
	}

	f, err := os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func inputFile(filename string) *os.File {
	if filename == "-" || filename == "" {
		return os.Stdin
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
