package main

import (
	"fmt"
	"sort"
)

// Given an array nums containing n distinct numbers
// in the range [0, n], return the only number in the
// range that is missing from the array.

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	}

	for i, _ := range nums {
		if i < len(nums)-1 {
			if nums[i]+1 != nums[i+1] {
				return nums[i] + 1
			}
		} else {
			return nums[i] + 1
		}
	}
	return 0
}

type testcase struct {
	Input  []int
	Output int
}

func main() {

	testcases := []testcase{
		{
			Input:  []int{3, 0, 1},
			Output: 2,
		},
		{

			Input:  []int{0, 1},
			Output: 2,
		},
		{

			Input:  []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			Output: 8,
		},
		{
			Input:  []int{},
			Output: 0,
		},
		{
			Input:  []int{1, 2},
			Output: 0,
		},
		{
			Input:  []int{0, 1, 3},
			Output: 2,
		},
	}
	for index, t := range testcases {
		fmt.Printf("test #%d\n", index)
		fmt.Printf("input: %d\n", t.Input)
		r := missingNumber(t.Input)
		fmt.Printf("output actual: %d; output expected: %d\n", r, t.Output)
	}
}
