package goLeetCode

func longestConsecutive(nums []int) int {
	ans := 0
	hs := map[int]bool{}

	for _, v := range nums {
		hs[v] = true
	}

	for num := range hs {
		cnt := 1
		if !hs[num-1] {
			for hs[num+cnt] {
				cnt++
			}
			if cnt > ans {
				ans = cnt
			}
		}
	}

	return ans
}
