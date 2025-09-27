package main

import ( 
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    file, _ := os.Open("./13-input.txt")
    scanner := bufio.NewScanner(file)

//TEMP:
fmt.Println(file, scanner)
//aX: 26
//bX: 67
//verhouding 1:1: (1 *) 26 + (1 *) 67 = 93
fmt.Println(10000000012748 / 120)
//TODO: zoek naar verhouding A:B waarbij het grote getal deelbaar is door het totaal
//van A en B in die verhouding.
//Vraag: hoe weet ik wanneer ik alle mogelijkheden heb gehad?
return

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

        // Find the cheapest combination. -1 means no combination has been found.
        cheapest := -1 
outer:
        // Try every button A amount with every button B amount, both up until 100.
        for a := 0; a <= 100; a++ {
            x := a * aX
            y := a * aY
            if x > prizeX || y > prizeY {
                break
            }
            for b := 1; b <= 100; b++ {
                x += bX
                y += bY
                if x == prizeX && y == prizeY {
                    price := a * 3 + b
                    if cheapest == -1 || price < cheapest {
                        cheapest = price
                    }
                    continue outer
                }
                if x > prizeX || y > prizeY {
                    continue outer
                }
            }
        }
        // Ignore if no combination was found.
        if cheapest != -1 {
            totalTokens += cheapest
        }
    }
    fmt.Println(totalTokens)
}