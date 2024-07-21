package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var (
		max   int
		count int
	)

	hash := make(map[byte]bool)
	hash[s[0]] = true
	for i := 0; i < len(s); i++ {
		hash = make(map[byte]bool)
		hash[s[i]] = true
		count = 1

		for j := i + 1; j < len(s); j++ {
			if hash[s[j]] {
				break
			} else {
				hash[s[j]] = true
				count++
			}
		}
		if max < count {
			max = count
		}

	}
	return max
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
