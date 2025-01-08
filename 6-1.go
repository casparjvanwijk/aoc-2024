package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	tick = 70 * time.Millisecond

	fieldNotVisited = "."
	fieldVisited    = "X"
	fieldObstacle   = "#"

	guardDirUp    guardDir = "^"
	guardDirRight guardDir = ">"
	guardDirDown  guardDir = "v"
	guardDirLeft  guardDir = "<"
)

type coords struct {
	x int
	y int
}

type guardDir string

type grid struct {
	fields   map[coords]string
	width    int
	height   int
	guardPos coords
	guardDir guardDir
}

func main() {
	input, _ := os.ReadFile("./6-input.txt")
	g := newGrid(string(input))
	for {
		g.paint()
		time.Sleep(tick)
		g.update()
		if g.guardOffMap() {
			fmt.Printf("%v: %v\n", fieldVisited, g.visitedCount())
			return
		}
	}
}

func newGrid(s string) grid {
	fields := map[coords]string{}
	lines := strings.Split(s, "\n")
	guardPos := coords{}

	for y, line := range lines {
		x := 0
		for _, r := range line {
			field := string(r)

			// Set guard.
			if string(r) == string(guardDirUp) {
				guardPos = coords{x, y}
				field = fieldVisited
			}

			fields[coords{x, y}] = field
			x++
		}
	}

	return grid{
		fields:   fields,
		width:    utf8.RuneCountInString(lines[0]),
		height:   len(lines),
		guardPos: guardPos,
		guardDir: guardDirUp,
	}
}

func (g *grid) update() {
	// Move guard.
	for {
		// Get next position.
		nextGuardPos := coords{}
		switch g.guardDir {
		case guardDirUp:
			nextGuardPos = coords{g.guardPos.x, g.guardPos.y - 1}
		case guardDirRight:
			nextGuardPos = coords{g.guardPos.x + 1, g.guardPos.y}
		case guardDirDown:
			nextGuardPos = coords{g.guardPos.x, g.guardPos.y + 1}
		case guardDirLeft:
			nextGuardPos = coords{g.guardPos.x - 1, g.guardPos.y}
		}

		// Move to next position if there is no obstacle.
		if g.fields[nextGuardPos] != string(fieldObstacle) {
			g.guardPos = nextGuardPos
			break
		}

		// Turn at obstacle.
		switch g.guardDir {
		case guardDirUp:
			g.guardDir = guardDirRight
		case guardDirRight:
			g.guardDir = guardDirDown
		case guardDirDown:
			g.guardDir = guardDirLeft
		case guardDirLeft:
			g.guardDir = guardDirUp
		}
	}

	// Mark the new field as visited.
	g.fields[coords{g.guardPos.x, g.guardPos.y}] = fieldVisited
}

func (g grid) guardOffMap() bool {
	return g.guardPos.x < 0 || g.guardPos.x >= g.width ||
		g.guardPos.y < 0 || g.guardPos.y >= g.height
}

func (g grid) visitedCount() int {
	count := 0
	for y := range g.height {
		for x := range g.width {
			if g.fields[coords{x, y}] == fieldVisited {
				count++
			}
		}
	}
	return count
}

func (g grid) paint() {
	clearScreen()

	for y := range g.height {
		line := ""
		for x := range g.width {
			// Add the guard.
			if x == g.guardPos.x && y == g.guardPos.y {
				line += string(g.guardDir)
				continue
			}
			line += g.fields[coords{x, y}]
		}
		fmt.Println(line)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
