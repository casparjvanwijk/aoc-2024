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

	scoreSum := 0

	// BFS from each trailhead, looking for trail ends.
	for _, th := range trailheads {
		trailEnds := map[co]bool{}

		queue := []co{th}
		discovered := map[co]bool{th: true}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			if grid[cur] == 9 {
				trailEnds[cur] = true
				continue
			}

			neighbors := []co{
				{cur.x, cur.y - 1},
				{cur.x + 1, cur.y},
				{cur.x, cur.y + 1},
				{cur.x - 1, cur.y},
			}

			for _, n := range neighbors {
				if grid[n] == grid[cur]+1 && !discovered[n] {
					queue = append(queue, n)
					discovered[n] = true
				}
			}
		}
		scoreSum += len(trailEnds)
	}

	fmt.Println(scoreSum)
}
