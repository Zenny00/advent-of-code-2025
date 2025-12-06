package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func main() {
        // Open input file and check for error
        file, err := os.Open("./input")
        if err != nil {
                panic(err)
        }

	// fresh_ingredients := make(map[string]bool)
	var ranges [][]int
	var ingredients []int
	num_fresh := 0

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
			ingredients_found = true

			if current_line == "" {
				continue
			}

			ingredient, err := strconv.Atoi(current_line)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert ingredient: %s", err))
			}


			// Add ingredient to list
			ingredients = append(ingredients, ingredient)
		}
        }

	for _, ingredient := range ingredients {
		for _, fresh_range := range ranges {
			range_start := fresh_range[0]
			range_end := fresh_range[1]
			fmt.Println(ingredient, range_start, range_end)
			if ingredient >= range_start && ingredient <= range_end {
				num_fresh += 1
				break
			}
		}
	}

	fmt.Println(num_fresh)
}

