package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Possible search directions: horizontal, vertical, diagonal, and reverse
var directions = [][2]int{
	{0, 1},   // right
	{1, 0},   // down
	{1, 1},   // diagonal down-right
	{-1, 1},  // diagonal up-right
	{0, -1},  // left
	{-1, 0},  // up
	{-1, -1}, // diagonal up-left
	{1, -1},  // diagonal down-left
}

func countXMAS(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if findXMAS(grid, r, c, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func findXMAS(grid [][]rune, startR, startC, deltaR, deltaC int) bool {
	target := []rune("XMAS")
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < len(target); i++ {
		r := startR + i*deltaR
		c := startC + i*deltaC

		// Check if we're still within grid bounds
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}

		// Check if current letter matches
		if grid[r][c] != target[i] {
			return false
		}
	}

	return true
}

func main() {
	// Read the input file
	content, err := ioutil.ReadFile("Day4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert input to grid of runes
	lines := strings.Fields(string(content))
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// Count XMAS occurrences
	result := countXMAS(grid)
	fmt.Printf("Number of XMAS occurrences: %d\n", result)
}