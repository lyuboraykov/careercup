package exceptfortwo

// https://www.careercup.com/question?id=16306671

func getTwoOddNumbers(numbers []int) []int {
	oddMap := make(map[int]bool)
	for _, n := range numbers {
		if _, ok := oddMap[n]; ok {
			delete(oddMap, n)
		} else {
			oddMap[n] = true
		}
	}
	result := make([]int, 0)
	for n := range oddMap {
		result = append(result, n)
	}
	return result
}
