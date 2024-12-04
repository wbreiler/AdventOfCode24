package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func parseAndSumMultiplications(memoryString string) int {
	// Regular expression to find valid mul() instructions
	mulPattern := regexp.MustCompile(`mul\s*\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)`)
	
	// Find all matches
	matches := mulPattern.FindAllStringSubmatch(memoryString, -1)
	
	// Calculate total sum
	totalSum := 0
	for _, match := range matches {
		if len(match) == 3 {
			// Convert first and second captured groups to integers
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			totalSum += x * y
		}
	}
	
	return totalSum
}

func main() {
	// Read the input file
	content, err := ioutil.ReadFile("Day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	// Convert byte slice to string and trim whitespace
	memoryString := string(content)
	
	// Calculate and print the result
	result := parseAndSumMultiplications(memoryString)
	fmt.Printf("Sum of all multiplication results: %d\n", result)
}