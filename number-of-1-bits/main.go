package main

import (
	"fmt"
)

// Given a positive integer n, write a function that returns
// the number of set bits in its binary representation
// (also known as the Hamming weight).

func hammingWeight(n int) int {
	result := 0
	for n > 0 {
		if n&1 == 1 {
			result++
		}
		n = n >> 1
	}
	return result
}

type testcase struct {
	Input  int
	Output int
}

func main() {

	testcases := []testcase{
		{
			Input:  11,
			Output: 3,
		},
		{
			Input:  128,
			Output: 1,
		},
		{
			Input:  2147483645,
			Output: 30,
		},
	}
	for index, t := range testcases {
		fmt.Printf("test #%d\n", index)
		fmt.Printf("input: %d\n", t.Input)
		r := hammingWeight(t.Input)
		fmt.Printf("output actual: %d; output expected: %d\n", r, t.Output)
	}
}
