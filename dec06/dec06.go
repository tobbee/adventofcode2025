package main

import (
	"flag"
	"fmt"

	u "github.com/tobbee/adventofcode2025/utils"
)

func main() {
	lines := u.ReadLinesFromFileNoTrim("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

func task1(lines []string) int {
	problems := parseLines1(lines)
	return sumProblems(problems)
}

func sumProblems(problems []problem) int {
	totalSum := 0
	for _, p := range problems {
		switch p.op {
		case "+":
			sum := 0
			for _, n := range p.numbers {
				sum += n
			}
			totalSum += sum
		case "*":
			prod := 1
			for _, n := range p.numbers {
				prod *= n
			}
			totalSum += prod
		}
	}
	return totalSum
}

type problem struct {
	numbers []int
	op      string
}

func parseLines1(lines []string) []problem {
	allParts := make([][]string, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineParts := u.SplitWithSpace(line)
		allParts = append(allParts, lineParts)
	}
	var problems []problem
	for row, lineParts := range allParts {
		items := make([]string, 0)
		for _, p := range lineParts {
			if p == "" {
				continue
			}
			items = append(items, p)
		}
		if len(problems) == 0 {
			problems = make([]problem, len(items))
		}
		pNr := 0
		lastRow := len(allParts) - 1
		for _, p := range items {
			if row != lastRow {
				num := u.Atoi(p)
				problems[pNr].numbers = append(problems[pNr].numbers, num)
			} else {
				problems[pNr].op = p
			}
			pNr++
		}
	}
	return problems
}

func task2(lines []string) int {
	problems := parseLines2(lines)
	return sumProblems(problems)
}

func parseLines2(lines []string) []problem {
	problems := make([]problem, 0)
	// Look at last lines for operations and positions
	col := 0
	opRow := len(lines) - 1
	var currProblem problem
	for {
		if col >= len(lines[opRow]) {
			problems = append(problems, currProblem)
			break
		}
		if isEmptyCol(lines, col) {
			problems = append(problems, currProblem)
			col++
			continue
		}
		switch lines[opRow][col] {
		case '+':
			currProblem.op = "+"
			currProblem.numbers = make([]int, 0)
		case '*':
			currProblem.op = "*"
			currProblem.numbers = make([]int, 0)
		default:
			// Stay with current problem
		}
		number := 0
		power := 1
		for r := opRow - 1; r >= 0; r-- {
			if lines[r][col] == ' ' {
				continue
			}
			digit := lines[r][col] - '0'
			number += int(digit) * power
			power *= 10
		}
		currProblem.numbers = append(currProblem.numbers, number)
		col++
	}

	return problems
}

func isEmptyCol(lines []string, col int) bool {
	for _, line := range lines {
		if line[col] != ' ' {
			return false
		}
	}
	return true
}
