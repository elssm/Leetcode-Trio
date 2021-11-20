package goLeetCode

import "sort"

func merge(intervals [][]int) [][]int {
	var res [][]int
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))
	for i, v := range intervals {
		start[i] = v[0]
		end[i] = v[1]
	}
	sort.Ints(start)
	sort.Ints(end)
	for i, j := 0, 0; i < len(intervals); i++ {
		if i == len(intervals)-1 || start[i+1] > end[i] {
			res = append(res, []int{start[j], end[i]})
			j = i + 1
		}
	}
	return res
}
