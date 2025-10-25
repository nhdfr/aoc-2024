package main

import (
	"strings"
)

func solvePart2(input string) interface{} {
	input = strings.TrimSpace(input)
	floor := 0
	
	for i, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		
		if floor == -1 {
			return i + 1  // 1-indexed position
		}
	}
	
	return -1  // Never reached basement
}