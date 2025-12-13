package main

import (
	"flag"
	"fmt"
	"strings"

	u "github.com/tobbee/adventofcode2025/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

func task1(lines []string) int {
	shapes, regions := parseInput(lines)
	fmt.Println("Shapes:", shapes)
	var smallerThanRegion []int
	for reg, region := range regions {
		totalTileArea := 0
		totalGridArea := region.W * region.H
		for i, nrShapes := range region.NrShapes {
			shape := shapes[i]
			totalTileArea += nrShapes * shape.nrTiles
		}
		fmt.Println("reg:", reg, "Diff", totalGridArea-totalTileArea, "Total tile area:", totalTileArea, "Total grid area:", totalGridArea)
		if totalTileArea < totalGridArea-350 {
			smallerThanRegion = append(smallerThanRegion, reg)
		}
	}
	fmt.Println("Regions smaller than grid area:", len(smallerThanRegion), "of total", len(regions))

	return 0
}

type Shape struct {
	nrTiles int
	tiles   u.Grid2D[string]
}

type Region struct {
	W, H     int
	NrShapes []int
}

func parseInput(lines []string) ([]Shape, []Region) {
	var shapes []Shape
	var regions []Region
	shapeLines := make([]string, 0, 3)
	parsingShapes := true
	for _, line := range lines {
		if strings.Contains(line, "x") {
			parsingShapes = false
		}
		if parsingShapes {
			switch {
			case strings.Contains(line, ":"):
				shapeLines = shapeLines[:0]
			case line == "":
				g := u.CreateCharGridFromLines(shapeLines)
				nrTiles := 0
				for r := range g.Height {
					for c := range g.Width {
						if g.At(r, c) == "#" {
							nrTiles++
						}
					}
				}
				shape := Shape{
					nrTiles: nrTiles,
					tiles:   g,
				}
				shapes = append(shapes, shape)
			default:
				shapeLines = append(shapeLines, line)
			}
			continue
		}
		// Parsing regions
		parts := u.SplitWithSpace(line)
		w, h, _ := strings.Cut(parts[0][:len(parts[0])-1], "x")
		width, height := u.Atoi(w), u.Atoi(h)
		nrShapes := make([]int, len(parts)-1)
		for i, ns := range parts[1:] {
			nrShapes[i] = u.Atoi(ns)
		}
		regions = append(regions, Region{
			W:        width,
			H:        height,
			NrShapes: nrShapes,
		})
	}
	return shapes, regions
}

func task2(lines []string) int {
	return 0
}
