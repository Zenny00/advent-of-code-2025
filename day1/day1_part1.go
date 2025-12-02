package main

import (
  "os"
  "bufio"
  "strconv"
  "fmt"
)

func main() {
	// Open input file and check for error
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	// Store the values 0 - 90 in the array, index is the same as the value
	var dial_values [100]int
	for i := range 99 {
		dial_values[i] = i
	}

	// Read the file line by line
	num_zeros := 0
	dial_location := 50
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Determine if index should go up or down
		prefix := string(line[0])
		direction := 1
		if prefix == "L" {
			direction *= -1
		}

		// Cast the remainder of the line to a number value
		magnitude, err := strconv.Atoi(line[1:])

		// Check if conversion to int was successful
		if err != nil {
			panic(err)
		}

		for magnitude > 0 {
			// Move in specified direction mod 99 (this allows us to loop around the "dial")
			dial_location = (dial_location + direction) % 100
			if dial_location < 0 {
				dial_location += 100
			}

			magnitude -= 1
		}

		if dial_location == 0 {
			num_zeros += 1
		}	
	}

	// Close file and panic if cannot close
	if err := file.Close(); err != nil {
		panic(err)
	}

	// Print the number of zeros
	fmt.Println(num_zeros)
}
