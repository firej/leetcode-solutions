package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maxsum := nums[0]
	currentsum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentsum+nums[i] < nums[i] {
			currentsum = nums[i]
		} else {
			currentsum += nums[i]
		}

		if currentsum > maxsum {
			maxsum = currentsum
		}
	}

	return maxsum
}

func main() {
	fmt.Println("test #1")
	fmt.Println("result:", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println("test #2")
	fmt.Println("result:", maxSubArray([]int{1}))
	fmt.Println("test #3")
	fmt.Println("result:", maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println("test #4")
	fmt.Println("result:", maxSubArray([]int{-1}))
	fmt.Println("test #5")
	fmt.Println("result:", maxSubArray([]int{-2, 1}))
	fmt.Println("test #6")
	fmt.Println("result:", maxSubArray([]int{-2, -1}))
	fmt.Println("test #7")
	fmt.Println("result:", maxSubArray([]int{0, -3, 1, 1}))
}
