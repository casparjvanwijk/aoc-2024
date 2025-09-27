package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	adv = 0
	bxl = 1
	bst = 2
	jnz = 3
	bxc = 4
	out = 5
	bdv = 6
	cdv = 7
)

func main() {
	var (
		a, b, c int
		program []int8
	)

	// Parse input.
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	a, _ = strconv.Atoi(scanner.Text()[12:])
	scanner.Scan()
	b, _ = strconv.Atoi(scanner.Text()[12:])
	scanner.Scan()
	c, _ = strconv.Atoi(scanner.Text()[12:])
	scanner.Scan()
	scanner.Scan()
	for _, rune := range scanner.Text()[9:] {
		switch rune {
		case ',':
			continue
		default:
			num, _ := strconv.Atoi(string(rune))
			program = append(program, int8(num))
		}
	}

	// Run program.
	output := []string{}
	pointer := 0
	for pointer < len(program) {
		opcode, operand := program[pointer], program[pointer+1]
		switch opcode {
		case adv:
			a = int(float64(a) / math.Pow(float64(2), float64(comboOperand(operand, a, b, c))))
		case bdv:
			b = int(float64(a) / math.Pow(float64(2), float64(comboOperand(operand, a, b, c))))
		case cdv:
			c = int(float64(a) / math.Pow(float64(2), float64(comboOperand(operand, a, b, c))))
		case bxl:
			b = b ^ int(operand)
		case bst:
			b = comboOperand(operand, a, b, c) % 8
		case jnz:
			if a == 0 {
				break
			}
			pointer = int(operand)
			continue
		case bxc:
			b = b ^ c
		case out:
			output = append(output, strconv.Itoa(comboOperand(operand, a, b, c)%8))
		}
		pointer += 2
	}

	fmt.Println(strings.Join(output, ","))
}

func comboOperand(operand int8, a, b, c int) int {
	switch operand {
	case 0, 1, 2, 3:
		return int(operand)
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		panic("Invalid operand.")
	}
}
