package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoolGrid(t *testing.T) {
	g := CreateGrid2D[bool](4, 2)
	g.Set(true, 0, 1)
	require.Equal(t, true, g.Get(0, 1))
	r, c, ok := g.Find(true)
	require.Equal(t, 0, r)
	require.Equal(t, 1, c)
	require.Equal(t, true, ok)
}

func TestCreateCharGridFromLines(t *testing.T) {
	lines := []string{"abc", "def"}
	g := CreateCharGridFromLines(lines)
	require.Equal(t, 3, g.Width)
	require.Equal(t, 2, g.Height)
	require.Equal(t, "a", g.At(0, 0))
	require.Equal(t, "c", g.At(0, 2))
	require.Equal(t, "f", g.At(1, 2))
}

func TestCreateDigitGridFromLines(t *testing.T) {
	lines := []string{"123", "456"}
	g := CreateDigitGridFromLines(lines)
	require.Equal(t, 3, g.Width)
	require.Equal(t, 2, g.Height)
	require.Equal(t, 1, g.At(0, 0))
	require.Equal(t, 3, g.At(0, 2))
	require.Equal(t, 6, g.At(1, 2))
}

func TestCreateRuneGridFromLines(t *testing.T) {
	lines := []string{"abc", "def"}
	g := CreateRuneGridFromLines(lines)
	require.Equal(t, 3, g.Width)
	require.Equal(t, 2, g.Height)
	require.Equal(t, 'a', g.At(0, 0))
	require.Equal(t, 'c', g.At(0, 2))
	require.Equal(t, 'f', g.At(1, 2))
}
