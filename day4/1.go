package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		// Check for XMAS in every direction, if within bounds, from every X.
		for col, char := range line {
			if string(char) != "X" {
				continue
			}

			// Right
			if col <= colMax-3 &&
				string(lines[row][col+1]) == "M" &&
				string(lines[row][col+2]) == "A" &&
				string(lines[row][col+3]) == "S" {
				count++
			}

			// Left
			if col >= 3 &&
				string(lines[row][col-1]) == "M" &&
				string(lines[row][col-2]) == "A" &&
				string(lines[row][col-3]) == "S" {
				count++
			}

			// Down
			if row <= rowMax-3 &&
				string(lines[row+1][col]) == "M" &&
				string(lines[row+2][col]) == "A" &&
				string(lines[row+3][col]) == "S" {
				count++
			}

			// Up
			if row >= 3 &&
				string(lines[row-1][col]) == "M" &&
				string(lines[row-2][col]) == "A" &&
				string(lines[row-3][col]) == "S" {
				count++
			}

			// Right down
			if col <= colMax-3 && row <= rowMax-3 &&
				string(lines[row+1][col+1]) == "M" &&
				string(lines[row+2][col+2]) == "A" &&
				string(lines[row+3][col+3]) == "S" {
				count++
			}

			// Right up
			if col <= colMax-3 && row >= 3 &&
				string(lines[row-1][col+1]) == "M" &&
				string(lines[row-2][col+2]) == "A" &&
				string(lines[row-3][col+3]) == "S" {
				count++
			}

			// Left down
			if col >= 3 && row <= rowMax-3 &&
				string(lines[row+1][col-1]) == "M" &&
				string(lines[row+2][col-2]) == "A" &&
				string(lines[row+3][col-3]) == "S" {
				count++
			}

			// Left up
			if col >= 3 && row >= 3 &&
				string(lines[row-1][col-1]) == "M" &&
				string(lines[row-2][col-2]) == "A" &&
				string(lines[row-3][col-3]) == "S" {
				count++
			}
		}
	}
	fmt.Println(count)
}
