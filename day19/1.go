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

	memo := make(map[string]bool)
	var isPossible func(design string) bool
	isPossible = func(design string) bool {
		if res, ok := memo[design]; ok {
			return res
		}

		// Base case.
		if design == "" {
			return true
		}

		// Try all patterns recursively.
		for _, p := range patterns {
			if strings.HasPrefix(design, p) {
				if isPossible(design[len(p):]) {
					memo[design] = true
					return true
				}
			}
		}
		memo[design] = false
		return false
	}

	// Check possible designs.
	possible := 0
	for _, d := range designs {
		if isPossible(d) {
			possible++
		}
	}
	fmt.Println(possible)
}
