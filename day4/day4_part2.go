package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func count_adjacent(forklift_room [][]string, num_rows int, row_length int, row int, column int) int {
	// Check adjacent cells from row - 1 to row + 1 and column - 1 to column + 1
	row_start := row - 1
	row_end := row + 1

	column_start := column - 1
	column_end := column + 1

	num_adjacent := 0

	for row_index := row_start; row_index <= row_end; row_index += 1 {

		// If outside bounds, continue, there are not rolls outside the array
		if row_index < 0 || row_index > num_rows - 1 {
			continue
		}

		for column_index := column_start; column_index <= column_end; column_index += 1 {
			if column_index < 0 || column_index > row_length - 1 || (row_index == row && column_index == column) {
				continue
			}

			// Otherwise, if the character at the index is not ".", there is a roll there
			if forklift_room[row_index][column_index] != "." {
				num_adjacent += 1
			}
		}
	}

	return num_adjacent
}

func get_accessible_rolls(forklift_room [][]string, num_rows int, row_length int) int {
	num_accessible := 0

	for i, row := range forklift_room {
		for j, value := range row {
			// Check any roll
			if value != "." {
				count := count_adjacent(forklift_room, num_rows, row_length, i, j)

				if count < 4 {
					num_accessible += 1
					forklift_room[i][j] = "."
				}
			}
		}

	}

	return num_accessible
}

func main() {
        // Open input file and check for error
        file, err := os.Open("./input")
        if err != nil {
                panic(err)
        }


	var forklift_room [][]string
	num_rows := 0
	row_length := 0

        // Read the file's line by line and make into 2d array
	scanner := bufio.NewScanner(file)
        for scanner.Scan() {
		line := scanner.Text()
		current_line := string(line)
		row_length = len(current_line)

		var row []string
		values := strings.Split(current_line, "")
		
		for _, value := range values {
			row = append(row, value)
		}

		forklift_room = append(forklift_room, row)
		num_rows += 1
	}

	total := 0

	amount := get_accessible_rolls(forklift_room, num_rows, row_length)
	for amount > 0 {
		total += amount

		amount = get_accessible_rolls(forklift_room, num_rows, row_length)
	}

	fmt.Println(total)
}
