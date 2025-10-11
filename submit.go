package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
)

type Meta struct {
	Part1Submitted bool `json:"part1Submitted"`
	Part2Submitted bool `json:"part2Submitted"`
}

type SubmissionResult struct {
	Success     bool
	Message     string
	RateLimit   bool
	WaitTime    time.Duration
	AlreadyDone bool
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run submit.go <year> <day>")
		os.Exit(1)
	}

	year, _ := strconv.Atoi(os.Args[1])
	day, _ := strconv.Atoi(os.Args[2])

	// Load environment variables
	godotenv.Load()

	dayPadded := fmt.Sprintf("%02d", day)
	dayDir := fmt.Sprintf("%d/day%s", year, dayPadded)

	fmt.Printf("%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", colorPurple, colorReset)
	fmt.Printf("%sAOC %d Day %d - Submission Workflow%s\n", colorPurple, year, day, colorReset)
	fmt.Printf("%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", colorPurple, colorReset)

	// Check if day directory exists
	if _, err := os.Stat(dayDir); os.IsNotExist(err) {
		fmt.Printf("%sDay %d not found. Run './aoc.sh %d' first to set it up.%s\n", colorRed, day, day, colorReset)
		os.Exit(1)
	}

	// Change to day directory
	os.Chdir(dayDir)

	// Load meta.json
	meta := loadMeta()

	if !meta.Part1Submitted {
		runPart(1, year, day, &meta)
	}

	// Reload meta after potential part 1 submission
	meta = loadMeta()

	if meta.Part1Submitted && !meta.Part2Submitted {
		fmt.Printf("\n%s──────────────────────────────────────────────────%s\n", colorCyan, colorReset)
		runPart(2, year, day, &meta)
	}

	// Final status
	meta = loadMeta()
	if meta.Part1Submitted && meta.Part2Submitted {
		fmt.Printf("\n%s Day %d Complete! Both parts submitted successfully!%s\n", colorGreen, day, colorReset)
	}
}

func runPart(part, year, day int, meta *Meta) {
	fmt.Printf("\n%s Running Part %d...%s\n", colorBlue, part, colorReset)

	// Run the solution
	cmd := exec.Command("go", "run", ".")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s Error running solution: %v%s\n", colorRed, err, colorReset)
		return
	}

	// Parse output for the specific part
	lines := strings.Split(string(output), "\n")
	var result string
	
	for _, line := range lines {
		if strings.Contains(line, fmt.Sprintf("Part %d:", part)) {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				result = strings.TrimSpace(parts[1])
				break
			}
		}
	}

	if result == "" {
		fmt.Printf("%s Could not parse Part %d result from output%s\n", colorRed, part, colorReset)
		fmt.Printf("%sOutput was:%s\n%s\n", colorYellow, colorReset, string(output))
		return
	}

	fmt.Printf("%s Part %d Result: %s%s\n", colorCyan, part, result, colorReset)

	// Ask for submission
	fmt.Printf("\n%s Submit Part %d answer '%s'? (y/N): %s", colorYellow, part, result, colorReset)
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response == "y" || response == "yes" {
		fmt.Printf("%s═══════════════════════════════════════%s\n", colorPurple, colorReset)
		fmt.Printf("%s Submitting Part %d answer: %s%s\n", colorBlue, part, result, colorReset)
		fmt.Printf("%s Waiting for AOC response...%s", colorYellow, colorReset)
		
		submissionResult := submitAnswer(year, day, part, result)
		fmt.Printf(" ✓\n")
		
		if submissionResult.Success {
			fmt.Printf("%s CORRECT! Part %d solved successfully!%s\n", colorGreen, part, colorReset)
			if part == 1 {
				fmt.Printf("%s You earned one gold star!%s\n", colorYellow, colorReset)
			} else {
				fmt.Printf("%s You earned two gold stars! Day complete!%s\n", colorYellow, colorReset)
			}
			updateSubmissionStatus(part, true)
		} else if submissionResult.AlreadyDone {
			fmt.Printf("%s Part %d already completed! Marking as submitted.%s\n", colorGreen, part, colorReset)
			updateSubmissionStatus(part, true)
		} else if submissionResult.RateLimit {
			fmt.Printf("%s RATE LIMITED: You submitted too recently.%s\n", colorYellow, colorReset)
			fmt.Printf("%s Please wait %v before submitting again.%s\n", colorCyan, submissionResult.WaitTime, colorReset)
		} else {
			fmt.Printf("%s INCORRECT answer for Part %d.%s\n", colorRed, part, colorReset)
			if submissionResult.Message != "" {
				fmt.Printf("%s Hint: %s%s\n", colorCyan, submissionResult.Message, colorReset)
			}
			
			// Ask if they want to retry
			fmt.Printf("\n%s Want to retry with a different answer? (y/N): %s", colorYellow, colorReset)
			retryResponse, _ := reader.ReadString('\n')
			retryResponse = strings.TrimSpace(strings.ToLower(retryResponse))
			
			if retryResponse == "y" || retryResponse == "yes" {
				fmt.Printf("%sEnter new answer: %s", colorCyan, colorReset)
				newAnswer, _ := reader.ReadString('\n')
				newAnswer = strings.TrimSpace(newAnswer)
				
				if newAnswer != "" {
					fmt.Printf("%s Submitting retry answer: %s%s\n", colorBlue, newAnswer, colorReset)
					retryResult := submitAnswer(year, day, part, newAnswer)
					
					if retryResult.Success {
						fmt.Printf("%s CORRECT! Part %d solved successfully!%s\n", colorGreen, part, colorReset)
						updateSubmissionStatus(part, true)
					} else {
						fmt.Printf("%s Still incorrect: %s%s\n", colorRed, retryResult.Message, colorReset)
					}
				}
			}
		}
		fmt.Printf("%s═══════════════════════════════════════%s\n", colorPurple, colorReset)
	} else {
		fmt.Printf("%s  Skipped submission for Part %d%s\n", colorCyan, part, colorReset)
	}
}

