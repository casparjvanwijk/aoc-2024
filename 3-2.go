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
	
	// Find mul(a,b), don't() and do() matches.
	exp, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
	matches := exp.FindAllStringSubmatch(input, -1)

	sum := 0
	enabled := true
	for _, m := range matches {
		switch m[0] {
		case "don't()":
			enabled = false
		case "do()":
			enabled = true
		default:
			if !enabled {
				continue
			}
			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])
			sum += a * b
		}
	}
	fmt.Println(sum)
}
