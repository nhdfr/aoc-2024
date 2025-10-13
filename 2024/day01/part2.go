package main

import (
	"strconv"
	"strings"
)

func solvePart2(input string) interface{} {
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

	rightCount := make(map[int]int)
	for _, r := range right {
		rightCount[r]++
	}

	result := 0
	for _, l := range left {
		result += l * rightCount[l]
	}

	return result
}
