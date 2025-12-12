package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  //"strconv"
)

func compute_paths(manifolds [][]string, paths [][]int, row int, column int) [][]int {

	// Go through the 2D array and split where the character ahead is "^"
	for i := row; i < len(manifolds) - 1; i++ {
		for j := 0; j < len(manifolds[i]); j++ {
			if manifolds[i][j] == "|" {
				left_spot := j - 1
				right_spot := j + 1

				if manifolds[i+1][j] == "^" && left_spot >= 0 && right_spot < len(manifolds[i]) {
					manifolds[i+1][left_spot] = "|"
					manifolds[i+1][right_spot] = "|"
					paths[i+1][left_spot] = paths[i][j] + paths[i+1][left_spot]
					paths[i+1][right_spot] = paths[i][j] + paths[i+1][right_spot]
				} else {
					manifolds[i+1][j] = "|"
					paths[i+1][j] += paths[i][j]
				}
			}
		}
	}

	return paths 
}

func main() {
        // Open input file and check for error
        file, err := os.Open("./input")
        if err != nil {
                panic(err)
        }

        var tachyon_manifolds [][]string
	num_rows := 0
	column_length := 0

        // Read the file's line by line and make into 2d array
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                current_line := string(line)

		values := strings.Split(current_line, "")

		num_rows++
		column_length = len(values)
		tachyon_manifolds = append(tachyon_manifolds, values)
        }

	paths := make([][]int, num_rows)

	for i := 0; i < num_rows; i++ {
		paths[i] = make([]int, column_length)
	}

	start_row := 0
	start_column := 0

	for i := 0; i < len(tachyon_manifolds[start_row]); i++ {
		if tachyon_manifolds[start_row][i] == "S" {
			start_column = i
		}
	}

	tachyon_manifolds[start_row][start_column] = "|"
	paths[start_row][start_column] = 1

	paths = compute_paths(tachyon_manifolds, paths, start_row, start_column)

	// We total up the counts in the last row to get the total number of paths
	total_paths := 0
	last_row := paths[len(paths) - 1]

	for i := 0; i < len(last_row); i++ {
		total_paths += last_row[i]
	}

	fmt.Println(total_paths)
}

