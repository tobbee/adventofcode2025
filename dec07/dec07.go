package main

import (
	"flag"
	"fmt"

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
	g := u.CreateCharGridFromLines(lines)
	s := findStart(g)
	beams := u.CreateSet[int]()
	beams.Add(s)
	totalNrSplits := 0
	for row := 1; row < g.Height; row++ {
		var nrSplits int
		nrSplits, beams = processRow(g, beams, row)
		totalNrSplits += nrSplits
	}
	return totalNrSplits
}

func processRow(g u.Grid2D[string], beams u.Set[int], row int) (int, u.Set[int]) {
	nrSplits := 0
	newRows := u.CreateSet[int]()
	for _, col := range beams.Values() {
		if g.At(row, col) == "^" {
			nrSplits++
			newRows.Add(col - 1)
			newRows.Add(col + 1)
		} else {
			newRows.Add(col)
		}
	}
	return nrSplits, newRows
}

func findStart(g u.Grid2D[string]) int {
	for x := 0; x < g.Width; x++ {
		if g.At(0, x) == "S" {
			return x
		}
	}
	panic("no start found")
}

func task2(lines []string) int {
	g := u.CreateCharGridFromLines(lines)
	s := findStart(g)
	nrPathsToPos := make(map[int]int)
	nrPathsToPos[s] = 1
	for row := 1; row < g.Height; row++ {
		nrPathsToPos = processQuantumRow(g, nrPathsToPos, row)
	}
	totalNrPaths := 0
	for _, v := range nrPathsToPos {
		totalNrPaths += v
	}
	return totalNrPaths
}

// find number of paths to each position in the given row from the previous row
func processQuantumRow(g u.Grid2D[string], nrPathsToPos map[int]int, row int) map[int]int {
	newNrPathsToPos := make(map[int]int)
	for col, nrPaths := range nrPathsToPos {
		if g.At(row, col) == "^" {
			newNrPathsToPos[col-1] += nrPaths
			newNrPathsToPos[col+1] += nrPaths
		} else {
			newNrPathsToPos[col] += nrPaths
		}
	}
	return newNrPathsToPos
}
