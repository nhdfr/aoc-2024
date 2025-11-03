package main

import (
	"strconv"
	"strings"
)

func solvePart1(input string) interface{} {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	rules := make(map[int]map[int]bool)
	var updates [][]int
	parsingRules := true

	for _, line := range lines {
		if line == "" {
			parsingRules = false
			continue
		}
		if parsingRules {
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])
			if rules[before] == nil {
				rules[before] = make(map[int]bool)
			}
			rules[before][after] = true
		} else {
			parts := strings.Split(line, ",")
			var update []int
			for _, p := range parts {
				num, _ := strconv.Atoi(p)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	result := 0
	for _, update := range updates {
		if isCorrect(update, rules) {
			result += update[len(update)/2]
		}
	}

	return result
}

func isCorrect(update []int, rules map[int]map[int]bool) bool {
	pos := make(map[int]int)
	for i, num := range update {
		pos[num] = i
	}
	for before, afters := range rules {
		if _, ok := pos[before]; !ok {
			continue
		}
		for after := range afters {
			if _, ok := pos[after]; !ok {
				continue
			}
			if pos[before] > pos[after] {
				return false
			}
		}
	}
	return true
}
