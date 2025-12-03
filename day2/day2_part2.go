package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func find_invalids(start string, end string) int {
	invalid_sum := 0
	
	// Convert strings to ints to allow iteration over range
	start_value, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}

	end_value, err := strconv.Atoi(end)
	if err != nil {
		panic(err)
	}

	for i := start_value; i <= end_value; i++ {
		// Convert i into a string so we can check if invalid
		id := strconv.Itoa(i)
	
		slice_index := 1
		id_length := len(id)
		is_invalid := true

		for slice_index <= id_length / 2 {
			is_invalid = true

			slice := id[0:slice_index]
			step := len(slice)

			for remaining_index := 0; remaining_index < len(id); remaining_index += step {
				
				// If we are over the length we know this is not invalid
				if remaining_index + step > len(id) {
					is_invalid = false
					break
				}

				// If the next slice does not match what we have, we know the id is not invalid
				if id[remaining_index:remaining_index+step] != slice {
					is_invalid = false
					break
				}
			}


			// If invalid, add to total and break early for efficiency and to avoid double counting
			if is_invalid {
				invalid_sum += i
				break
			}

			slice_index += 1
		}

	}

	return invalid_sum
}

func main() {
	// Open input file and check for error
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	// Read the file's single line
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	ranges := strings.Split(line, ",")

	// Store the total sum
	total_sum := 0

	// Go through each of the provided ranges
	for _, id_range := range ranges {
		range_boundaries := strings.Split(id_range, "-")
		total_sum += find_invalids(range_boundaries[0], range_boundaries[1])
	}

	fmt.Println(total_sum)
}

