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
	var operators []string
	rows := 0

        // Read the file's line by line and make into 2d array
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                current_line := string(line)


		values := strings.Split(current_line, "")
		if values[0] != "*" && values[0] != "+" {
			numbers = append(numbers, values)
			rows++
		} else {
			values := strings.Fields(current_line)
			operators = values
		}
        }

	// Parse out the numbers
	var new_numbers [][]string
	var operands []string
	for i := len(numbers[0]) - 1; i >= 0 ; i-- {
		number := ""
		for j := 0; j < rows; j++ {
			if numbers[j][i] != " " {
				number += numbers[j][i]
			}
		}

		if number != "" {
			operands = append(operands, number)
		} else {
			if len(operands) > 0 {
				operands_copy := make([]string, len(operands))
				copy(operands_copy, operands)

				new_numbers = append(new_numbers, operands_copy)
				operands = []string{}
			}
		}
	}


	new_numbers = append(new_numbers, operands)

	grand_total := 0

	operand_index := 0
	for i := len(operators) - 1; i >= 0; i-- {
		local_total := 0
		operator := operators[i]

		for j := 0; j < len(new_numbers[operand_index]); j++ {
			number, err := strconv.Atoi(new_numbers[operand_index][j])

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
		operand_index++
	}

	fmt.Println(grand_total)
}

