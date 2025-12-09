package main

import (
	"flag"
	"fmt"
	"sort"

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
	reds := readRedTiles(lines)
	maxArea := 0
	for i := 0; i < len(reds); i++ {
		for j := i + 1; j < len(reds); j++ {
			area := rectArea(reds[i], reds[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func readRedTiles(lines []string) []u.Pos2D {
	var tiles []u.Pos2D
	for _, line := range lines {
		var col, row int
		_, err := fmt.Sscanf(line, "%d,%d", &col, &row)
		if err != nil {
			panic(err)
		}
		tiles = append(tiles, u.Pos2D{Row: row, Col: col})
	}
	return tiles
}

func rectArea(p1, p2 u.Pos2D) int {
	width := u.Abs(p1.Col-p2.Col) + 1
	height := u.Abs(p1.Row-p2.Row) + 1
	return width * height
}

// Not rightt: 4591195600
// Too low: 1291825827

func task2(lines []string) int {
	reds := readRedTiles(lines)
	edges := edgesByLength(reds)
	longestEdge := edges[0]
	fmt.Println("Longest edge: ", longestEdge, " length: ", u.ManhattanDistance(longestEdge.a, longestEdge.b))
	horizontal := longestEdge.a.Row == longestEdge.b.Row
	if horizontal {
		fmt.Println("Longest edge is horizontal")
	} else {
		fmt.Println("Longest edge is vertical")
	}
	var tl, br u.Pos2D
	maxArea := 0
	for i := 0; i < len(reds); i++ {
		for j := i + 1; j < len(reds); j++ {
			// This is only a valid rectangle if
			// the border (other red tiles) are spanning
			// the full rectangle.
			// I.e.
			// 1. no interior red tiles
			// 2. Other red tiles spanning the opposite corners'
			// 3. Do not cross the longest edge (which is essentially splitting a sphere in two halfs heuristic)
			minRow := min2(reds[i].Row, reds[j].Row)
			maxRow := max2(reds[i].Row, reds[j].Row)
			minCol := min2(reds[i].Col, reds[j].Col)
			maxCol := max2(reds[i].Col, reds[j].Col)
			if horizontal && minRow < longestEdge.a.Row && maxRow > longestEdge.a.Row {
				continue
			}
			if !horizontal && minCol < longestEdge.a.Col && maxCol > longestEdge.a.Col {
				continue
			}
			redInteriorFound := false
			topLeft := u.Pos2D{Row: minRow, Col: minCol}
			topRight := u.Pos2D{Row: minRow, Col: maxCol}
			bottomLeft := u.Pos2D{Row: maxRow, Col: minCol}
			bottomRight := u.Pos2D{Row: maxRow, Col: maxCol}
			topLeftFound := false
			topRightFound := false
			bottomLeftFound := false
			bottomRightFound := false
			for k := 0; k < len(reds); k++ {
				if reds[k].Col <= topLeft.Col && reds[k].Row <= topLeft.Row {
					topLeftFound = true
				}
				if reds[k].Col >= topRight.Col && reds[k].Row <= topRight.Row {
					topRightFound = true
				}
				if reds[k].Col <= bottomLeft.Col && reds[k].Row >= bottomLeft.Row {
					bottomLeftFound = true
				}
				if reds[k].Col >= bottomRight.Col && reds[k].Row >= bottomRight.Row {
					bottomRightFound = true
				}
				if reds[k].Row > minRow && reds[k].Row < maxRow &&
					reds[k].Col > minCol && reds[k].Col < maxCol {
					redInteriorFound = true
					break
				}
			}
			if !redInteriorFound && topLeftFound && topRightFound && bottomLeftFound && bottomRightFound {
				// Valid rectangle
				area := rectArea(reds[i], reds[j])
				if area > maxArea {
					fmt.Printf("New max area %d. Corners: %#v %#v\n", area, reds[i], reds[j])
					maxArea = area
					tl = topLeft
					br = bottomRight
				}
			}
		}
	}
	for _, red := range reds {
		if red.Row > tl.Row && red.Row < br.Row &&
			red.Col > tl.Col && red.Col < br.Col {
			fmt.Printf("Red interior tile found: %#v\n", red)
		}
	}
	return maxArea
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pair struct {
	a, b u.Pos2D
}

func edgesByLength(boxes []u.Pos2D) []pair {
	nrEdges := len(boxes)
	edges := make([]pair, 0, nrEdges)
	for i := 0; i < nrEdges; i++ {
		j := (i + 1) % nrEdges
		edges = append(edges, pair{boxes[i], boxes[j]})
	}
	sort.Slice(edges, func(i, j int) bool {
		a := edges[i]
		b := edges[j]
		lenA := u.ManhattanDistance(a.a, a.b)
		lenB := u.ManhattanDistance(b.a, b.b)
		return lenA > lenB
	})
	return edges
}
