package exceptfortwo

import "testing"

func TestGetTwoOddNumbers(t *testing.T) {
	res := getTwoOddNumbers([]int{1, 1, 2, 3, 3, 4})
	if res[0] != 2 && res[1] != 4 {
		t.Error("Wrong odd numbers", res)
	}
}
