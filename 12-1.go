package main

import (
    "bufio"
	"fmt"
	"os"
)

type co struct {
    x int
    y int
}

type region struct{
    perimeter int
    area int
}

func main() {
    file, _ := os.Open("./12-input.txt")
    scanner := bufio.NewScanner(file)

    grid := map[co]rune{}
    y := 0
    for scanner.Scan() {
        for x, r := range scanner.Text() {
            grid[co{x, y}] = r
        }
        y++
    }

    visited := map[co]bool{}
    regions := []region{}
    
    // Scan the grid and collect regions.
    for pos := range grid {
        if visited[pos] {
            continue
        }
        visited[pos] = true

        // Start a new region. BFS to find all connected plots with the same plant.
        r := region{}
        queue := []co{pos}
        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:]

            r.area++

            neighbors := []co{
                co{cur.x, cur.y-1},
                co{cur.x+1, cur.y},
                co{cur.x, cur.y+1},
                co{cur.x-1, cur.y},
            }
            for _, n := range neighbors {
                switch grid[n] {
                case grid[cur]:
                    // Find neighbors to visit next.
                    if visited[n] {
                        continue
                    }
                    visited[n] = true
                    queue = append(queue, n)
                default:
                    // Every neighbor which is a different plant needs a fence.
                    r.perimeter++
                }
            }
        }
        regions = append(regions, r)
    }

    totalPrice := 0
    for _, r := range regions {
        totalPrice += r.area * r.perimeter
    }
    fmt.Println(totalPrice)
}
