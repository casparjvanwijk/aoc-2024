package main

import (
	"fmt"
	"os"
	"strconv"
)

const free = -1

func main() {
	input, _ := os.ReadFile("./9-input.txt")

	// Parse input blocks.
	var blocks []int
	for i, byte := range input {
		// Append free space on odd indexes and incrementing file ID blocks on even indexes.
		val := free
		if i%2 == 0 {
			val = i / 2
		}
		for count := num(byte); count > 0; count-- {
			blocks = append(blocks, val)
		}
	}

	// Pick file blocks from left pointer.
	// When hitting free space, pick file block from right pointer.
	var compact []int
	j := len(blocks) - 1
	for i := 0; i <= tail; i++ {
		if blocks[i] == free {
			compact = append(compact, blocks[j])
			j--
			// Tail pointer skips free space.
			for blocks[j] == free {
				j--
			}
			continue
		}
		compact = append(compact, blocks[i])
	}

	// Calculate checksum.
	checksum := 0
	for i, block := range compact {
		checksum += i * block
	}
	fmt.Println(checksum)
}

func num(b byte) int {
	result, _ := strconv.Atoi(string(b))
	return result
}
