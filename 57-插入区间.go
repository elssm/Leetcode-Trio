package goLeetCode

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	start := make([]int, len(intervals)+1)
	end := make([]int, len(intervals)+1)
	for i, v := range intervals {
		start[i] = v[0]
		end[i] = v[1]
	}
	start[len(intervals)] = newInterval[0]
	end[len(intervals)] = newInterval[1]
	sort.Ints(start)
	sort.Ints(end)
	for i, j := 0, 0; i < len(intervals)+1; i++ {
		if i == len(intervals) || start[i+1] > end[i] {
			res = append(res, []int{start[j], end[i]})
			j = i + 1
		}
	}
	return res

}
