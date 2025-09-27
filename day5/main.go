package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// rule indicates that left must come before right.
type rule struct {
	left  string
	right string
}

func main() {
	file, _ := os.Open("./5-input.txt")
	scanner := bufio.NewScanner(file)

	// Process rules.
	hasRule := map[rule]bool{}
	for scanner.Scan() {
		// Scan until the first blank line.
		if scanner.Text() == "" {
			break
		}

		spl := strings.Split(scanner.Text(), "|")
		hasRule[rule{
			left:  spl[0],
			right: spl[1],
		}] = true
	}

	isUpdateCorrect := func(pages []string) bool {
		// Check whether rules exist for all adjacent pages.
		// NB: this assumes that rules exist for all combinations of pages,
		// which is the case in the puzzle input.
		for i := 0; i < len(pages)-1; i++ {
			if !hasRule[rule{
				left:  pages[i],
				right: pages[i+1],
			}] {
				return false
			}
		}
		return true
	}

	orderUpdate := func(pages []string) []string {
		result := make([]string, 0, len(pages))
	pageLoop:
		for _, a := range pages {
			for i, b := range result {
				// Insert on finding a matching rule.
				if hasRule[rule{
					left:  a,
					right: b,
				}] {
					newResult := append([]string{}, result[:i]...)
					newResult = append(newResult, a)
					result = append(newResult, result[i:]...)
					continue pageLoop
				}
			}
			// Otherwise, append.
			result = append(result, a)
		}
		return result
	}

	// Process updates.
	sumCorrect, sumIncorrectOrdered := 0, 0
	for scanner.Scan() {
		pages := strings.Split(scanner.Text(), ",")
		if isUpdateCorrect(pages) {
			middle, _ := strconv.Atoi(pages[len(pages)/2])
			sumCorrect += middle
			continue
		}

		ordered := orderUpdate(pages)
		middle, _ := strconv.Atoi(ordered[len(ordered)/2])
		sumIncorrectOrdered += middle
	}
	fmt.Println("Part One:", sumCorrect)
	fmt.Println("Part Two:", sumIncorrectOrdered)
}
