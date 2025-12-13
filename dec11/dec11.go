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

func parseInput(lines []string) map[string][]string {
	conns := make(map[string][]string)
	for _, line := range lines {
		parts := u.SplitWithSpace(line)
		input := parts[0][0 : len(parts[0])-1]
		outputs := parts[1:]
		conns[input] = outputs
	}
	return conns
}

func task1(lines []string) int {
	conns := parseInput(lines)
	fmt.Println(conns)
	return nrTraversals("you", conns, u.CreateSet[string]())
}

func nrTraversals(curr string, conns map[string][]string, visited u.Set[string]) int {
	if curr == "out" {
		return 1
	}
	nextLevel := visited.Clone()
	nextLevel.Add(curr)
	endTraversals := 0
	for _, next := range conns[curr] {
		if visited.Contains(next) {
			continue
		}
		endTraversals += nrTraversals(next, conns, nextLevel)
	}
	return endTraversals
}

func task2(lines []string) int {
	conns := parseInput(lines)
	narrowSections := [][]string{
		{"cgb", "zod", "jur"},
		// {"fft"}, // Needs to pass this single node
		{"lfr", "vfm", "xsh", "dhl"},
		{"yvu", "qcf", "mrm", "hvt"},
		{"nid", "xxa", "ptd", "icr"},
		//{"dac"}, // Needs to pass this single node
		{"dev", "gbd", "npt", "oki", "you"},
		{"out"},
	}
	startPoints := []string{"svr"}
	nrPaths := make(map[string]int)
	nrPaths["svr"] = 1
	for _, ns := range narrowSections {
		for _, start := range startPoints {
			for _, nn := range ns {
				target := nn
				visited := u.CreateSet[string]()
				for _, nn2 := range ns {
					if nn2 != nn {
						visited.Add(nn2)
					}
				}
				nrP := nrTraversals2(start, target, conns, visited)
				nrPaths[target] += nrP * nrPaths[start]
			}
		}
		startPoints = ns
	}
	return nrPaths["out"]
}

func nrTraversals2(curr, to string, conns map[string][]string, visited u.Set[string]) int {
	if curr == to {
		switch to {
		case "lfr", "vfm", "xsh", "dhl":
			if visited.Contains("fft") {
				return 1
			}
			return 0
		case "dev", "gbd", "npt", "oki", "you":
			if visited.Contains("dac") {
				return 1
			}
			return 0
		}
		return 1
	}
	nextLevel := visited.Clone()
	nextLevel.Add(curr)
	endTraversals := 0
	for _, next := range conns[curr] {
		if visited.Contains(next) {
			continue
		}
		endTraversals += nrTraversals2(next, to, conns, nextLevel)
	}
	return endTraversals
}

func setFromStrings(elems ...string) u.Set[string] {
	s := u.CreateSet[string]()
	for _, e := range elems {
		s.Add(e)
	}
	return s
}
