package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	i := 0
	for {
		for _, s := range strs {
			if i >= len(s) {
				return strs[0][:i]
			}
			if s[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
		i++
	}
}

func main() {
	fmt.Println("test #1")
	res1 := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println("result:", res1)
	fmt.Println("test #2")
	res2 := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Println("result:", res2)

}
