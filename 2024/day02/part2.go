package main

import (
	"strconv"
	"strings"
)

func solvePart2(input string) interface{} {
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
			continue
		}
		safe := false
		for i := 0; i < len(levels); i++ {
			newLevels := make([]int, 0, len(levels)-1)
			newLevels = append(newLevels, levels[:i]...)
			newLevels = append(newLevels, levels[i+1:]...)
			if isSafe(newLevels) {
				safe = true
				break
			}
		}
		if safe {
			result++
		}
	}

	return result
}
