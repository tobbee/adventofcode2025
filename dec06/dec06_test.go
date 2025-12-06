package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	u "github.com/tobbee/adventofcode2025/utils"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("testinput")
	result := task1(lines)
	require.Equal(t, 4277556, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFileNoTrim("testinput")
	result := task2(lines)
	require.Equal(t, 3263827, result)
}
