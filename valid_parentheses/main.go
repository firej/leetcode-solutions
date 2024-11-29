package main

import (
	"fmt"
)

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
//
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// 3. Every close bracket has a corresponding open bracket of the same type.

type stack []rune

func (s *stack) push(letter rune) {
	*s = append(*s, letter)
}
func (s *stack) pop() rune {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func isValid(s string) bool {
	validator := stack{}
	for _, next_bracket := range s {
		if next_bracket == '(' || next_bracket == '[' || next_bracket == '{' {
			validator.push(next_bracket)
		} else if next_bracket == ')' || next_bracket == ']' || next_bracket == '}' {
			if len(validator) == 0 {
				return false
			}
			opened_bracket := validator.pop()
			if next_bracket == ')' && opened_bracket != '(' {
				return false
			} else if next_bracket == ']' && opened_bracket != '[' {
				return false
			} else if next_bracket == '}' && opened_bracket != '{' {
				return false
			}
		} else {
			return false
		}
	}
	return len(validator) == 0
}

type testcase struct {
	Input  string
	Output bool
}

func main() {

	testcases := []testcase{
		{
			Input:  "()",
			Output: true,
		},

		{
			Input:  "()[]{}",
			Output: true,
		},

		{
			Input:  "(]",
			Output: false,
		},
		{
			Input:  "([])",
			Output: true,
		},
	}
	for index, t := range testcases {
		fmt.Printf("test #%d\n", index)
		fmt.Printf("input: %s\n", t.Input)
		r := isValid(t.Input)
		fmt.Printf("output actual: %t; output expected: %t\n", r, t.Output)
	}
}
