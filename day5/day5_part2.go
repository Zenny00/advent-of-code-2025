package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func merge_ranges(first_range []int, second_range []int) int {
	return 0
}

func main() {
        // Open input file and check for error
        file, err := os.Open("./input")
        if err != nil {
                panic(err)
        }

	// fresh_ingredients := make(map[string]bool)
	var ranges [][]int
	//num_fresh := 0

        // Read the file's line by line and make into 2d array
        scanner := bufio.NewScanner(file)
	ingredients_found := false
        for scanner.Scan() {
                line := scanner.Text()
                current_line := string(line)

		// Add the ranges to a hashmap
		if current_line != "" && !ingredients_found {
			fresh_range := strings.Split(current_line, "-")
			range_start, err := strconv.Atoi(fresh_range[0])

			if err != nil {
				panic(fmt.Sprintf("Failed to convert start: %s", err))
			}

			range_end, err := strconv.Atoi(fresh_range[1])

			if err != nil {
				panic(fmt.Sprintf("Failed to convert end: %s", err))
			}

			range_array := []int{range_start, range_end}
			ranges = append(ranges, range_array)
		} else {
			break
		}
	}

	// Sort the array using the start of the range
	for i := 0; i < len(ranges) - 1; i++ {
		swapped := false
		for j := 0; j < len(ranges) - i - 1; j++ {
			if ranges[j][0] > ranges[j + 1][0] {
				tmp := ranges[j]
				ranges[j] = ranges[j + 1]
				ranges[j + 1] = tmp

				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	// Go through and merge the arrays
	var new_ranges [][]int
	index := 0
	start := ranges[index][0]
	end := ranges[index][1]

	index++

	for index < len(ranges) {
		if end >= ranges[index][0] {
			if ranges[index][1] >= end {
				end = ranges[index][1]
			}
		} else {
			new_range := []int{start, end}
			new_ranges = append(new_ranges, new_range)
			start = ranges[index][0]
			end = ranges[index][1]
		}

		index++
	}

	new_range := []int{start, end}
	new_ranges = append(new_ranges, new_range)

	total := 0
	for i := 0; i < len(new_ranges); i++ {
		range_start := new_ranges[i][0]
		range_end := new_ranges[i][1]

		total += range_end - range_start + 1
	}

	fmt.Println(total)
}

