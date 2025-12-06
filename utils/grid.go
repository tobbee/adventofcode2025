package utils

import (
	"strconv"
)

type Grid2D[T comparable] struct {
	Grid   [][]T
	Width  int
	Height int
}

func CreateGrid2D[T comparable](width, height int) Grid2D[T] {
	grid := Grid2D[T]{
		Grid:   make([][]T, 0, height),
		Width:  width,
		Height: height}

	for i := 0; i < grid.Height; i++ {
		grid.Grid = append(grid.Grid, make([]T, grid.Width))
	}
	return grid
}

// InBounds checks if (row, col) inside grid
func (g Grid2D[T]) InBounds(row, col int) bool {
	return 0 <= row && row < g.Height && 0 <= col && col < g.Width
}

// AtBorder checks if (row, col) is at the border of the grid
func (g Grid2D[T]) AtBorder(row, col int) bool {
	return row == 0 || row == g.Height-1 || col == 0 || col == g.Width-1
}

// SetAll sets all cells in the grid to the given value
func (g *Grid2D[T]) SetAll(value T) {
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			g.Grid[row][col] = value
		}
	}
}

func (g *Grid2D[T]) Find(value T) (row, col int, ok bool) {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == value {
				return r, c, true
			}
		}
	}
	return -1, -1, false
}

func (g *Grid2D[T]) At(row, col int) T {
	return g.Grid[row][col]
}

// Set sets an individual cell in the grid
func (g *Grid2D[T]) Set(val T, row, col int) {
	g.Grid[row][col] = val
}

func (g Grid2D[T]) Get(row, col int) T {
	return g.Grid[row][col]
}

// Factory functions for creating Grid2D with specific types

func CreateDigitGridFromLines(lines []string) Grid2D[int] {
	if len(lines) == 0 {
		return CreateGrid2D[int](0, 0)
	}
	width := len(lines[0])
	g := CreateGrid2D[int](width, len(lines))
	for i, line := range lines {
		if len(line) != width {
			panic("non-rectangular grid")
		}
		digits := SplitToChars(line)
		for j, digit := range digits {
			nr, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}
			g.Set(nr, i, j)
		}
	}
	return g
}

func CreateCharGridFromLines(lines []string) Grid2D[string] {
	if len(lines) == 0 {
		return CreateGrid2D[string](0, 0)
	}
	width := len(lines[0])
	g := CreateGrid2D[string](width, len(lines))
	for i, line := range lines {
		if len(line) != width {
			panic("non-rectangular grid")
		}
		row := SplitToChars(line)
		for j, char := range row {
			g.Set(char, i, j)
		}
	}
	return g
}

func CreateRuneGridFromLines(lines []string) Grid2D[rune] {
	if len(lines) == 0 {
		return CreateGrid2D[rune](0, 0)
	}
	width := len(lines[0])
	g := CreateGrid2D[rune](width, len(lines))
	for i, line := range lines {
		if len(line) != width {
			panic("non-rectangular grid")
		}
		row := SplitToRunes(line)
		for j, r := range row {
			g.Set(r, i, j)
		}
	}
	return g
}

func CreateZeroDigitGrid(width, height int) Grid2D[int] {
	return CreateGrid2D[int](width, height)
}
