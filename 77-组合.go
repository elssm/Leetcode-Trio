package goLeetCode

var ans [][]int

func combine(n int, k int) [][]int {
	begin := make([]int, 0)

	ans = make([][]int, 0)
	//不要空的子集,所以從1開始
	//dfs(0, begin, k, n)
	dfs(1, begin, k, n)
	return ans
}
func dfs(start int, set []int, setLen int, r int) {

	//限制長度為2,當長度為2時,copy到答案
	if len(set) == setLen {
		temp := make([]int, len(set))

		copy(temp, set)

		ans = append(ans, temp)
	}

	//nums := make([]int, 0)
	//nums = append(nums, r)
	if len(set)+r-start+1 < setLen {
		return
	}

	//回朔算法
	for i := start; i <= r; i++ {

		set = append(set, i)

		dfs(i+1, set, setLen, r)

		set = set[:len(set)-1]
	}

}

func combine2(n int, k int) [][]int {
	nums, res := []int{}, [][]int{}
	findCombine(nums, &res, n, 1, k)
	return res
}

func findCombine(nums []int,res *[][]int, n, index, k int){
	// fmt.Println(nums)
	if len(nums) == k{
		tmp := make([]int, k)
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i:=index; i <= n; i++{
		nums = append(nums, i)
		findCombine(nums, res, n, i+1, k)
		nums = nums[:len(nums)-1]
	}
	return
}