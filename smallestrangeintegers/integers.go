package smallestrangeintegers

type candidate struct {
	index int
	value int
}

// https://www.careercup.com/question?id=16759664
func findSmallestRange(numbers [][]int) [2]int {
	k := len(numbers)
	candidates := make([]candidate, k)
	minCandidateI := 0
	maxCandidateI := 0
	for i, v := range numbers {
		candidates[i] = candidate{
			0,
			v[0],
		}
		if candidates[i].value < candidates[minCandidateI].value {
			minCandidateI = i
		}
		if candidates[i].value > candidates[maxCandidateI].value {
			maxCandidateI = i
		}
	}

	minCandidate := &candidates[minCandidateI]
	maxCandidate := &candidates[maxCandidateI]
	minRange := maxCandidate.value - minCandidate.value
	minRangeValues := [2]int{minCandidate.value, maxCandidate.value}
	for {
		minCandidate.index++
		if minCandidate.index < len(numbers[minCandidateI]) {
			minCandidate.value = numbers[minCandidateI][minCandidate.index]
		} else {
			break
		}
		for i, c := range candidates {
			if c.value < minCandidate.value {
				minCandidate = &candidates[i]
				minCandidateI = i
			}
			if c.value > maxCandidate.value {
				maxCandidate = &candidates[i]
				maxCandidateI = i
			}
		}
		newMinRange := maxCandidate.value - minCandidate.value
		if newMinRange < minRange {
			minRange = newMinRange
			minRangeValues = [2]int{minCandidate.value, maxCandidate.value}
		}
	}

	return minRangeValues
}
