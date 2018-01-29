package smallestrangeintegers

import "testing"

func TestFindSmallestRange(t *testing.T) {
	got := findSmallestRange([][]int{
		[]int{4, 10, 15, 24, 26},
		[]int{0, 9, 12, 20},
		[]int{5, 18, 22, 30},
	})
	expected := [2]int{20, 24}
	if got != expected {
		t.Error("got different from expected", got, expected)
	}
}
