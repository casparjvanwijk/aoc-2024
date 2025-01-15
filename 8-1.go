package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

const emptyCell = '.'

type co struct {
	x int
	y int
}

func (c co) onMap(xMax int, yMax int) bool {
	return c.x >= 0 && c.x <= xMax && c.y >= 0 && c.y <= yMax
}

func main() {
	file, _ := os.Open("./8-input.txt")
	scanner := bufio.NewScanner(file)

	// Find all nodes.
	xMax := 0
	y := 0
	nodesByFreq := map[rune][]co{}
	for scanner.Scan() {
		if xMax == 0 {
			xMax = utf8.RuneCountInString(scanner.Text()) - 1
		}
		for x, r := range scanner.Text() {
			if r == emptyCell {
				continue
			}
			nodesByFreq[r] = append(nodesByFreq[r], co{x, y})
		}
		y++
	}
	yMax := y - 1

	// Find all unique antinode locations.
	antiNodes := map[co]bool{}
	for _, nodes := range nodesByFreq {
		// Pair each node with each other node with the same frequency.
		for i, n1 := range nodes {
			for _, n2 := range nodes[i+1:] {
				dX := n2.x - n1.x
				dY := n2.y - n1.y

				// Antinode 1 is before node 1.
				anti1 := co{n1.x - dX, n1.y - dY}
				// Antinode 2 is after node 2.
				anti2 := co{n2.x + dX, n2.y + dY}

				if anti1.onMap(xMax, yMax) {
					antiNodes[anti1] = true
				}
				if anti2.onMap(xMax, yMax) {
					antiNodes[anti2] = true
				}
			}
		}
	}
	fmt.Println(len(antiNodes))
}
