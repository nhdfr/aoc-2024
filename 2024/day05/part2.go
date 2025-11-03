package main

import (
	"strconv"
	"strings"
)

func solvePart2(input string) interface{} {
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
		if !isCorrect(update, rules) {
			sorted := topologicalSort(update, rules)
			result += sorted[len(sorted)/2]
		}
	}

	return result
}

func topologicalSort(update []int, rules map[int]map[int]bool) []int {
	pageSet := make(map[int]bool)
	for _, p := range update {
		pageSet[p] = true
	}

	// Build graph and indegree for pages in update
	graph := make(map[int][]int)
	indegree := make(map[int]int)
	for _, p := range update {
		indegree[p] = 0
	}

	for before, afters := range rules {
		if !pageSet[before] {
			continue
		}
		for after := range afters {
			if !pageSet[after] {
				continue
			}
			graph[before] = append(graph[before], after)
			indegree[after]++
		}
	}

	var queue []int
	for p, deg := range indegree {
		if deg == 0 {
			queue = append(queue, p)
		}
	}

	var result []int
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr)

		for _, neigh := range graph[curr] {
			indegree[neigh]--
			if indegree[neigh] == 0 {
				queue = append(queue, neigh)
			}
		}
	}

	return result
}
