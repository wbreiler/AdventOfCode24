package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file containing the reports
	file, err := os.Open("Day2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Parse the input into a list of reports
	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := parseLineToInts(line)
		reports = append(reports, levels)
	}

	// Count safe reports
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) || isSafeWithOneRemoval(report) {
			safeCount++
		}
	}

	// Print the number of safe reports
	fmt.Println("Number of safe reports:", safeCount)
}

// Parse a line of space-separated integers into a slice of ints
func parseLineToInts(line string) []int {
	parts := strings.Fields(line)
	levels := make([]int, len(parts))
	for i, part := range parts {
		levels[i], _ = strconv.Atoi(part)
	}
	return levels
}

// Check if a report is safe
func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	// Determine if the sequence is increasing or decreasing
	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check if the difference is outside the range [1, 3]
		if diff < -3 || diff > 3 {
			return false
		}

		// Check for consistent direction (all increasing or all decreasing)
		if isIncreasing && diff <= 0 {
			return false
		}
		if !isIncreasing && diff >= 0 {
			return false
		}
	}

	return true
}

// Check if a report can be safe by removing one level
func isSafeWithOneRemoval(report []int) bool {
	for i := range report {
		// Create a new report excluding the current level
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)

		// Check if the modified report is safe
		if isSafe(modifiedReport) {
			return true
		}
	}
	return false
}