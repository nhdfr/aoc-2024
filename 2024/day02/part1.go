package main

import (
	"strconv"
	"strings"
)

func solvePart1(input string) interface{} {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	result := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		levels := make([]int, len(parts))
		for i, p := range parts {
			levels[i], _ = strconv.Atoi(p)
		}
		if isSafe(levels) {
			result++
		}
	}

	return result
}
