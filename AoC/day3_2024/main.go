package main

import (
	"fmt"
	"strconv"

	"github.com/clombo/Aoc/utils/fileUtils"
	"github.com/clombo/Aoc/utils/mathUtils"
	"github.com/clombo/Aoc/utils/stringUtils"
)

func main() {

	//Read input from file
	content, err := fileUtils.ReadFileContents("corruptMemory.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	matches, totalMatches := stringUtils.FindAllByRegex(content, `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	fmt.Printf("Total matches found: %d\n", totalMatches)

	mulList := []int{}
	enabled := true

	for _, match := range matches {

		total := 0

		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			x, err1 := strconv.Atoi(match[1])
			y, err2 := strconv.Atoi(match[2])

			if err1 != nil || err2 != nil {
				fmt.Println("Error converting string to int:", err1, err2)
				continue
			}

			total = x * y
		}

		mulList = append(mulList, total)

	}

	sum := mathUtils.SumOfInts(mulList)

	fmt.Println("Total sum of multiplications:", sum)
}
