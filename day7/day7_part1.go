package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  //"strconv"
)

func propagate(manifolds [][]string, start_row int, start_column int) int {
	// Go through the 2D array and split where the character ahead is "^"
	for i := start_row; i < len(manifolds) - 1; i++ {
		for j := 0; j < len(manifolds[i]); j++ {
			if manifolds[i][j] == "|" {
				left_spot := j - 1
				right_spot := j + 1

				if manifolds[i+1][j] == "^" && left_spot >= 0 && right_spot < len(manifolds[i]) {
					manifolds[i+1][left_spot] = "|"
					manifolds[i+1][right_spot] = "|"
					splits += 1
				} else {
					manifolds[i+1][j] = "|"
				}
			}
		}
	}

	fmt.Println(manifolds)

	return splits
}

func main() {
        // Open input file and check for error
        file, err := os.Open("./example")
        if err != nil {
                panic(err)
        }

        var tachyon_manifolds [][]string

        // Read the file's line by line and make into 2d array
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                current_line := string(line)

		values := strings.Split(current_line, "")

		tachyon_manifolds = append(tachyon_manifolds, values)
        }

	start_row := 0
	start_column := 0

	for i := 0; i < len(tachyon_manifolds[start_row]); i++ {
		if tachyon_manifolds[start_row][i] == "S" {
			start_column = i
		}
	}

	fmt.Println(tachyon_manifolds)
	fmt.Println(start_row, start_column)

	start_row = 1
	tachyon_manifolds[start_row][start_column] = "|"

	tachyon_manifold_splits := propagate(tachyon_manifolds, start_row, start_column)

	fmt.Println(tachyon_manifold_splits)
}

