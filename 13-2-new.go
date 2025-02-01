package main

import ( 
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


/*
TEMP

total: 10
A: 2
B: 3

combinaties:
A0: -
A1: -
A2: B2: prijs: 8
A3: - 
A4: -
A5: B0: prijs: 15

==> Observaties
- Prijs wordt steeds hoger, dus latere opties zijn eigenlijk niet relevant.
- Prijs gaat per optie regelmatig omhoog. TODO: dit checken bij eerste puzzel.
- Afstand tussen correcte A-aantallen is B (dit zal anders werken wanneer er x en y is) 


==> Dalend/stijgend
- A / B == 3: prijs is altijd hetzelfde.
- A / B < 3: prijs gaat omhoog (dus latere opties zijn niet relevant)
- A / B > 3: prijs gaat omlaag (dus laatste optie is relevant)


total: 10
A: 6
B: 2
A0: B5: prijs: 5 
A1: B2: prijs: 5

total: 10
A: 3
B: 2
A0: B5: prijs: 5 
A1: -
A2: B2: prijs: 8 
A3: -

total: 20
A: 7
B: 1
A0: B20: prijs: 20 
A1: B13: prijs: 16 
A2: B6: prijs: 12 

total: 20
A: 6
B: 1
A0: B20: prijs: 20 
A1: B14: prijs: 17 
A2: B8: prijs: 14 
A3: B2: prijs: 11 

total: 20
A: 5
B: 1
A0: B20: prijs: 20 
A1: B15: prijs: 18 
A2: B10: prijs: 16 
A3: B5: prijs: 14 
A4: B0: prijs: 12

total: 20
A: 3
B: 1
A0: B20: prijs: 20 
A1: B17: prijs: 20 
A2: B14: prijs: 20 
etc...

total: 20
A: 2
B: 1
A0: B20: prijs: 20 
A1: B18: prijs: 21 
A2: B16: prijs: 22 
A3: B5: prijs: 14 
A4: B0: prijs: 12
etc...

*/

/*

total: 10
A: 2
B: 3

t = 10
a = 2
b = 3

NA * a + NB * b = t
aNA + bNB = t
NA = (t - bNB) / a
NB = (t - aNA) / b


NA / NB = ( (t - bNB) / a ) / ((t - aNA) / b)
NA / NB = ( (t - bNB) / a ) * (b / (t - aNA)) = b(t-bNB) / a(t-aNA) = ...



aNA + bNB = t



*/

const (
    costA = 3
    costB = 1
    positionCorrection = 10000000000000
)

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

        prizeX += positionCorrection
        prizeY += positionCorrection

        // Empty line.
        scanner.Scan()

        cheapest := 0 
        // Try for every button A amount if prize can be reached with button B presses. 
        //TEMP disabled:
        for a := 0; a <= 1; a++ {
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
    // TEMP disabled:
    // fmt.Println(totalTokens)


    // TEMP:
    for i := 0; i < 1000000000000; i++ {
        rest := 10000000012748 % (93 + 26 * i)
        if rest == 0 {
            fmt.Println("FOUND", rest, i)
            return
        }
    }
}