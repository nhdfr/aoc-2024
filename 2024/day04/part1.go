package main

import (
	"strings"
)

func solvePart1(input string) interface{} {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	rows := len(grid)
	cols := len(grid[0])
	result := 0

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'X' {
				for _, dir := range directions {
					if checkXMAS(grid, i, j, dir[0], dir[1]) {
						result++
					}
				}
			}
		}
	}

	return result
}

func checkXMAS(grid [][]rune, x, y, dx, dy int) bool {
	rows := len(grid)
	cols := len(grid[0])
	target := []rune{'X', 'M', 'A', 'S'}
	for k := 0; k < 4; k++ {
		nx := x + k*dx
		ny := y + k*dy
		if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != target[k] {
			return false
		}
	}
	return true
}
