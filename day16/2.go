package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type co struct {
	x, y int
}

var nullPos = co{-1, -1}

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

type queue []queueItem

type queueItem struct {
	node node
	cost int
	prev node
}

func (q *queue) insertSorted(node node, cost int, prev node) {
	item := queueItem{node: node, cost: cost, prev: prev}
	// Find first index where cost is higher than inserting cost.
	idx := slices.IndexFunc(*q, func(qi queueItem) bool {
		return qi.cost > item.cost
	})
	// If inserting cost is highest, append to end.
	if idx == -1 {
		*q = append(*q, item)
		return
	}
	// Insert before.
	*q = slices.Insert(*q, idx, item)
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

	paint(walls, startPos, endPos, nil, nil)

	var q queue
	q.insertSorted(
		node{pos: startPos, d: direction{1, 0}},
		0,
		node{pos: nullPos},
	)

	cost := map[node]int{}
	prev := map[node][]node{}

	var (
		endNode    node
		endReached bool
	)
	for len(q) > 0 {
		time.Sleep(time.Millisecond * 10)

		cur := q[0]
		q = q[1:]
		if c, ok := cost[cur.node]; ok {
			if cur.cost == c {
				prev[cur.node] = append(prev[cur.node], cur.prev)
			}
			continue
		}
		cost[cur.node] = cur.cost
		prev[cur.node] = []node{cur.prev}
		// Do not break immediately when end is reached, because multiple cheapest routes could be found.
		if endReached && cur.cost > cost[endNode] {
			break
		}

		visited := []co{}
		for node := range cost {
			visited = append(visited, node.pos)
		}
		paint(walls, startPos, endPos, visited, nil)

		if cur.node.pos == endPos {
			endReached = true
			endNode = cur.node
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
			q.insertSorted(next, cost, cur.node)
		}
	}

	routeFields := map[co]struct{}{}
	routeQueue := []node{endNode}
	for len(routeQueue) > 0 {
		cur := routeQueue[0]
		routeQueue = routeQueue[1:]
		if cur.pos == nullPos {
			continue
		}
		routeFields[cur.pos] = struct{}{}
		routeQueue = append(routeQueue, prev[cur]...)
	}

	visited := []co{}
	for node := range cost {
		visited = append(visited, node.pos)
	}
	routeArr := []co{}
	for pos := range routeFields {
		routeArr = append(routeArr, pos)
	}
	paint(walls, startPos, endPos, visited, routeArr)
	fmt.Println("Cost:", cost[endNode])
	fmt.Println("Fields on shortest routes:", len(routeFields))
}

func paint(walls [][]bool, startPos, endPos co, visited []co, routeFields []co) {
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
	for _, pos := range visited {
		frame[pos.y][pos.x] = '▒'
	}

	// Add fields visited on shortest routes if available.
	for _, pos := range routeFields {
		frame[pos.y][pos.x] = '█'
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
