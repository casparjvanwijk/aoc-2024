package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

type (
	fieldType int
	direction int
)

const (
	debug bool = false
	tick       = 20 * time.Millisecond

	fieldTypePlain    fieldType = 0
	fieldTypeObstacle fieldType = 1

	dirUp    direction = 0
	dirRight direction = 1
	dirDown  direction = 2
	dirLeft  direction = 3

	guardInitDir = dirUp

	symbolEmpty         = '.'
	symbolObstacle      = '#'
	symbolAddedObstacle = 'O'
	symbolHorizontal    = '-'
	symbolVertical      = '|'
	symbolCrossing      = '+'
	symbolRepeated      = 'R'
)

var symbolGuard = map[direction]rune{
	dirUp:    '^',
	dirRight: '>',
	dirDown:  'v',
	dirLeft:  '<',
}

type co struct {
	x int
	y int
}

type field struct {
	typ      fieldType
	visited  map[direction]bool
	repeated bool // Whether the field was visited more than once in the same direction.
}

type grid struct {
	fields           map[co]field
	width            int
	height           int
	guardInitPos     co
	guardPos         co
	guardDir         direction
	addedObstaclePos co
}

func main() {
	startTime := time.Now()
	input, _ := os.ReadFile("./6-input.txt")

	// Find all guard positions during a normal run.
	guardPositions := map[co]bool{}
	g := newGrid(string(input))
	for {
		if debug {
			g.paint()
			time.Sleep(tick)
		}
		guardPositions[g.guardPos] = true
		g.update()
		if g.guardOffMap() {
			break
		}
	}

	// The initial guard position cannot be used for the added obstacle.
	delete(guardPositions, g.guardInitPos)

	// Run the simulation with added obstacles on each guard position and look for loops.
	totalRuns := len(guardPositions)
	runCount := 0
	totalLoops := 0
outerLoop:
	for obstaclePos := range guardPositions {
		runCount++
		clearScreen()
		fmt.Printf("%v / %v\n", runCount, totalRuns)

		g.reset()
		g.addedObstaclePos = obstaclePos

		// Run the simulation.
		for {
			if debug {
				g.paint()
				fmt.Println("Obstacle:", obstaclePos)
				fmt.Println("Total loops:", totalLoops)
				time.Sleep(tick)
			}
			g.update()
			if g.guardOffMap() {
				continue outerLoop
			}
			if g.guardInLoop() {
				totalLoops++
				continue outerLoop
			}
		}
	}
	fmt.Println(totalLoops)
	fmt.Println(time.Since(startTime))
}

func newGrid(s string) grid {
	fields := map[co]field{}
	lines := strings.Split(s, "\n")
	guardPos := co{}

	for y, line := range lines {
		x := 0
		for _, r := range line {
			switch r {
			case symbolEmpty:
				fields[co{x, y}] = field{
					typ:     fieldTypePlain,
					visited: map[direction]bool{},
				}
			case symbolObstacle:
				fields[co{x, y}] = field{
					typ:     fieldTypeObstacle,
					visited: map[direction]bool{},
				}
			case symbolGuard[guardInitDir]:
				guardPos = co{x, y}
				f := field{
					typ:     fieldTypePlain,
					visited: map[direction]bool{},
				}
				f.visited[guardInitDir] = true
				fields[co{x, y}] = f
			}
			x++
		}
	}

	return grid{
		fields:           fields,
		width:            utf8.RuneCountInString(lines[0]),
		height:           len(lines),
		guardPos:         guardPos,
		guardInitPos:     guardPos,
		guardDir:         guardInitDir,
		addedObstaclePos: co{-1, -1},
	}
}

func (g *grid) reset() {
	// Reset fields.
	for pos, field := range g.fields {
		field.visited = map[direction]bool{}
		field.repeated = false
		g.fields[pos] = field
	}

	// Reset guard.
	g.guardPos = g.guardInitPos
	g.guardDir = guardInitDir
}

func (g *grid) update() {
	nextGuardPos := g.nextGuardPos()

	// Turn right if there is an obstacle, otherwise move forward.
	if g.fields[nextGuardPos].typ == fieldTypeObstacle || nextGuardPos == g.addedObstaclePos {
		g.guardDir = (g.guardDir + 1) % 4
	} else {
		g.guardPos = nextGuardPos
	}

	if g.guardOffMap() {
		return
	}

	// If the field was already visited in the same direction: mark as repeated.
	if g.fields[g.guardPos].visited[g.guardDir] {
		cpy := g.fields[g.guardPos]
		cpy.repeated = true
		g.fields[g.guardPos] = cpy
		return
	}

	// Mark the new field as visited.
	// If the guard only turned: still revisit the current field to make a crossing.
	g.fields[g.guardPos].visited[g.guardDir] = true
}

func (g grid) nextGuardPos() co {
	switch g.guardDir {
	case dirUp:
		return co{g.guardPos.x, g.guardPos.y - 1}
	case dirRight:
		return co{g.guardPos.x + 1, g.guardPos.y}
	case dirDown:
		return co{g.guardPos.x, g.guardPos.y + 1}
	case dirLeft:
		return co{g.guardPos.x - 1, g.guardPos.y}
	default:
		return g.guardPos
	}
}

func (g grid) guardOffMap() bool {
	return g.guardPos.x < 0 || g.guardPos.x >= g.width ||
		g.guardPos.y < 0 || g.guardPos.y >= g.height
}

func (g grid) guardInLoop() bool {
	return g.fields[g.guardPos].repeated
}

func (g grid) paint() {
	clearScreen()

	for y := range g.height {
		var line strings.Builder
		for x := range g.width {
			field := g.fields[co{x, y}]
			var r rune
			switch {
			case x == g.guardPos.x && y == g.guardPos.y:
				r = symbolGuard[g.guardDir]
			case x == g.addedObstaclePos.x && y == g.addedObstaclePos.y:
				r = symbolAddedObstacle
			case field.typ == fieldTypeObstacle:
				r = symbolObstacle
			case field.repeated:
				r = symbolRepeated
			case len(field.visited) == 0:
				r = symbolEmpty
			case len(field.visited) == 1:
				if field.visited[dirUp] || field.visited[dirDown] {
					r = symbolVertical
					break
				}
				if field.visited[dirRight] || field.visited[dirLeft] {
					r = symbolHorizontal
					break
				}
			case len(field.visited) == 2:
				if field.visited[dirUp] && field.visited[dirDown] {
					r = symbolVertical
					break
				}
				if field.visited[dirRight] && field.visited[dirLeft] {
					r = symbolHorizontal
					break
				}
				r = symbolCrossing
			case len(field.visited) > 2:
				r = symbolCrossing
			}
			line.WriteRune(r)
		}
		fmt.Println(line.String())
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
