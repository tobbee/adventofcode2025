package main

import (
	"flag"
	"fmt"
	"math"

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
	machines := make([]machine, 0)
	for _, line := range lines {
		m := parseLine(line)
		machines = append(machines, m)
	}
	total := 0
	for _, m := range machines {
		n := m.matchIndicatorTarget()
		total += n
	}
	return total
}

func (m *machine) matchIndicatorTarget() int {
	// Try all button combinations
	numButtons := len(m.buttons)
	numCombinations := 1 << numButtons
	leastPresses := math.MaxInt
	for comb := 0; comb < numCombinations; comb++ {
		m.reset()
		nrPresses := 0
		for btnIdx := range numButtons {
			if (comb & (1 << btnIdx)) != 0 {
				m.pushLightButton(btnIdx)
				nrPresses++
			}
		}
		if m.matchesLightTarget() && nrPresses < leastPresses {
			leastPresses = nrPresses
		}
	}
	return leastPresses
}

type machine struct {
	nrLights  int
	target    []bool
	indicator []bool
	buttons   [][]bool
	zoltages  []int
	counters  []int
}

func (m *machine) reset() {
	for i := range m.indicator {
		m.indicator[i] = false
		m.counters = make([]int, len(m.zoltages))
	}
}

func (m *machine) pushLightButton(btnIdx int) {
	for lightIdx, toggles := range m.buttons[btnIdx] {
		if toggles {
			m.indicator[lightIdx] = !m.indicator[lightIdx]
		}
	}
}

func (m *machine) matchesLightTarget() bool {
	for i, lightOn := range m.indicator {
		if lightOn != m.target[i] {
			return false
		}
	}
	return true
}

func task2(_ []string) int {
	// Done in Python using SciPy linear programming MILP solver
	return 19210
}

// A line lookes like [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
func parseLine(line string) machine {
	m := machine{
		buttons: make([][]bool, 0),
	}
	parts := u.SplitWithSpace(line)
	targetStr := parts[0]
	m.target = make([]bool, len(targetStr)-2)
	for i, ch := range targetStr[1 : len(targetStr)-1] {
		if ch == '#' {
			m.target[i] = true
		}
	}
	m.nrLights = len(m.target)
	m.indicator = make([]bool, m.nrLights)
	for i := 1; i < len(parts)-1; i++ {
		btnPart := parts[i]
		btnPart = btnPart[1 : len(btnPart)-1]
		btnNums := u.SplitToInts(btnPart)
		toggles := make([]bool, m.nrLights)
		for _, n := range btnNums {
			toggles[n] = true
		}
		m.buttons = append(m.buttons, toggles)
	}
	zoltagePart := parts[len(parts)-1]
	zoltagePart = zoltagePart[1 : len(zoltagePart)-1]
	m.zoltages = u.SplitToInts(zoltagePart)
	return m
}
