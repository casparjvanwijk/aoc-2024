package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	width, height = 101, 103
	seconds       = 99999
)

type co struct {
	x int
	y int
}

type robot struct {
	pos co
	vX  int
	vY  int
}

func main() {
	file, _ := os.Open("./14-input.txt")
	scanner := bufio.NewScanner(file)
	reg, _ := regexp.Compile(`p=(.+),(.+)\sv=(.+),(.+)`)

	// Parse robots.
	robots := []robot{}
	for scanner.Scan() {
		res := reg.FindStringSubmatch(scanner.Text())
		x, _ := strconv.Atoi(res[1])
		y, _ := strconv.Atoi(res[2])
		vX, _ := strconv.Atoi(res[3])
		vY, _ := strconv.Atoi(res[4])
		robots = append(robots, robot{co{x, y}, vX, vY})
	}

	for i := range seconds {
		// Move the robots.
		for j := range robots {
			robots[j].pos.x = mod(robots[j].pos.x+robots[j].vX, width)
			robots[j].pos.y = mod(robots[j].pos.y+robots[j].vY, height)
		}

		// Put the robot counts on the grid.
		grid := map[co]int{}
		noDoubles := true
		for _, r := range robots {
			if grid[r.pos] == 1 {
				noDoubles = false
			}
			grid[r.pos]++
		}

		clearScreen()
		fmt.Println("Seconds:", i+1)

		// Print the grid.
		for y := 0; y < height; y++ {
			line := ""
			for x := 0; x < width; x++ {
				if grid[co{x, y}] == 0 {
					line += "."
					continue
				}
				line += strconv.Itoa(grid[co{x, y}])
			}
			fmt.Println(line)
		}

		// Assuming that the robots have a fixed position when making the Christmas tree,
		// stop if all fields have max 1 robot, as that is probably the grid with the Christmas tree.
		if noDoubles {
			break
		}

		time.Sleep(time.Millisecond * 5)
	}
}

// mod calculates the Euclidian modulo (instead of the truncated modulo which is % in Go).
func mod(a, b int) int {
	return (a%b + b) % b
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}