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

		// If the string is not even, then it does not meet the criteria for being invalid
		if len(id) % 2 == 0 {
			mid := len(id) / 2
			
			// Check if the first and second half of the string match
			if id[:mid] == id[mid:] {
				invalid_sum += i
			}
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

