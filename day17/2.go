package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
		a, b, c    int
		program    []int8
		programStr string
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
	programStr = scanner.Text()[9:]
	for _, rune := range programStr {
		switch rune {
		case ',':
			continue
		default:
			num, _ := strconv.Atoi(string(rune))
			program = append(program, int8(num))
		}
	}

	//TODO: patroon begrijpen: achterste cijfers van a bepalen voorste cijfers van program?
	// Komt door %8 bij output?
	// Original value for a is ignored.
	for a = 100000000000000; a < 999999999999999; a++ {
		fmt.Println(a)
		if outputsSelf(program, a, b, c) {
			fmt.Println(a)
			return
		}
	}
	fmt.Println("No value for A found.")
}

func outputsSelf(program []int8, a, b, c int) bool {
	// output := []int8{}
	outCount := 0
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
			val := int8(comboOperand(operand, a, b, c) % 8)
			if val != program[outCount] {
				return false
			}
			outCount++
			// output = append(output, val)
		}
		pointer += 2
	}
	return outCount == len(program)
	//TODO: clean up?
	// return slices.Equal(output, program)
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
