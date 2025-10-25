package main

import (
	"strings"
)

func solvePart1(input string) interface{} {
	input = strings.TrimSpace(input)
	floor := 0
	
	for _, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	
	return floor
}