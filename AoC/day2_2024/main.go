package main

import (
	"fmt"

	fileUtils "github.com/clombo/Aoc/utils/fileUtils"
	mathUtils "github.com/clombo/Aoc/utils/mathUtils"
)

func main() {

	// This is the main entry point for the Day 2 challenge of 2024.
	// You can implement your solution here.

	data, err := fileUtils.ReadFileLines("reportLevels.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	count := 0
	for _, line := range data {
		isSafe := isSafeWithDampening(line)
		if isSafe {
			count++
		}
		fmt.Printf("Report %d is safe: %v\n", line, isSafe)
	}
	fmt.Printf("Total safe reports: %d\n", count)
}

func isValidLevels(levels []int) bool {

	if len(levels) < 2 {
		return false // Need at least 2 values to compare
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// Rule 2: step size must be between 1 and 3
		if mathUtils.Abs(diff) < 1 || mathUtils.Abs(diff) > 3 {
			return false
		}

		// Track if it's still all increasing
		if diff <= 0 {
			increasing = false
		}

		// Track if it's still all decreasing
		if diff >= 0 {
			decreasing = false
		}
	}

	// Rule 1: Must be all increasing OR all decreasing
	return increasing || decreasing
}

func isSafeWithDampening(levels []int) bool {

	// Base case if already valid
	if isValidLevels(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		//Create a copy of the level slice up to the current index
		modified := append([]int{}, levels[:i]...)
		// Append the rest of the levels after the current index
		modified = append(modified, levels[i+1:]...)
		if isValidLevels(modified) {
			return true
		}
	}

	return false
}
