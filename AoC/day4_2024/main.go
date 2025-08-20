package main

import (
	"fmt"
	"strings"

	"github.com/clombo/Aoc/utils/fileUtils"
)

// All 8 possible directions: up, up-right, right, down-right, down, down-left, left, up-left
var directions = [][2]int{
	{-1, 0},  // up
	{-1, 1},  // up-right
	{0, 1},   // right
	{1, 1},   // down-right
	{1, 0},   // down
	{1, -1},  // down-left
	{0, -1},  // left
	{-1, -1}, // up-left
}

func main() {
	//Read input file contents
	contents, err := fileUtils.ReadFileContents("wordSearch.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//Number of columns for the grid
	cols := 4

	//Build the grid from the contents
	grid := buildGrid(contents, cols)

	word := "XMAS"
	count := countWordOccurrences(grid, word)
	fmt.Printf("Word %q occurs %d times in the grid.\n", word, count)

	//This is not working as expected
	// Instead make a grid of runes with 4 rows and 4 columns and split string into runes
	// and count occurrences of the word "XMAS" in all 8 directions.
}

// buildGrid takes a long string and splits it into rows based on column count.
func buildGrid(longStr string, cols int) [][]rune {

	longStr = strings.ReplaceAll(longStr, " ", "")  // remove spaces
	longStr = strings.ReplaceAll(longStr, "\n", "") // remove newlines

	letters := []rune(longStr)
	rows := len(letters) / cols
	grid := make([][]rune, rows)

	for r := 0; r < rows; r++ {
		grid[r] = letters[r*cols : (r+1)*cols]
	}

	return grid
}

// countWordOccurrences searches the grid and counts total matches of the word in all 8 directions.
func countWordOccurrences(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordRunes := []rune(word)
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			count += searchFrom(grid, r, c, wordRunes)
		}
	}
	return count
}

// searchFrom checks all 8 directions starting from (r, c) and returns how many matches were found.
func searchFrom(grid [][]rune, r, c int, word []rune) int {
	matchCount := 0
	for _, dir := range directions {
		rr, cc := r, c
		found := true
		for i := 0; i < len(word); i++ {
			if rr < 0 || rr >= len(grid) || cc < 0 || cc >= len(grid[0]) || grid[rr][cc] != word[i] {
				found = false
				break
			}
			rr += dir[0]
			cc += dir[1]
		}
		if found {
			matchCount++
		}
	}
	return matchCount
}
