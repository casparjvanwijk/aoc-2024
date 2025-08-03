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
    fieldEmpty field = 0
    fieldWall field = 1
    fieldBoxLeft field = 2
    fieldBoxRight field = 3
    //TODO: delete?
    fieldBox field = 4 
    symbolEmpty = '.'
    symbolWall = '#'
    symbolBoxLeft = '['
    symbolBoxRight = ']'
    //TODO: delete?
    symbolBox = 'O'
    symbolRobot = '@'
)

var symbolByField = map[field]rune{
    fieldEmpty: symbolEmpty,
    fieldWall: symbolWall,
    fieldBoxLeft: symbolBoxLeft,
    fieldBoxRight: symbolBoxRight,
} 

var directionByMove = map[rune]co{
    '^': co{0, -1},
    '>': co{1, 0}, 
    'v': co{0, 1}, 
    '<': co{-1, 0}, 
}

func main() {
    file, _ := os.Open("15-input.txt")
    scanner := bufio.NewScanner(file)
    scanner.Scan()

    // Parse grid.
    grid := make(map[co]field)
    robot := co{}
    width := utf8.RuneCountInString(scanner.Text()) * 2
    y := 0
    for ; scanner.Text() != ""; y++ {
        for i, symbol := range scanner.Text() {
            x := 2 * i
            switch symbol {
            case symbolRobot:
                robot = co{x, y}
            case symbolWall: 
                grid[co{x, y}], grid[co{x+1, y}] = fieldWall, fieldWall
            case symbolBox:
                grid[co{x, y}], grid[co{x+1, y}] = fieldBoxLeft, fieldBoxRight
            default:
                continue
            }
        }
        scanner.Scan()
    }
    height := y

    // Parse moves.
    moves := ""
    for scanner.Scan() {
        moves += scanner.Text()
    }

    clearScreen()
    paint(grid, width, height, robot)
    time.Sleep(time.Millisecond * 10)

//TEMP:
return

    for _, move := range moves {
        // Move robot.
        d := directionByMove[move]
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

    // Calculate answer.
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
            if (co{x, y} == robot) {
                line += string(symbolRobot)
                continue
            } 
            line += string(symbolByField[grid[co{x, y}]])
        }
        fmt.Println(line)
    }
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}