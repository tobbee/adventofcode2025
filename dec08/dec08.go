package main

import (
	"flag"
	"fmt"
	"math"
	"sort"

	u "github.com/tobbee/adventofcode2025/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("task1: ", task1(lines, 1000))
	} else {
		fmt.Println("task2: ", task2(lines))
	}
}

func task1(lines []string, n int) int {
	boxes := parseJunctionBoxes(lines)
	dists := findNClosestPairs(boxes, n)
	circuits := make([]u.Set[Point3D], 0)
	for _, pd := range dists {
		a, b := pd.p.A, pd.p.B
		new := true
		for _, c := range circuits {
			if c.Contains(a) || c.Contains(b) {
				c.Add(a)
				c.Add(b)
				new = false
				break
			}
		}
		if new {
			c := u.CreateSet[Point3D]()
			c.Add(a)
			c.Add(b)
			circuits = append(circuits, c)
		}
	}
	circuits = mergeAll(circuits)

	circuitSizes := make([]int, 0, len(circuits))
	for _, c := range circuits {
		circuitSizes = append(circuitSizes, len(c))
	}
	sort.Ints(circuitSizes)
	prod := circuitSizes[len(circuitSizes)-1] * circuitSizes[len(circuitSizes)-2] * circuitSizes[len(circuitSizes)-3]
	return prod
}

func mergeAll(circuits []u.Set[Point3D]) []u.Set[Point3D] {
	merged := true
	for merged {
		circuits, merged = mergeCircuits(circuits)
	}
	return circuits
}

func mergeCircuits(circuits []u.Set[Point3D]) ([]u.Set[Point3D], bool) {
	merged := false
	newCircuits := make([]u.Set[Point3D], 0, len(circuits))
	used := u.CreateSet[int]()
	for i := 0; i < len(circuits); i++ {
		c1 := circuits[i]
		for j := i + 1; j < len(circuits); j++ {
			c2 := circuits[j]
			if c1.Intersects(c2) {
				c := c1.Clone()
				c.Extend(c2)
				newCircuits = append(newCircuits, c)
				merged = true
				used.Add(j)
				used.Add(i)
				break
			}
		}
		if !used.Contains(i) {
			used.Add(i)
			newCircuits = append(newCircuits, c1)
		}
	}
	return newCircuits, merged
}

type Point3D struct {
	X, Y, Z int
}

type Pair struct {
	A, B Point3D
}

type pairDistance struct {
	p    Pair
	dist float64
}

func parseJunctionBoxes(lines []string) u.Set[Point3D] {
	boxes := u.CreateSet[Point3D]()
	for _, line := range lines {
		x, y, z := 0, 0, 0
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic(err)
		}
		boxes.Add(Point3D{X: x, Y: y, Z: z})
	}
	return boxes
}

func findNClosestPairs(boxes u.Set[Point3D], n int) []pairDistance {
	distances := make([]pairDistance, 0, n+1)

	boxList := boxes.Values()

	for i := 1; i < len(boxList); i++ {
		for j := 0; j < i; j++ {
			a := boxList[i]
			b := boxList[j]
			distance := EuclidianDistance(a, b)
			pairDistance := pairDistance{
				p:    Pair{A: a, B: b},
				dist: distance,
			}
			if len(distances) == 0 {
				distances = append(distances, pairDistance)
				continue
			}
			index := sort.Search(len(distances), func(i int) bool {
				return distances[i].dist >= distance
			})
			if index == len(distances) {
				distances = append(distances, pairDistance)
			} else {
				distances = append(distances[:index+1], distances[index:]...)
				distances[index] = pairDistance
			}
			if len(distances) > n {
				distances = distances[:n]
			}
		}
	}
	return distances
}

func EuclidianDistance(a, b Point3D) float64 {
	xd := float64(a.X - b.X)
	yd := float64(a.Y - b.Y)
	zd := float64(a.Z - b.Z)
	return math.Sqrt(xd*xd + yd*yd + zd*zd)
}

func task2(lines []string) int {
	boxes := parseJunctionBoxes(lines)
	n := len(boxes.Values())
	var nClosests []pairDistance
	for {
		nClosests = findNClosestPairs(boxes, n)
		a, b, done := mergeUntilOne(nClosests, len(boxes))
		if done {
			fmt.Println("Final pair is", a, "and", b, "n=", n)
			return a.X * b.X
		}
		n *= 10
	}
}

func newCircuit(a, b Point3D) u.Set[Point3D] {
	c := u.CreateSet[Point3D]()
	c.Add(a)
	c.Add(b)
	return c
}

func mergeUntilOne(distances []pairDistance, nrBoxes int) (a, b Point3D, done bool) {
	circuits := make([]u.Set[Point3D], 0)
	circuits = append(circuits, newCircuit(distances[0].p.A, distances[0].p.B))
	for _, pd := range distances[1:] {
		a, b := pd.p.A, pd.p.B
		new := true
		for _, c := range circuits {
			if c.Contains(a) || c.Contains(b) {
				c.Add(a)
				c.Add(b)
				new = false
				break
			}
		}
		circuits = mergeAll(circuits)
		if len(circuits) == 1 && len(circuits[0]) == nrBoxes {
			return a, b, true
		}
		if new {
			circuits = append(circuits, newCircuit(a, b))
		}
	}
	return Point3D{}, Point3D{}, false
}
