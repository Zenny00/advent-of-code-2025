package main

import (
  "os"
  "bufio"
  "strconv"
  "fmt"
)

func find_max_joltage(bank string) int {
	max_index := 0
	max_value, err := strconv.Atoi(string(bank[max_index]))

	if err != nil {
		panic("Failed to convert!")
	}

	/**
	 * Go through and find the max value, this is our starting point
	 * Special case, if our max is at the end, we have no other digit to grab so we need to find the next largest
	 * To achieve this, we can just exclude the last index of the string in our search
	 */
	for i := 0; i < len(bank) - 1; i += 1 {
		current, err := strconv.Atoi(string(bank[i]))

		if err != nil {
			panic("Failed to convert!")
		}

		if current > max_value {
			max_value = current
			max_index = i
		}
	}


	// We then start at our max index to find the next largest value
	second_max_index := max_index + 1
	second_max_value, err := strconv.Atoi(string(bank[second_max_index]))

	if err != nil {
		panic("Failed to convert!")
	}

	for i := max_index + 1; i < len(bank); i += 1 {
		current, err := strconv.Atoi(string(bank[i]))

		if err != nil {
			panic("Failed to convert!")
		}

		if current > second_max_value {
			second_max_value = current
			second_max_index = i
		}
	}

	max_joltage_string := fmt.Sprintf("%d%d", max_value, second_max_value)
	max_joltage, err := strconv.Atoi(max_joltage_string)

	if err != nil {
		panic("Could not convert!")
	}

	return max_joltage
}

func main() {
	// Open input file and check for error
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	// Read the file line by line
	total_joltage := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := string(scanner.Text())
		max_joltage := find_max_joltage(line)
		total_joltage += max_joltage
	}

	fmt.Println(fmt.Sprintf("The total max joltage is: %d", total_joltage))
}
	