func submitAnswer(year, day, part int, answer string) SubmissionResult {
	sessionToken := os.Getenv("AOC_SESSION")
	if sessionToken == "" {
		return SubmissionResult{Success: false, Message: "No AOC_SESSION token found in .env"}
	}

	submitURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)

	data := url.Values{}
	data.Set("level", strconv.Itoa(part))
	data.Set("answer", answer)

	req, err := http.NewRequest("POST", submitURL, strings.NewReader(data.Encode()))
	if err != nil {
		return SubmissionResult{Success: false, Message: fmt.Sprintf("Request failed: %v", err)}
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return SubmissionResult{Success: false, Message: fmt.Sprintf("Request failed: %v", err)}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SubmissionResult{Success: false, Message: fmt.Sprintf("Failed to read response: %v", err)}
	}

	responseText := string(body)

	// Parse response
	if strings.Contains(responseText, "That's the right answer") {
		return SubmissionResult{Success: true}
	} else if strings.Contains(responseText, "You don't seem to be solving the right level") || 
	          strings.Contains(responseText, "Did you already complete it") {
		return SubmissionResult{AlreadyDone: true}
	} else if strings.Contains(responseText, "You gave an answer too recently") {
		waitTime := time.Minute // Default
		if strings.Contains(responseText, "one minute") {
			waitTime = time.Minute
		} else if strings.Contains(responseText, "5 minutes") {
			waitTime = 5 * time.Minute
		}
		return SubmissionResult{RateLimit: true, WaitTime: waitTime}
	} else if strings.Contains(responseText, "That's not the right answer") {
		message := "Wrong answer"
		if strings.Contains(responseText, "too high") {
			message = "Your answer is too high"
		} else if strings.Contains(responseText, "too low") {
			message = "Your answer is too low"
		}
		return SubmissionResult{Success: false, Message: message}
	}

	return SubmissionResult{Success: false, Message: "Unknown response from AOC"}
}

func loadMeta() Meta {
	var meta Meta
	data, err := os.ReadFile("meta.json")
	if err != nil {
		// Return default if file doesn't exist
		return Meta{Part1Submitted: false, Part2Submitted: false}
	}
	
	json.Unmarshal(data, &meta)
	return meta
}

func updateSubmissionStatus(part int, submitted bool) {
	meta := loadMeta()
	
	if part == 1 {
		meta.Part1Submitted = submitted
	} else {
		meta.Part2Submitted = submitted
	}

	data, _ := json.MarshalIndent(meta, "", "  ")
	os.WriteFile("meta.json", data, 0644)
}
