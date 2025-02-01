package main

import ( 
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const costA, costB = 3, 1

func main() {
    file, _ := os.Open("./13-input.txt")
    scanner := bufio.NewScanner(file)

    totalTokens := 0
    for scanner.Scan() {
        spl := strings.Split(scanner.Text(), ",")
        aX, _ := strconv.Atoi(spl[0][12:])
        aY, _ := strconv.Atoi(spl[1][3:])
        
        scanner.Scan()
        spl = strings.Split(scanner.Text(), ",")
        bX, _ := strconv.Atoi(spl[0][12:])
        bY, _ := strconv.Atoi(spl[1][3:])

        scanner.Scan()
        spl = strings.Split(scanner.Text(), ",")
        prizeX, _ := strconv.Atoi(spl[0][9:])
        prizeY, _ := strconv.Atoi(spl[1][3:])

        // Empty line.
        scanner.Scan()

        cheapest := 0 
        // Check for every button A amount if prize can be reached with button B presses. 
        for a := 0; a <= 100; a++ {
            x, y := a * aX, a * aY
            if x > prizeX || y > prizeY {
                break
            }

            restX, restY := prizeX - x, prizeY - y
            if restX % bX != 0 || restY % bY != 0 {
                continue
            }
            
            b := restX / bX
            if restY / bY != b {
                continue
            }

            cost := costA * a + costB * b 
            if cheapest == 0 || cost < cheapest {
                cheapest = cost
            }
        }
        totalTokens += cheapest
    }
    fmt.Println(totalTokens)
}