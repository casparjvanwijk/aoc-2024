package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Parse input.
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")
	scanner.Scan() // Empty line.
	designs := []string{}
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}

	memo := make(map[string]int)
	var possibilities func(design string) int
	possibilities = func(design string) int {
		if res, ok := memo[design]; ok {
			return res
		}

		// Base case.
		if design == "" {
			return 1
		}

		// Try all patterns recursively.
		result := 0
		for _, p := range patterns {
			if strings.HasPrefix(design, p) {
				result += possibilities(design[len(p):])
			}
		}
		memo[design] = result
		return result
	}

	// Check design possibilities.
	result := 0
	for _, d := range designs {
		result += possibilities(d)
	}
	fmt.Println(result)
}
