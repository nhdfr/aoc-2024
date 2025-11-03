package main

import (
	"regexp"
	"strconv"
	"strings"
)

func solvePart1(input string) interface{} {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	result := 0

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			result += a * b
		}
	}

	return result
}
