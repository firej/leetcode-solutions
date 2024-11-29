package main

import (
	"fmt"
)

// Whitespace: Ignore any leading whitespace (" ").
// Signedness: Determine the sign by checking if the next
//		character is '-' or '+', assuming positivity if neither present.
// Conversion: Read the integer by skipping leading zeros until
//		a non-digit character is encountered or the end of the
//		string is reached. If no digits were read, then the result is 0.
// Rounding: If the integer is out of the 32-bit signed integer range
//		[-231, 231 - 1], then round the integer to remain in the range.
//		Specifically, integers less than -231 should be rounded to -231,
//		and integers greater than 231 - 1 should be rounded to 231 - 1.

func myAtoi(s string) int {
	const MaxUint = ^uint32(0)
	const MaxInt = int32(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	var accumulator int32
	var started bool
	var isMinus bool
	for i, _ := range s {
		if '0' <= s[i] && s[i] <= '9' {
			buf := int(accumulator)*10 + int(s[i]-'0')
			if buf > int(MaxInt) {
				if isMinus {
					return int(MinInt)
				} else {
					return int(MaxInt)
				}
			}
			accumulator = int32(buf)
			started = true
		} else if s[i] == '-' {
			if !started {
				isMinus = true
				started = true
			} else {
				break
			}
		} else if s[i] == '+' {
			if started {
				break
			}
			started = true
		} else if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if started {
				break
			}
		} else {
			break
		}
	}
	if isMinus {
		return int(-accumulator)
	} else {
		return int(accumulator)
	}
}

func main() {
	fmt.Println("test #1")
	fmt.Println("result:", myAtoi("42"))
	fmt.Println("test #2")
	fmt.Println("result:", myAtoi(" -042"))
	fmt.Println("test #3")
	fmt.Println("result:", myAtoi("1337c0d3"))
	fmt.Println("test #4")
	fmt.Println("result:", myAtoi("0-1"))
	fmt.Println("test #5")
	fmt.Println("result:", myAtoi("-91283472332"))

}
