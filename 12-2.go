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
    area int
    sides int
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
        outerCorners := 0
        innerCorners := 0
        queue := []co{pos}
        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:]

            r.area++
            inner, outer := countCorners(cur, grid)
            innerCorners += inner
            outerCorners += outer

            // Find neighbors to visit next.
            neighbors := []co{
                co{cur.x, cur.y-1},
                co{cur.x+1, cur.y},
                co{cur.x, cur.y+1},
                co{cur.x-1, cur.y},
            }
            for _, n := range neighbors {
                if grid[n] == grid[cur] && !visited[n] {
                    visited[n] = true
                    queue = append(queue, n)
                }
            }
        }
        
        // Inner corners are shared by three plots so must be divided by three to get the sides count.
        r.sides = outerCorners + innerCorners / 3
        regions = append(regions, r)
    }

    totalPrice := 0
    for _, r := range regions {
        totalPrice += r.area * r.sides
    }
    fmt.Println(totalPrice)
}

func countCorners(pos co, grid map[co]rune) (inner int, outer int) {
    // Corner types:
    // - Inner corners are shared between three plots.
    // - Outer corners touch only one plot.
    inner, outer = 0, 0

    cornerDeltas := []co{
        {1, -1}, // Top right
        {1, 1}, // Bottom right
        {-1, 1}, // Bottom left
        {-1, -1}, // Top left
    }

    // Check all corners of the plot.
    for _, d := range cornerDeltas {
        // Check for matching plots diagonally, vertically or horizontally.
        D := grid[co{pos.x + d.x, pos.y + d.y}] == grid[pos]
        V := grid[co{pos.x, pos.y + d.y}] == grid[pos]
        H := grid[co{pos.x + d.x, pos.y}] == grid[pos]

        // Check corner type. 
        // Diagrams: consider the top right corner of the bottom left X.
        switch {
            case (D && V && H) || (!D && !V && H) || (!D && V && !H):
                /* 
                X X   O O   X O
                X X   X X   X O
                */
                // No corner.
            case (!D && V && H) || (D && !V && H) || (D && V && !H):
                /* 
                X O   O X   X X
                X X   X X   X O
                */
                inner++
            case (!D && !V && !H) || (D && !V && !H):
                /* 
                O O   O X
                X O   X O
                */
                outer++
        }
    }
    return inner, outer
}