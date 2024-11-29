package main

import (
	"fmt"
)

// The Hamming distance between two integers is the number of positions
// at which the corresponding bits are different.
// Given two integers x and y, return the Hamming distance between them.

func hammingDistance(x int, y int) int {
	result := 0
	for i := 0; i < 33; i++ {
		if (x&1)^(y&1) > 0 {
			result++
		}
		x = x >> 1
		y = y >> 1
	}
	return result
}

type testcase struct {
	InputX int
	InputY int
	Output int
}

func main() {

	testcases := []testcase{
		{
			InputX: 1,
			InputY: 4,
			Output: 2,
		},
		{
			InputX: 3,
			InputY: 1,
			Output: 1,
		},
	}
	for index, t := range testcases {
		fmt.Printf("test #%d\n", index)
		fmt.Printf("input: X: %d, Y: %d\n", t.InputX, t.InputY)
		r := hammingDistance(t.InputX, t.InputY)
		fmt.Printf("output actual: %d; output expected: %d\n", r, t.Output)
	}
}
