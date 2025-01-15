package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./7-input.txt")
	scanner := bufio.NewScanner(file)

	resultNoConcat := 0
	resultWithConcat := 0
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), ": ")
		wanted, _ := strconv.Atoi(spl[0])
		numbersRaw := strings.Split(spl[1], " ")
		numbers := make([]int, len(numbersRaw))
		for i, str := range numbersRaw {
			numbers[i], _ = strconv.Atoi(str)
		}

		if isEquationValid(numbers, wanted) {
			resultNoConcat += wanted
		}
		if isEquationValidWithConcat(numbers, wanted) {
			resultWithConcat += wanted
		}
	}
	fmt.Println("Part One:", resultNoConcat)
	fmt.Println("Part Two:", resultWithConcat)
}

func isEquationValid(numbers []int, wanted int) bool {
	// Use first number as starting accumulator.
	return checkEquation(numbers[1:], wanted, numbers[0])
}

func checkEquation(numbers []int, wanted, acc int) bool {
	if len(numbers) == 0 {
		return acc == wanted
	}
	return checkEquation(numbers[1:], wanted, acc+numbers[0]) ||
		checkEquation(numbers[1:], wanted, acc*numbers[0])
}

func isEquationValidWithConcat(numbers []int, wanted int) bool {
	// Use first number as starting accumulator.
	return checkEquationWithConcat(numbers[1:], wanted, numbers[0])
}

func checkEquationWithConcat(numbers []int, wanted, acc int) bool {
	if len(numbers) == 0 {
		return acc == wanted
	}

	concat, _ := strconv.Atoi(strconv.Itoa(acc) + strconv.Itoa(numbers[0]))
	return checkEquationWithConcat(numbers[1:], wanted, acc+numbers[0]) ||
		checkEquationWithConcat(numbers[1:], wanted, acc*numbers[0]) ||
		checkEquationWithConcat(numbers[1:], wanted, concat)
}
