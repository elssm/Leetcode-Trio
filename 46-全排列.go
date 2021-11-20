package goLeetCode

func permute(nums []int) [][]int {
	var res [][]int
	var temp []int
	mark := make([]bool, len(nums))
	findPermute(nums, &res, temp, mark)
	return res
}

func findPermute(nums []int, res *[][]int, temp []int, mark []bool) {
	if len(temp) == len(nums) {
		tt := make([]int, len(temp))
		copy(tt, temp)
		*res = append(*res, tt)
		return
	}
	for i := 0; i < len(nums); i++ {
		if mark[i] != true {
			temp = append(temp, nums[i])
			mark[i] = true
			findPermute(nums, res, temp, mark)
			temp = temp[:len(temp)-1]
			mark[i] = false
		}
	}
	return
}
