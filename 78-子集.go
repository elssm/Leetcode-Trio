package goLeetCode

func subsets(nums []int) [][]int {
	var set []int
	var res [][]int
	findSubsets(nums, set, &res, 0)
	return res
}

func findSubsets(nums, set []int, res *[][]int, index int) {
	tmp := make([]int, len(set))
	copy(tmp, set)
	*res = append(*res, tmp)
	for i := index; i < len(nums); i++ {
		set = append(set, nums[i])
		findSubsets(nums, set, res, i+1)
		set = set[:len(set)-1]
	}
	return
}
