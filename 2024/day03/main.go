package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	example, err := os.ReadFile("example.txt")
	hasExample := err == nil && len(example) > 0

	if hasExample {
		fmt.Println("Example Tests:")
		fmt.Println("Part 1 (Example):", solvePart1(string(example)))
		fmt.Println("Part 2 (Example):", solvePart2(string(example)))
		fmt.Println("──────────────────────────────")
	}

	fmt.Println("Actual Input:")
	fmt.Println("Part 1:", solvePart1(string(input)))
	fmt.Println("Part 2:", solvePart2(string(input)))
}
