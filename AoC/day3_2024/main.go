package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/clombo/Aoc/utils/mathUtils"
)

func main() {

	//Regex to match the pattern "mul(x,y)" where x and y are integers
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	//Read input from file
	data, err := os.ReadFile("corruptMemory.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//Convert data to string
	content := string(data)

	//Return all matches of the regex in the content
	//Regex package n parameter: < 0 means to return all matches, > 0 means return at most n matches, = 0 means return no matches
	matches := re.FindAllStringSubmatch(content, -1)
	totalMatches := len(matches)

	fmt.Printf("Total matches found: %d\n", totalMatches)

	mulList := []int{}

	for _, match := range matches {

		x, err1 := strconv.Atoi(match[1])
		y, err2 := strconv.Atoi(match[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting string to int:", err1, err2)
			continue
		}

		total := x * y
		mulList = append(mulList, total)
	}

	sum := mathUtils.SumOfInts(mulList)

	fmt.Println("Total sum of multiplications:", sum)

}
