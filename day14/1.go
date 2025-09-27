package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

const width, height = 101, 103
const seconds = 100

func main() {
	file, _ := os.Open("./14-input.txt")
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile(`p=(.+),(.+)\sv=(.+),(.+)`)

	topLeft, botLeft, topRight, botRight := 0, 0, 0, 0
	for scanner.Scan() {
		res := r.FindStringSubmatch(scanner.Text())
		xStart, _ := strconv.Atoi(res[1])
		yStart, _ := strconv.Atoi(res[2])
		vX, _ := strconv.Atoi(res[3])
		vY, _ := strconv.Atoi(res[4])

		x := mod(xStart + seconds * vX, width) 
		y := mod(yStart + seconds * vY, height) 

		switch {
		case x < width / 2 && y < height / 2:
			topLeft++
		case x < width / 2 && y > height / 2:
			botLeft++
		case x > width / 2 && y < height / 2:
			topRight++
		case x > width / 2 && y > height / 2:
			botRight++
		}
	}

	fmt.Println(topLeft * botLeft * topRight * botRight)
}

// mod calculates the Euclidian modulo (instead of the truncated modulo which is % in Go).
func mod(a, b int) int {
	return (a % b + b) % b
}
