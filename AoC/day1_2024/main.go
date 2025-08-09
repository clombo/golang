package main

import (
	"fmt"
	"math"
	"sort"

	fileUtils "github.com/clombo/Aoc/utils/fileUtils"
	"github.com/clombo/Aoc/utils/mathUtils"
)

func main() {
	// This is the main entry point for the Day 1 challenge of 2024.
	// You can implement your solution here.

	data, err := fileUtils.ReadFileColumns("locationId.txt", 2)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	list1 := data[0]
	list2 := data[1]

	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Create a new list to hold the differences
	diffList := createDifferenceList(list1, list2)

	total := mathUtils.SumOfInts(diffList)
	fmt.Println("Total sum of differences:", total)

	list2Occurrences := getValueOccurrences(list2)

	var simList []int

	for _, value := range list1 {
		if _, exists := list2Occurrences[value]; exists {
			occurence := list2Occurrences[value]
			simList = append(simList, value*occurence)
		}
	}

	simScore := mathUtils.SumOfInts(simList)
	fmt.Println("Similarity score:", simScore)

	//fmt.Println("Differences:", diffList)
}

func createDifferenceList(list1, list2 []int) []int {
	// Create a new list to hold the differences
	var diffList []int

	for i := 0; i < len(list1) && i < len(list2); i++ {
		diff := int(math.Abs(float64(list1[i] - list2[i])))
		diffList = append(diffList, diff)
	}

	return diffList
}

func getValueOccurrences(list []int) map[int]int {
	// Create a map to hold the occurrences of each value
	valueOccurrences := make(map[int]int)

	for _, value := range list {
		valueOccurrences[value]++
	}

	return valueOccurrences
}
