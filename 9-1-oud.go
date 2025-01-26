package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("./9-input.txt")

	var compact []int

	// Pointers.
	head := 0
	tail := len(input) - 1

	headID := 0
	tailID := len(input) / 2
	// tailRest is the amount of files to be moved with the current tailID.
	tailRest := num(input[tail])

outer:
	for {
		// Odd head index: add amount of blocks indicated by head pointer.
		// Prefer adding blocks from the tail rest, so skip this if head ID has caught up with tail ID.
		if headID < tailID {
			for i := num(input[head]); i > 0; i-- {
				compact = append(compact, headID)
			}
		}

		head++

		// Even head index: add blocks from tail side.
		for i := num(input[head]); i > 0; i-- {
			compact = append(compact, tailID)
			tailRest--
			// Add as many blocks as available per file, then move to the next file from tail side.
			if tailRest == 0 {
				tail -= 2
				tailID--
				tailRest = num(input[tail])
				// End when tail ID has caught up with headID, meaning all files have been added.
				if tailID <= headID {
					break outer
				}
			}
		}

		head++
		headID++
	}

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
