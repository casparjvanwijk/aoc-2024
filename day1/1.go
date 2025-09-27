package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./1-input.txt")
	scanner := bufio.NewScanner(file)

	// Split input into two lists.
	list1 := []int{}
	list2 := []int{}
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), "   ")
		val1, _ := strconv.Atoi(spl[0])
		val2, _ := strconv.Atoi(spl[1])
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	// Sum differences.
	dSum := 0
	for i := range list1 {
		dSum += abs(list1[i] - list2[i])
	}
	fmt.Println(dSum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
