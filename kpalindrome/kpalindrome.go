package kpalindrome

// https://www.careercup.com/question?id=6287528252407808

import "math"

var cache = make(map[string]int)

func isKpalindrome(s string, k int) bool {
	return numPalindrome(s) <= k
}

func numPalindrome(s string) int {
	if val, ok := cache[s]; ok {
		return val
	}
	n := len(s)
	diffFirst := 0
	diffLast := 0
	for i := range s {
		backI := n - i - 1
		if i >= backI {
			return 0
		}
		if s[i] != s[backI] {
			diffFirst = i
			diffLast = backI
			break
		}
	}
	result := int(math.Min(
		float64(numPalindrome(s[:diffFirst]+s[diffFirst+1:])),
		float64(numPalindrome(s[:diffLast]+s[diffLast+1:])))) + 1
	cache[s] = result
	return result
}
