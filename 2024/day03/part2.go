package main

import (
	"regexp"
	"strconv"
)

func solvePart2(input string) interface{} {
	result := 0
	enabled := true

	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled && match[2] != "" && match[3] != "" {
			a, _ := strconv.Atoi(match[2])
			b, _ := strconv.Atoi(match[3])
			result += a * b
		}
	}

	return result
}
