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
	nrAvail := 0
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			if g.At(row, col) == "." {
				continue
			}
			nrNeighbors := 0
			for r := row - 1; r <= row+1; r++ {
				for c := col - 1; c <= col+1; c++ {
					if r == row && c == col {
						continue
					}
					if g.InBounds(r, c) && g.At(r, c) == "@" {
						nrNeighbors++
					}
				}
			}
			if nrNeighbors < 4 {
				nrAvail++
			}
		}
	}
	return nrAvail
}

func task2(lines []string) int {
	g := u.CreateCharGridFromLines(lines)

	avails := findAvails(g)
	startAvails := len(avails)
	nrAvail := startAvails
	fmt.Println("Starting with", startAvails, "availables")
	removed := 0
	for {
		n := removeAvails(g, avails)
		removed += n
		fmt.Println("Removed", n, "availables, total:", nrAvail)
		avails = findAvails(g)
		nrAvail += len(avails)
		if len(avails) == 0 {
			break
		}
	}

	return removed
}

func findAvails(g u.Grid2D[string]) []u.Pos2D {
	avails := make([]u.Pos2D, 0, 1000)
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			if g.At(row, col) == "." {
				continue
			}
			nrNeighbors := 0
			for r := row - 1; r <= row+1; r++ {
				for c := col - 1; c <= col+1; c++ {
					if r == row && c == col {
						continue
					}
					if g.InBounds(r, c) && g.At(r, c) == "@" {
						nrNeighbors++
					}
				}
			}
			if nrNeighbors < 4 {
				avails = append(avails, u.Pos2D{Row: row, Col: col})
			}
		}
	}
	return avails
}

func removeAvails(g u.Grid2D[string], avails []u.Pos2D) int {
	for _, p := range avails {
		g.Set(".", p.Row, p.Col)
	}
	return len(avails)
}
