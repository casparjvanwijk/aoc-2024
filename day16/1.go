package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type co struct {
	x, y int
}

type direction struct {
	x, y int
}

// node is the combination of a position and the direction in which it was visited.
type node struct {
	pos co
	d   direction
}

const (
	emptySymbol = '.'
	wallSymbol  = '#'
	startSymbol = 'S'
	endSymbol   = 'E'
)

var directions = []direction{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func main() {
	walls := [][]bool{}
	var (
		startPos co
		endPos   co
	)

	file, _ := os.Open("./16-input.txt")
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := []bool{}
		for x, rune := range scanner.Text() {
			switch rune {
			case wallSymbol:
				line = append(line, true)
			case startSymbol:
				startPos = co{x, y}
				line = append(line, false)
			case endSymbol:
				endPos = co{x, y}
				line = append(line, false)
			default:
				line = append(line, false)
			}
		}
		walls = append(walls, line)
	}

	// visited is map of visited nodes and their previous node.
	visited := make(map[node]node)

	paint(walls, startPos, endPos, visited, nil)

	type queueItem struct {
		node node
		cost int
		prev node
	}
	queue := []queueItem{
		{
			node: node{pos: startPos, d: direction{1, 0}},
			cost: 0,
		},
	}

	insertSorted := func(queue []queueItem, item queueItem) []queueItem {
		// Find first index where cost is higher than inserting cost.
		idx := slices.IndexFunc(queue, func(qi queueItem) bool {
			return qi.cost > item.cost
		})
		// If inserting cost is highest, append to end.
		if idx == -1 {
			return append(queue, item)
		}
		// Insert before.
		return slices.Insert(queue, idx, item)
	}

	var (
		cost    int
		endNode node
	)
	for len(queue) > 0 {
		// time.Sleep(time.Millisecond * 10)

		cur := queue[0]
		queue = queue[1:]
		if _, ok := visited[cur.node]; ok {
			continue
		}
		visited[cur.node] = cur.prev
		// paint(walls, startPos, endPos, visited, nil)

		if cur.node.pos == endPos {
			cost = cur.cost
			endNode = cur.node
			break
		}

		for _, d := range directions {
			if d == (direction{-cur.node.d.x, -cur.node.d.y}) {
				// Do not go back.
				continue
			}
			next := node{
				pos: co{cur.node.pos.x + d.x, cur.node.pos.y + d.y},
				d:   d,
			}
			if walls[next.pos.y][next.pos.x] {
				continue
			}
			cost := cur.cost + 1
			if d != cur.node.d {
				cost += 1000
			}
			queue = insertSorted(queue, queueItem{
				node: next,
				cost: cost,
				prev: cur.node,
			})
		}
	}

	route := []co{}
	cur := endNode
	for {
		route = append(route, cur.pos)
		if cur.pos == startPos {
			break
		}
		cur = visited[cur]
	}

	paint(walls, startPos, endPos, visited, route)
	fmt.Println("Cost:", cost)
}

func paint(walls [][]bool, startPos, endPos co, visited map[node]node, route []co) {
	// Create empty frame with walls.
	frame := [][]rune{}
	for _, line := range walls {
		frameLine := []rune{}
		for _, field := range line {
			if field {
				frameLine = append(frameLine, wallSymbol)
				continue
			}
			frameLine = append(frameLine, emptySymbol)
		}
		frame = append(frame, frameLine)
	}

	// Add start and end positions.
	frame[startPos.y][startPos.x] = startSymbol
	frame[endPos.y][endPos.x] = endSymbol

	// Add visited fields.
	for node := range visited {
		frame[node.pos.y][node.pos.x] = '▒'
	}

	// Add shortest route if available.
	if route != nil {
		for _, pos := range route {
			frame[pos.y][pos.x] = '█'
		}

	}

	// Print.
	clearScreen()
	var builder strings.Builder
	for _, line := range frame {
		for _, rune := range line {
			builder.WriteRune(rune)
		}
		builder.WriteRune('\n')
	}
	fmt.Print(builder.String())
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
