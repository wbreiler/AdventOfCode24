package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file containing the location IDs
	file, err := os.Open("Day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the contents of the file into slices
	var list1, list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by spaces and convert each number to an integer
		nums := strings.Fields(line)
		if len(nums) != 2 {
			continue // Ensure there are two numbers in the line
		}

		// Convert to integers and append to respective lists
		num1, err1 := strconv.Atoi(nums[0])
		num2, err2 := strconv.Atoi(nums[1])
		if err1 != nil || err2 != nil {
			continue // Skip lines with invalid numbers
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	// Part 1: Calculate total distance
	// Sort both lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		totalDistance += abs(list1[i] - list2[i])
	}

	// Part 2: Calculate the similarity score
	// Count occurrences of each number in the right list
	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num]++
	}

	// Calculate the similarity score
	similarityScore := 0
	for _, num := range list1 {
		similarityScore += num * counts[num]
	}

	// Output both the total distance and the similarity score
	fmt.Println("Total distance:", totalDistance)
	fmt.Println("Similarity score:", similarityScore)
}

// Helper function to calculate the absolute difference
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}