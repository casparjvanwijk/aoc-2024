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

type move struct {
	x, y int
}

var moveBySymbol = map[rune]move{
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

type game struct {
	width, height int
	walls         map[co]struct{}
	robot         co
	boxes         []*box
	// boxMap keeps track of the boxes indexed by position for fast lookup.
	boxMap map[co]*box
}

func newGame(width, height int, walls map[co]struct{}, robot co, boxes []*box) game {
	g := game{
		width:  width,
		height: height,
		walls:  walls,
		robot:  robot,
		boxes:  boxes,
	}
	g.refreshBoxMap()
	return g
}

func (g *game) refreshBoxMap() {
	g.boxMap = make(map[co]*box)
	for _, b := range g.boxes {
		for _, f := range b.fields() {
			g.boxMap[f] = b
		}
	}
}

func (g *game) update(d move) {
	robotNext := co{g.robot.x + d.x, g.robot.y + d.y}
	// Check for wall collision.
	if _, ok := g.walls[robotNext]; ok {
		return
	}
	// Check for box collision.
	b, ok := g.boxMap[robotNext]
	if ok {
		canMove, boxesTouched := g.canMoveBox(b, d)
		if !canMove {
			return
		}
		for _, b := range boxesTouched {
			b.pos.x += d.x
			b.pos.y += d.y
		}
		g.refreshBoxMap()
	}
	g.robot = robotNext
}

// canMoveBox checks if a box can be moved and also returns all touched boxes.
func (g *game) canMoveBox(b *box, d move) (canMove bool, boxesTouched []*box) {
	visited := make(map[*box]struct{})
	queue := []*box{b}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if _, ok := visited[cur]; ok {
			continue
		}
		visited[cur] = struct{}{}
		boxesTouched = append(boxesTouched, cur)

		for _, f := range cur.fields() {
			fNext := co{f.x + d.x, f.y + d.y}
			if _, ok := g.walls[fNext]; ok {
				return false, nil
			}
			if boxNext, ok := g.boxMap[fNext]; ok && boxNext != cur {
				queue = append(queue, boxNext)
			}
		}
	}
	return true, boxesTouched
}

func (g *game) paint() {
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

func main() {
	var (
		walls = make(map[co]struct{})
		boxes = []*box{}
		robot co
		moves []move
	)

	file, _ := os.Open("15-input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Parse game.
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
				boxes = append(boxes, &box{co{x, y}})
			default:
				continue
			}
		}
		scanner.Scan()
	}
	height := y

	g := newGame(width, height, walls, robot, boxes)

	// Parse moves.
	for scanner.Scan() {
		line := scanner.Text()
		for _, s := range line {
			moves = append(moves, moveBySymbol[s])
		}
	}

	g.paint()
	time.Sleep(delay)
	for _, m := range moves {
		g.update(m)
		g.paint()
		time.Sleep(delay)
	}

	// Calculate answer.
	result := 0
	for _, b := range boxes {
		result += b.pos.x + 100*b.pos.y
	}
	fmt.Println(result)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
