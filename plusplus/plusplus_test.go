package plusplus

import "testing"

func TestPlusPlus(t *testing.T) {
	testData := [][][]int{
		{{1}, {2}},
		{{9, 9}, {1, 0, 0}},
		{{1, 1, 9}, {1, 2, 0}},
		{{9, 8}, {9, 9}},
		{{1, 0, 0}, {1, 0, 1}},
	}

	for _, params := range testData {
		expected := params[1]
		got := plusplus(params[0])
		if !areSlicesEqual(expected, got) {
			t.Error("Wrong addition operation", expected, got)
		}
	}
}

func areSlicesEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
