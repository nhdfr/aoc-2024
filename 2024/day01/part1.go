package main

import (
	"sort"
	"strconv"
	"strings"
)

func solvePart1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left, right []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	result := 0
	for i := range left {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		result += diff
	}

	return result
}
