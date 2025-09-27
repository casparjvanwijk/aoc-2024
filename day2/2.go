package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./2-input.txt")
	scanner := bufio.NewScanner(file)

	safeCount := 0
	for scanner.Scan() {
		// Parse report.
		spl := strings.Split(scanner.Text(), " ")
		levels := make([]int, len(spl))
		for i, str := range spl {
			levels[i], _ = strconv.Atoi(str)
		}

		safe, badIdx := isReportSafe(levels)
		if safe {
			safeCount++
			continue
		}

		// Try without first bad level.
		dampened := removeIndex(levels, badIdx)
		if safe, _ = isReportSafe(dampened); safe {
			safeCount++
			continue
		}

		// Try without level before first bad level.
		dampened = removeIndex(levels, badIdx-1)
		if safe, _ = isReportSafe(dampened); safe {
			safeCount++
			continue
		}

		// If possible, try without level before that.
		if badIdx < 2 {
			continue
		}
		dampened = removeIndex(levels, badIdx-2)
		if safe, _ = isReportSafe(dampened); safe {
			safeCount++
			continue
		}
	}
	fmt.Println(safeCount)
}

func isReportSafe(levels []int) (safe bool, badIdx int) {
	direction := "asc"
	if levels[1]-levels[0] < 0 {
		direction = "desc"
	}

	for i := range levels {
		// First level is always good.
		if i == 0 {
			continue
		}

		// Check with previous level.
		d := levels[i] - levels[i-1]
		if (d < 0 && direction == "asc") ||
			(d > 0 && direction == "desc") ||
			abs(d) < 1 || abs(d) > 3 {
			return false, i
		}
	}
	return true, 0
}

func removeIndex(s []int, idx int) []int {
	res := append([]int{}, s[:idx]...)
	return append(res, s[idx+1:]...)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
