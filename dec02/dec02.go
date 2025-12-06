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
	ranges := splitRanges(lines[0])
	invalid_sum := 0
	for _, r := range ranges {
		for n := r[0]; n <= r[1]; n++ {
			n_str := fmt.Sprintf("%d", n)
			if len(n_str)%2 == 1 {
				continue
			}
			half_len := len(n_str) / 2
			first_half := n_str[:half_len]
			second_half := n_str[half_len:]
			if first_half == second_half {
				invalid_sum += n
			}
		}
	}
	return invalid_sum
}

func splitRanges(s string) [][2]int {
	var result [][2]int
	var start, end int
	parts := strings.Split(s, ",")
	for _, part := range parts {
		fmt.Sscanf(part, "%d-%d", &start, &end)
		result = append(result, [2]int{start, end})
	}
	return result
}

func task2(lines []string) int {
	ranges := splitRanges(lines[0])
	invalid_sum := 0
	for _, r := range ranges {
		for n := r[0]; n <= r[1]; n++ {
			n_str := fmt.Sprintf("%d", n)
			for l := 1; l < len(n_str); l++ {
				if len(n_str)%l != 0 {
					continue
				}
				nr_parts := len(n_str) / l
				rep_part := n_str[:l]
				rep := ""
				for i := 0; i < nr_parts; i++ {
					rep += rep_part
				}
				if rep == n_str {
					invalid_sum += n
					break
				}
			}
		}
	}
	return invalid_sum
}
