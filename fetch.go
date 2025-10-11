package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run fetch.go <year> <day>")
		os.Exit(1)
	}

	year, _ := strconv.Atoi(os.Args[1])
	day, _ := strconv.Atoi(os.Args[2])

	// Load environment variables
	godotenv.Load()

	sessionToken := os.Getenv("AOC_SESSION")
	if sessionToken == "" {
		fmt.Println(" No AOC_SESSION token found in .env file")
		os.Exit(1)
	}

	// Fetch input
	inputURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	
	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		fmt.Printf(" Failed to create request: %v\n", err)
		os.Exit(1)
	}
	
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf(" Failed to fetch input: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf(" Failed to fetch input: HTTP %d\n", resp.StatusCode)
		if resp.StatusCode == 404 {
			fmt.Printf("Day %d might not be available yet\n", day)
		}
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf(" Failed to read response: %v\n", err)
		os.Exit(1)
	}
	
	// Save input
	err = os.WriteFile("input.txt", body, 0644)
	if err != nil {
		fmt.Printf(" Failed to save input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input fetched successfully for Day %d\n", day)
}
