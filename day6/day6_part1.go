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

        var numbers [][]string
	var operators[]string
	rows := 0

        // Read the file's line by line and make into 2d array
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                current_line := string(line)

                values := strings.Fields(current_line)

		if values[0] != "*" && values[0] != "+" {
			numbers = append(numbers, values)
			rows++
		} else {
			operators = values
		}
        }

	grand_total := 0

	for i := 0; i < len(operators); i++ {
		local_total := 0
		operator := operators[i]


		for j := 0; j < rows; j++ {
			number, err := strconv.Atoi(numbers[j][i])

			if err != nil {
				panic("Failed to convert a number!")
			}

			if operator == "*" {
				if local_total == 0 {
					local_total = 1
				}
				local_total *= number
			} else {
				local_total += number
			}
		}
		grand_total += local_total
	}

	fmt.Println(grand_total)
}

