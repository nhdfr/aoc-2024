package main

import (
	"strings"
)

func solvePart2(input string) interface{} {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	rows := len(grid)
	cols := len(grid[0])
	result := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] == 'A' {
				ul := grid[i-1][j-1]
				ur := grid[i-1][j+1]
				dl := grid[i+1][j-1]
				dr := grid[i+1][j+1]
				if ((ul == 'M' && dr == 'S') || (ul == 'S' && dr == 'M')) &&
					((ur == 'M' && dl == 'S') || (ur == 'S' && dl == 'M')) {
					result++
				}
			}
		}
	}

	return result
}
