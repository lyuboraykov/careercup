package main

import "fmt"

type trueValue struct{}
type set map[int]trueValue

var t = trueValue{}

func main() {
	fmt.Println(criticalConnections(4, [][]int{
		[]int{0, 1}, []int{1, 2}, []int{2, 0}, []int{1, 3},
	}))
}

func criticalConnections(n int, connections [][]int) [][]int {
	result := make([][]int, 0)

	verticesSet := make(set)
	vertices := make([]int, 0)
	for _, c := range connections {
		if _, ok := verticesSet[c[0]]; !ok {
			vertices = append(vertices, c[0])
			verticesSet[c[0]] = t
		}
		if _, ok := verticesSet[c[1]]; !ok {
			vertices = append(vertices, c[1])
			verticesSet[c[1]] = t
		}
	}

	for i := range connections {
		copyConns := make([][]int, len(connections))
		copy(copyConns, connections)
		var connectionsWithoutCurrent [][]int
		if i == 0 {
			connectionsWithoutCurrent = copyConns[1:]
		} else if i == len(connections)-1 {
			connectionsWithoutCurrent = copyConns[:i]
		} else {
			connectionsWithoutCurrent = append(copyConns[:i], copyConns[i+1:]...)
		}

		if countComponents(connectionsWithoutCurrent, vertices) > 1 {
			result = append(result, connections[i])
		}
	}

	return result
}

func countComponents(connections [][]int, vertices []int) int {
	sets := make([]set, len(vertices))

	for i, v := range vertices {
		sets[i] = set{v: t}
	}

	for _, c := range connections {
		var s1, s2 set
		var i1, i2 int

		bothInSame := false
		for i, s := range sets {
			_, ok1 := s[c[0]]
			_, ok2 := s[c[1]]

			if ok1 && ok2 {
				bothInSame = true
				break
			}

			if ok1 {
				i1 = i
				s1 = s
			}
			if ok2 {
				i2 = i
				s2 = s
			}
			if s1 != nil && s2 != nil {
				break
			}
		}

		if bothInSame {
			// already connected, not essential connection
			continue
		}

		if s1 != nil && s2 != nil {
			jointSet := make(set)
			for k, v := range s1 {
				jointSet[k] = v
			}
			for k, v := range s2 {
				jointSet[k] = v
			}
			newSets := make([]set, 0, len(sets)-1)
			for i, s := range sets {
				if i != i1 && i != i2 {
					newSets = append(newSets, s)
				}
			}
			newSets = append(newSets, jointSet)
			sets = newSets
		} else if s1 != nil {
			s1[c[1]] = t
		} else if s2 != nil {
			s2[c[0]] = t
		} else {
			newSet := make(set)
			newSet[c[0]] = t
			newSet[c[1]] = t
			sets = append(sets, newSet)
		}

	}

	return len(sets)
}
