package kpalindrome

import "testing"

var palindromeTests = []struct {
	s   string
	k   int
	res bool
}{
	{"abxa", 1, true},
	{"abxda", 1, false},
	{"abccba", 0, true},
}

func TestIsKPalindrome(t *testing.T) {
	for _, example := range palindromeTests {
		if isKpalindrome(example.s, example.k) != example.res {
			t.Error("Wrong result on isKPalindrome", example.s, example.k, example.res)
		}
	}
}
