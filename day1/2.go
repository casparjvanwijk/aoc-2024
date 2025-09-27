package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./1-input.txt")
	scanner := bufio.NewScanner(file)

	// Split input into left list and counts of values in right list.
	list1 := []int{}
	list2Counts := make(map[int]int)
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), "   ")
		val1, _ := strconv.Atoi(spl[0])
		val2, _ := strconv.Atoi(spl[1])
		list1 = append(list1, val1)
		list2Counts[val2]++
	}

	result := 0
	for _, n := range list1 {
		result += n * list2Counts[n]
	}
	fmt.Println(result)
}
