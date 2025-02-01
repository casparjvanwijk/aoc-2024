package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.ReadFile("./11-input.txt")
	spl := strings.Split(string(in), " ")
	stones := make([]int, len(spl))
	for i, str := range spl {
		stones[i], _ = strconv.Atoi(str)
	}

	total := 0
	for _, stone := range stones {
		total += countStones(stone, 75)
	}
	fmt.Println(total)
}

type key struct {
	val    int
	blinks int
}

var memo = map[key]int{}

func countStones(val int, blinks int) int {
	if m, ok := memo[key{val, blinks}]; ok {
		return m
	}

	result := countStonesInner(val, blinks)
	memo[key{val, blinks}] = result
	return result
}

func countStonesInner(val int, blinks int) int {
	if blinks == 0 {
		return 1
	}

	if val == 0 {
		return countStones(1, blinks-1)
	}

	s := strconv.Itoa(val)
	if len(s)%2 == 0 {
		left, _ := strconv.Atoi(s[:len(s)/2])
		right, _ := strconv.Atoi(s[len(s)/2:])
		return countStones(left, blinks-1) + countStones(right, blinks-1)
	}

	return countStones(val*2024, blinks-1)
}
