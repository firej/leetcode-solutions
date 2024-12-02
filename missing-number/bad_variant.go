package main

import "fmt"

func find(nums []int) int {
	if len(nums) <= 1 {
		panic(fmt.Sprintf("bad nums! %+v", nums))
	}
	if len(nums) < 3 {
		if nums[0] != nums[1]+1 && nums[0]+1 != nums[1] {
			if nums[1] > nums[0] {
				return nums[0] + 1
			} else {
				return nums[1] + 1
			}
		}
	}
	quantity := len(nums)
	middle := nums[quantity/2]
	left := make([]int, 0, len(nums))
	maxLeft := 0
	missingFromLeft := 0
	right := make([]int, 0, len(nums))
	minRight := middle
	missingFromRight := 0
	for _, v := range nums {
		if v < middle {
			left = append(left, v)
			if maxLeft < v {
				maxLeft = v
			}
		} else if v > middle {
			right = append(right, v)
			if minRight > v {
				minRight = v
			}
		}
	}
	if maxLeft != 0 && (maxLeft+1 != middle) {
		return maxLeft + 1
	} else if len(left) >= 2 {
		missingFromLeft = find(left)
	}

	if minRight != middle && minRight-1 != middle {
		return minRight - 1
	} else if minRight == middle {
		return middle + 1
	} else if len(right) >= 2 {
		missingFromRight = find(right)
	}

	if missingFromLeft != 0 {
		return missingFromLeft
	}
	if missingFromRight != 0 {
		return missingFromRight
	}

	return 0
}

func missingNumberBad(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	r := find(nums)
	if r != 0 {
		return r
	}
	max := nums[0]
	min := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	if min > 0 {
		return 0
	}
	return max + 1
}
