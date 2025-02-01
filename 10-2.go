package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type co struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("./10-input.txt")
	scanner := bufio.NewScanner(file)

	grid := map[co]int{}
	trailheads := []co{}
	for y := 0; scanner.Scan(); y++ {
		for x, r := range scanner.Text() {
			val, _ := strconv.Atoi(string(r))
			grid[co{x, y}] = val
			if val == 0 {
				trailheads = append(trailheads, co{x, y})
			}
		}
	}

	ratingSum := 0

	// DFS from each trailhead, looking for unique routes to trail ends.
	for _, th := range trailheads {
		ratingSum += pathsFrom(th, grid)
	}

	fmt.Println(ratingSum)
}

func pathsFrom(pos co, grid map[co]int) int {
	if grid[pos] == 9 {
		return 1
	}

	found := 0
	neighbors := []co{
		{pos.x, pos.y - 1},
		{pos.x + 1, pos.y},
		{pos.x, pos.y + 1},
		{pos.x - 1, pos.y},
	}
	for _, n := range neighbors {
		if grid[n] == grid[pos]+1 {
			found += pathsFrom(n, grid)
		}
	}
	return found
}
