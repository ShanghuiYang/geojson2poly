package main

import (
	"fmt"
	"os"

	"./geojson"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("error: invail args")
		fmt.Println("usage: geojson2poly -i infile.json -o outfile.poly")
		return
	}
	var infile string
	var outfile string
	for i, arg := range os.Args {
		if arg == "-i" || arg == "-I" {
			infile = os.Args[i+1]
		}
		if arg == "-o" || arg == "-O" {
			outfile = os.Args[i+1]
		}
	}

	fmt.Printf("input: %v\n", infile)
	fmt.Printf("output: %v\n", outfile)

	g := geojson.NewGeojson()
	err := g.Load(infile)
	if err != nil {
		fmt.Printf("error: fail to load geojson from: %v\n", infile)
		fmt.Printf("err msg: %v\n", err)
	}

	err = g.ToPoly(outfile)
	if err != nil {
		fmt.Printf("error: fail to save poly file: %v\n", outfile)
		fmt.Printf("err msg: %v\n", err)
	}

	fmt.Println("done in success")
	return
}
