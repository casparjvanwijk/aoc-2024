package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.ReadFile("./3-input.txt")
	input := string(file)

	// Find mul(a,b) matches.
	exp, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := exp.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		sum += a * b
	}
	fmt.Println(sum)
}
