package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

const delay = time.Microsecond * 1

type co struct {
	x, y int
}

var moveBySymbol = map[rune]co{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
}

type box struct {
	pos co // Position of the left side of the box.
}

func (b box) fields() []co {
	return []co{b.pos, co{b.pos.x + 1, b.pos.y}}
}

type grid struct {
	width, height int
	walls         map[co]struct{}
	robot         co
	boxes         map[co]*box
}

func main() {
	var (
		walls = make(map[co]struct{})
		boxes = make(map[co]*box)
		robot co
		moves []co
	)

	file, _ := os.Open("15-input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Parse grid.
	width := utf8.RuneCountInString(scanner.Text()) * 2
	y := 0
	for ; scanner.Text() != ""; y++ {
		for i, rune := range scanner.Text() {
			x := 2 * i
			switch rune {
			case '@':
				robot = co{x, y}
			case '#':
				walls[co{x, y}], walls[co{x + 1, y}] = struct{}{}, struct{}{}
			case 'O':
				box := box{co{x, y}}
				boxes[co{x, y}], boxes[co{x + 1, y}] = &box, &box
			default:
				continue
			}
		}
		scanner.Scan()
	}
	height := y

	// Parse moves.
	for scanner.Scan() {
		line := scanner.Text()
		for _, s := range line {
			moves = append(moves, moveBySymbol[s])
		}
	}

	g := grid{width, height, walls, robot, boxes}

	g.paint()
	time.Sleep(delay)
	for _, m := range moves {
		g.update(m)
		g.paint()
		time.Sleep(delay)
	}

	// Calculate answer.
	uniqueBoxes := make(map[*box]struct{})
	for _, box := range g.boxes {
		uniqueBoxes[box] = struct{}{}
	}
	result := 0
	for box, _ := range uniqueBoxes {
		result += box.pos.x + 100*box.pos.y
	}
	fmt.Println(result)
}

func (g *grid) pushBox(b *box, d co) bool {
	boxesChecked := make(map[*box]struct{})

	var canPush func(b *box, d co) bool
	canPush = func(b *box, d co) bool {
		boxesChecked[b] = struct{}{}
		bNext := *b
		bNext.pos = co{b.pos.x + d.x, b.pos.y + d.y}

		for _, f := range bNext.fields() {
			if _, ok := g.walls[f]; ok {
				return false
			}
		}
		for _, f := range bNext.fields() {
			otherBox, ok := g.boxes[f]
			if !ok || otherBox == b {
				continue
			}
			_, ok = boxesChecked[otherBox]
			if ok {
				continue
			}
			if !canPush(otherBox, d) {
				return false
			}
		}
		return true
	}
	if !canPush(b, d) {
		return false
	}

	// Move the boxes.
	for b, _ := range boxesChecked {
		for _, f := range b.fields() {
			delete(g.boxes, f)
		}
		b.pos.x += d.x
		b.pos.y += d.y
	}
	for b, _ := range boxesChecked {
		for _, f := range b.fields() {
			g.boxes[f] = b
		}
	}
	return true
}

func (g *grid) update(d co) {
	robotNext := co{g.robot.x + d.x, g.robot.y + d.y}

	if _, ok := g.walls[robotNext]; ok {
		return
	}
	box, ok := g.boxes[robotNext]
	if ok && !g.pushBox(box, d) {
		return
	}

	g.robot = robotNext
}

func (g *grid) paint() {
	frame := make([][]rune, g.height)
	for y := range frame {
		frame[y] = make([]rune, g.width)
		for x := range frame[y] {
			if _, ok := g.walls[co{x, y}]; ok {
				frame[y][x] = '#'
				continue
			}
			frame[y][x] = '.'
		}
	}

	for _, b := range g.boxes {
		frame[b.pos.y][b.pos.x] = '['
		frame[b.pos.y][b.pos.x+1] = ']'
	}

	frame[g.robot.y][g.robot.x] = '@'

	var sb strings.Builder
	for _, line := range frame {
		for _, rune := range line {
			sb.WriteRune(rune)
		}
		sb.WriteRune('\n')
	}
	clearScreen()
	fmt.Print(sb.String())
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
