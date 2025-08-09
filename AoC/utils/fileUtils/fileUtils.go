package fileUtils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileLines(filePath string) ([][]int, error) {
	// Implementation for reading lines from a file
	var data [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		// Split the line into integers
		var numbers []int
		for _, part := range strings.Fields(line) {
			value, err := strconv.Atoi(part)
			if err == nil {
				numbers = append(numbers, value)
			}
		}

		data = append(data, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func ReadFileColumns(filePath string, columns int) ([][]int, error) {
	// Implementation for reading specific columns from a file

	// Prepare a slice of slices to hold the data
	data := make([][]int, columns)

	// Open the file
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		// Split the line into parts
		parts := strings.Fields(line)

		if len(parts) != columns {
			fmt.Println("Invalid number of columns in line:", line)
			continue
		}

		for i := 0; i < columns; i++ {
			value, err := strconv.Atoi(parts[i])
			if err != nil {
				fmt.Println("Invalid value in line:", line)
				continue
			}
			data[i] = append(data[i], value)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
