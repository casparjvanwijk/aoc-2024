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

		if isReportSafe(levels) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}

func isReportSafe(levels []int) bool {
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
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
