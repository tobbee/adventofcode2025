package main

import (
	"flag"
	"fmt"

	"github.com/tobbee/adventofcode2025/utils"
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
	nrZeros := 0
	pos := 50
	for _, line := range lines {
		dir := line[0]
		steps := utils.Atoi(line[1:])
		switch dir {
		case 'L':
			pos -= steps
		case 'R':
			pos += steps
		}
		pos %= 100
		if pos < 0 {
			pos += 100
		}
		if pos == 0 {
			nrZeros++
		}
	}
	return nrZeros
}

func task2(lines []string) int {
	nrZeros := 0
	pos := 50
	for _, line := range lines {
		dir := line[0]
		steps := utils.Atoi(line[1:])
		nrTurns := utils.Abs(steps / 100)
		deltaStep := steps % 100
		switch {
		case steps == 0:
		case dir == 'L' && pos-deltaStep <= 0 && pos != 0:
			nrZeros += 1
		case dir == 'R' && deltaStep+pos >= 100 && pos != 0:
			nrZeros += 1
		default:
			// no additional zero crossing
		}
		nrZeros += nrTurns
		switch dir {
		case 'L':
			pos -= steps
		case 'R':
			pos += steps
		}
		pos %= 100
		if pos < 0 {
			pos += 100
		}
	}
	return nrZeros
}
