package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var validClockwise = []string{"MMSS", "MSSM", "SMMS", "SSMM"}

func main() {
	file, _ := os.Open("4-input.txt")
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	colMax := len(lines[0]) - 1
	rowMax := len(lines) - 1
	count := 0
	for row, line := range lines {
		// Check surrounding letters, if within bounds, from every A.
		for col, char := range line {
			if string(char) != "A" {
				continue
			}

			if col == 0 || col >= colMax || row == 0 || row >= rowMax {
				continue
			}
			clockwise := string(lines[row-1][col+1]) +
				string(lines[row+1][col+1]) +
				string(lines[row+1][col-1]) +
				string(lines[row-1][col-1])
			if slices.Contains(validClockwise, clockwise) {
				count++
			}
		}
	}
	fmt.Println(count)
}
