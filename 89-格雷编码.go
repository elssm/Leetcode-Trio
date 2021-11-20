package goLeetCode

func grayCode(n int) []int {
	var ans []int
	powN := 1 << n
	for i := 0; i < powN; i++ {
		ans = append(ans, i^(i>>1))
	}
	return ans
}
