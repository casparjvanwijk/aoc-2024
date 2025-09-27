package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const width, height = 71, 71

type co struct {
	x, y int
}

type direction struct {
	x, y int
}

var directions = []direction{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func main() {
	// corrupted is a grid which tracks corrupted fields.
	corrupted := make([][]bool, height)
	for i := range corrupted {
		line := make([]bool, width)
		corrupted[i] = line
	}

	// Parse input.
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	for range 1024 {
		scanner.Scan()
		split := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		corrupted[y][x] = true
	}

	// BFS shortest path.
	endPos := co{width - 1, height - 1}
	visited := make(map[co]struct{})
	type queueItem struct {
		pos  co
		dist int
	}
	queue := []queueItem{{pos: co{0, 0}, dist: 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if _, ok := visited[cur.pos]; ok {
			continue
		}
		visited[cur.pos] = struct{}{}

		if cur.pos == endPos {
			fmt.Println("Distance:", cur.dist)
			return
		}

		for _, d := range directions {
			nextPos := co{cur.pos.x + d.x, cur.pos.y + d.y}
			// Check if next is on map.
			if nextPos.x < 0 || nextPos.x > width-1 || nextPos.y < 0 || nextPos.y > height-1 {
				continue
			}
			// Check if next is not corrupted.
			if corrupted[nextPos.y][nextPos.x] {
				continue
			}
			queue = append(queue, queueItem{
				pos:  nextPos,
				dist: cur.dist + 1,
			})
		}
	}

}
