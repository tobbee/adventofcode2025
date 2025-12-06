package main

import (
	"flag"
	"fmt"
	"strings"

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
	ranges, ingredients := parseInput(lines)
	nrFresh := 0
	for _, ingredient := range ingredients {
		if isFresh(ingredient, ranges) {
			nrFresh++
		}
	}
	return nrFresh
}

type rng struct{ min, max int }

func parseInput(lines []string) ([]rng, []int) {
	ranges := make([]rng, 0)
	ingredients := make([]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			ranges = append(ranges, rng{u.Atoi(parts[0]), u.Atoi(parts[1])})
			continue
		}
		if line == "" {
			continue
		}
		ingredients = append(ingredients, u.Atoi(line))
	}
	return ranges, ingredients
}

func isFresh(ingredient int, ranges []rng) bool {
	for _, r := range ranges {
		if ingredient >= r.min && ingredient <= r.max {
			return true
		}
	}
	return false
}

func task2(lines []string) int {
	ranges, _ := parseInput(lines)
	merged := mergeAll(ranges)
	fmt.Println("Merged Ranges:", len(merged))
	total := 0
	for _, m := range merged {
		total += m.max - m.min + 1
	}
	return total
}

func mergeAll(ranges []rng) []rng {
	merges := make([]rng, 0)
	fmt.Println("Nr Ranges:", len(ranges))
	for {
		firstMerged, rest := mergeFirst(ranges)
		merges = append(merges, firstMerged)
		fmt.Println("Merged:", firstMerged, "Remaining:", len(rest))
		if len(rest) == 0 {
			break
		}
		ranges = rest
	}
	return merges
}

func mergeFirst(ranges []rng) (rng, []rng) {
	first := ranges[0]
	used := u.CreateSet[int]()
	used.Add(0)
	for {
		nr_overlaps := 0
		for i, r := range ranges[1:] {
			if used.Contains(i + 1) {
				continue
			}
			if overlaps(first, r) {
				first = merge(first, r)
				used.Add(i + 1)
				nr_overlaps++
			}
		}
		if nr_overlaps == 0 {
			break
		}
	}
	rest := make([]rng, 0, len(ranges)-len(used))
	for r := range len(ranges) {
		if !used.Contains(r) {
			rest = append(rest, ranges[r])
		}
	}
	return first, rest
}

func merge(r1, r2 rng) rng {
	return rng{min2(r1.min, r2.min), max2(r1.max, r2.max)}
}

func overlaps(r1, r2 rng) bool {
	return r1.min <= r2.max && r2.min <= r1.max
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
