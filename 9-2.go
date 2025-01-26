package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const free = -1

func main() {
	input, _ := os.ReadFile("./9-input.txt")

	// Parse input blocks.
	var blocks []int
	for i, byte := range input {
		// Append free space on odd index and incrementing file ID block on even index.
		val := free
		if i%2 == 0 {
			val = i / 2
		}
		for count := num(byte); count > 0; count-- {
			blocks = append(blocks, val)
		}
	}

	// Make a copy so that files cannot be moved twice.
	compact := make([]int, len(blocks))
	copy(compact, blocks)

	// Make compact: loop over blocks from right to left until pointer reaches file 0.
	for i := len(blocks) - 1; blocks[i] != 0; i-- {
		// Skip empty blocks.
		if blocks[i] == free {
			continue
		}

		fileEnd := i
		fileSize := 0
		for blocks[fileEnd-fileSize] == blocks[fileEnd] {
			fileSize++
		}
		fileStart := fileEnd - fileSize + 1

		// Find end index of first free space large enough for file. -1 means that no space is found.
		spaceEnd := -1
		spaceSize := 0
		for j, block := range compact[:fileStart] {
			if block == free {
				spaceSize++
				if spaceSize == fileSize {
					spaceEnd = j
					break
				}
				continue
			}
			spaceSize = 0
		}

		// If space is found, move the file.
		if spaceEnd != -1 {
			for j := 0; j < fileSize; j++ {
				compact[fileEnd-j], compact[spaceEnd-j] = compact[spaceEnd-j], compact[fileEnd-j]
			}
		}

		// Skip rest of file.
		i = fileStart
	}

	// Calculate checksum.
	checksum := 0
	for i, block := range compact {
		if block == -1 {
			continue
		}
		checksum += i * block
	}
	fmt.Println(checksum)
}

func num(b byte) int {
	result, _ := strconv.Atoi(string(b))
	return result
}

func printBlocks(blocks []int) {
	var b strings.Builder
	for _, block := range blocks {
		switch {
		case block == -1:
			b.WriteString(".")
		default:
			b.WriteString(strconv.Itoa(block))
		}
	}
	fmt.Println(b.String())
}
