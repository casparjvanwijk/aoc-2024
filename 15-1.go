package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
    "unicode/utf8"
)

type co struct {
    x int
    y int
}

type field int

const (
    fieldWall field = 1
    fieldBox field = 2
)

var directionBySymbol = map[rune]co{
    '^': co{0, -1},
    '>': co{1, 0}, 
    'v': co{0, 1}, 
    '<': co{-1, 0}, 
}

//TODO: clean up, refactor
func main() {
    file, _ := os.Open("15-input.txt")
    scanner := bufio.NewScanner(file)
    scanner.Scan()
        
    grid := make(map[co]field)
    robot := co{}
    width := utf8.RuneCountInString(scanner.Text())
    y := 0
    for ; scanner.Text() != ""; y++ {
        for x, rune := range scanner.Text() {
            //TODO: extract field types / symbols etc.
            switch rune {
            case '#':
                grid[co{x, y}] = fieldWall
            case 'O':
                grid[co{x, y}] = fieldBox
            case '@':
                robot = co{x, y}
            }
        }
        scanner.Scan()
    }
    height := y

    moves := ""
    for scanner.Scan() {
        moves += scanner.Text()
    }

    //TODO: order? loop?
    clearScreen()
    paint(grid, width, height, robot)
    time.Sleep(time.Millisecond * 10)

    for _, move := range moves {
        d := directionBySymbol[move]
        targetPos := co{robot.x + d.x, robot.y + d.y}

        switch grid[targetPos] {
        case fieldWall:
            // Do nothing.
        case fieldBox:
            if (moveBox(grid, targetPos, d)) {
                robot = targetPos
            }
        default: 
            robot = targetPos
        }

        clearScreen()
        paint(grid, width, height, robot)
        time.Sleep(time.Millisecond * 10)
    }

    result := 0
    for pos, field := range grid {
        if field != fieldBox {
            continue
        } 
        result += pos.x + 100 * pos.y
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
    for y := range height {
        line := ""

        for x := range width {
            //TODO: refactor robot placement:
            if (co{x, y} == robot) {
                line += "@"
                continue
            } 

            switch grid[co{x, y}] {
            case fieldWall:
                line += "#"
            case fieldBox:
                line += "O"
            default:
                line += "."
            }
        }
        fmt.Println(line)
    }
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}