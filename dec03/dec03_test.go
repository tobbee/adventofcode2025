package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	u "github.com/tobbee/adventofcode2025/utils"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := taskX(lines, 2)
	require.Equal(t, 357, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := taskX(lines, 12)
	require.Equal(t, 3121910778619, result)
}
