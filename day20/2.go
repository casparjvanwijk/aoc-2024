package main

import (
	"bufio"
	"fmt"
	"os"
)

// minSaved is the least amount of time saved by cheating to count the cheat.
const minSaved = 100

// cheatMax is the max length of a cheat.
const cheatMax = 20

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
		// Check all target fields within cheat range of current field.
		for i := -cheatMax; i <= cheatMax; i++ {
			for j := -(cheatMax - abs(i)); j <= (cheatMax - abs(i)); j++ {
				x, y := pos.x+j, pos.y+i
				// Check if field is on map.
				if x < 0 || x > len(grid[0])-1 || y < 0 || y > len(grid)-1 {
					continue
				}
				// Check if field is valid cheat endpoint.
				if grid[y][x] != '.' && grid[y][x] != 'E' {
					continue
				}
				// If cheat saves enough distance, increment counter.
				cheatDist := abs(i) + abs(j)
				if distToEnd[pos]-distToEnd[co{x, y}]-cheatDist >= minSaved {
					cheatCount++
				}
			}
		}
	}

	fmt.Println(cheatCount)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
