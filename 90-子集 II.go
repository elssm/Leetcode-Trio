package goLeetCode

import "sort"

var res [][]int

func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)
	res = [][]int{}
	var track []int

	backTrack(nums, 0, track)
	return res
}
func backTrack(nums []int, start int, track []int) {

	temp := make([]int, len(track))
	copy(temp, track)
	res = append(res, temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backTrack(nums, i+1, track)
		track = track[:len(track)-1]

	}
}
