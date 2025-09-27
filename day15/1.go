package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

const delay = time.Millisecond * 2

type co struct {
	x int
	y int
}

type field int

const (
	fieldEmpty  field = 0
	fieldWall   field = 1
	fieldBox    field = 2
	symbolEmpty       = '.'
	symbolWall        = '#'
	symbolBox         = 'O'
	symbolRobot       = '@'
)

var symbolByField = map[field]rune{
	fieldEmpty: symbolEmpty,
	fieldWall:  symbolWall,
	fieldBox:   symbolBox,
}

var directionBySymbol = map[rune]co{
	'^': co{0, -1},
	'>': co{1, 0},
	'v': co{0, 1},
	'<': co{-1, 0},
}

func main() {
	fieldBySymbol := make(map[rune]field)
	for f, s := range symbolByField {
		fieldBySymbol[s] = f
	}

	file, _ := os.Open("15-input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Parse grid.
	grid := make(map[co]field)
	robot := co{}
	width := utf8.RuneCountInString(scanner.Text())
	y := 0
	for ; scanner.Text() != ""; y++ {
		for x, symbol := range scanner.Text() {
			switch symbol {
			case symbolRobot:
				robot = co{x, y}
			case symbolEmpty:
				continue
			default:
				grid[co{x, y}] = fieldBySymbol[symbol]
			}
		}
		scanner.Scan()
	}
	height := y

	// Parse moves.
	moves := []co{}
	for scanner.Scan() {
		line := scanner.Text()
		for _, s := range line {
			moves = append(moves, directionBySymbol[s])
		}
	}

	paint(grid, width, height, robot)
	time.Sleep(delay)

	for _, m := range moves {
		// Move robot.
		targetPos := co{robot.x + m.x, robot.y + m.y}
		switch grid[targetPos] {
		case fieldWall:
			// Do nothing.
		case fieldBox:
			if moveBox(grid, targetPos, m) {
				robot = targetPos
			}
		default:
			robot = targetPos
		}

		paint(grid, width, height, robot)
		time.Sleep(delay)
	}

	// Calculate answer.
	result := 0
	for pos, field := range grid {
		if field != fieldBox {
			continue
		}
		result += pos.x + 100*pos.y
	}
	fmt.Println(result)
}

// moveBox tries to move a box in a direction and returns true if the box was moved.
func moveBox(grid map[co]field, pos, d co) bool {
	targetPos := co{pos.x + d.x, pos.y + d.y}
	switch grid[targetPos] {
	case fieldWall:
		return false
	case fieldBox:
		if moveBox(grid, targetPos, d) {
			delete(grid, pos)
			grid[targetPos] = fieldBox
			return true
		}
		return false
	default:
		delete(grid, pos)
		grid[targetPos] = fieldBox
		return true
	}
}

func paint(grid map[co]field, width, height int, robot co) {
	var sb strings.Builder
	for y := range height {
		for x := range width {
			if (co{x, y} == robot) {
				sb.WriteRune(symbolRobot)
				continue
			}
			sb.WriteRune(symbolByField[grid[co{x, y}]])
		}
		sb.WriteRune('\n')
	}
	clearScreen()
	fmt.Print(sb.String())
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
