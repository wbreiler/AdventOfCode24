package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func parseAndSumMultiplications(memoryString string) int {
	mulPattern := regexp.MustCompile(`mul\s*\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)`)
	doPattern := regexp.MustCompile(`do\s*\(\s*\)`)
	dontPattern := regexp.MustCompile(`don't\s*\(\s*\)`)
	
	totalSum := 0
	enabled := true
	
	// Find all do(), don't(), and mul instructions in the order they appear
	fullPattern := regexp.MustCompile(`(mul\s*\(\s*\d{1,3}\s*,\s*\d{1,3}\s*\)|do\s*\(\s*\)|don't\s*\(\s*\))`)
	matches := fullPattern.FindAllStringSubmatch(memoryString, -1)
	
	for _, match := range matches {
		instruction := match[0]
		
		switch {
		case mulPattern.MatchString(instruction):
			// Multiplication instruction
			if enabled {
				submatches := mulPattern.FindStringSubmatch(instruction)
				x, _ := strconv.Atoi(submatches[1])
				y, _ := strconv.Atoi(submatches[2])
				totalSum += x * y
			}
		case doPattern.MatchString(instruction):
			// Enable multiplications
			enabled = true
		case dontPattern.MatchString(instruction):
			// Disable multiplications
			enabled = false
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
	fmt.Printf("Sum of enabled multiplication results: %d\n", result)
}