package main

import (
  "os"
  "bufio"
  "strconv"
  "fmt"
)

func find_max_joltage(bank string) int {
	num_digits := 12

	max_index := -1
	max_joltage := ""
	for num_digits > 0 {
		max_index += 1
		max_value, err := strconv.Atoi(string(bank[max_index]))

		if err != nil {
			panic("Failed to convert!")
		}

		for i := max_index; i < len(bank) - num_digits + 1; i++ {
			current_value := string(bank[i])
			current_value_int, err := strconv.Atoi(current_value)

			if err != nil {
				panic("Failed to convert!")
			}

			if current_value_int > max_value {
				max_value = current_value_int
				max_index = i
			}
		}

		max_joltage += strconv.Itoa(max_value)
		num_digits -= 1
	}
	
	fmt.Println(max_joltage)

	result, err := strconv.Atoi(max_joltage)

	if err != nil {
		panic("Failed to convert!")
	}

	return result
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
	
