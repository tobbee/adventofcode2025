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
		fmt.Println("task1: ", taskX(lines, 2))
	} else {
		fmt.Println("task2: ", taskX(lines, 12))
	}
}

func taskX(lines []string, n int) int {
	sum := 0
	for _, line := range lines {
		bc := u.SplitToChars(line)
		batteries := make([]int, len(bc))
		for i, b := range bc {
			batteries[i] = u.Atoi(b)
		}
		biggest := biggestJoltageN(batteries, n)
		sum += biggest
	}
	return sum
}

func biggestJoltageN(batteries []int, n int) int {
	digits := make([]int, n)
	lastIndex := -1
	for d := 0; d < n; d++ {
		digits[d] = -1
		for i := lastIndex + 1; i < len(batteries)-(n-1-d); i++ {
			if batteries[i] > digits[d] {
				digits[d] = batteries[i]
				lastIndex = i
			}
		}
	}
	joltage := 0
	for d := 0; d < n; d++ {
		joltage += digits[d] * pow10(n-1-d)
	}
	return joltage
}

func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
