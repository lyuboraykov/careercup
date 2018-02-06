package plusplus

// https://www.careercup.com/question?id=14370695

func plusplus(n []int) []int {
	res := make([]int, len(n))
	copy(res, n)
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = (res[i] + 1) % 10
		if res[i] != 0 {
			break
		}
	}
	if res[0] == 0 {
		res = append([]int{1}, res...)
	}
	return res
}
