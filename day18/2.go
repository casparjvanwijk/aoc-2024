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

	bytes := []co{}

	// Parse input.
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		bytes = append(bytes, co{x, y})
	}

	// Check paths.
	for _, b := range bytes {
		corrupted[b.y][b.x] = true
		if !pathPossible(corrupted) {
			fmt.Printf("%v,%v\n", b.x, b.y)
			return
		}
	}
}

func pathPossible(corrupted [][]bool) bool {
	// BFS shortest path.
	endPos := co{width - 1, height - 1}
	visited := make(map[co]struct{})
	queue := []co{{0, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if _, ok := visited[cur]; ok {
			continue
		}
		visited[cur] = struct{}{}

		if cur == endPos {
			return true
		}

		for _, d := range directions {
			next := co{cur.x + d.x, cur.y + d.y}
			// Check if next is on map.
			if next.x < 0 || next.x > width-1 || next.y < 0 || next.y > height-1 {
				continue
			}
			// Check if next is not corrupted.
			if corrupted[next.y][next.x] {
				continue
			}
			queue = append(queue, next)
		}
	}
	return false
}
