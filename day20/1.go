package main

import (
	"bufio"
	"fmt"
	"os"
)

// minSaved is the least amount of time saved by cheating to count the cheat.
const minSaved = 100

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
	grid := [][]rune{}
	var startPos co

	// Parse input.
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := []rune{}
		for x, rune := range scanner.Text() {
			line = append(line, rune)
			if rune == 'S' {
				startPos = co{x, y}
			}
		}
		grid = append(grid, line)
	}

	// Find route without cheats.
	var endDist int
	route := []co{}
	pos := startPos
	visited := make(map[co]struct{})
	for i := 0; ; i++ {
		route = append(route, pos)
		visited[pos] = struct{}{}
		if grid[pos.y][pos.x] == 'E' {
			endDist = i
			break
		}
		for _, d := range directions {
			next := co{pos.x + d.x, pos.y + d.y}
			if _, ok := visited[next]; ok {
				continue
			}
			if grid[next.y][next.x] == '#' {
				continue
			}
			pos = next
			break
		}
	}

	// Map distances for all fields.
	distToEnd := make(map[co]int)
	for i, pos := range route {
		distToEnd[pos] = endDist - i
	}

	// Look for cheats from every field on route.
	cheatCount := 0
	for _, pos := range route {
		for _, d := range directions {
			next := co{pos.x + d.x, pos.y + d.y}
			afterNext := co{pos.x + 2*d.x, pos.y + 2*d.y}

			// Check if after next is on map.
			if afterNext.x < 0 || afterNext.x > len(grid[0])-1 || afterNext.y < 0 || afterNext.y > len(grid)-1 {
				continue
			}

			// Check if cheat is possible and saves enough time (the cheat itself costs 2).
			if grid[next.y][next.x] == '#' && (grid[afterNext.y][afterNext.x] == '.' || grid[afterNext.y][afterNext.x] == 'E') && distToEnd[pos]-distToEnd[afterNext]-2 >= minSaved {
				cheatCount++
			}
		}
	}

	fmt.Println(cheatCount)
}
