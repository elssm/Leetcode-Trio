package goLeetCode

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	cols := len(matrix[0])
	dp := make([]int, cols)
	ans := 0
	for i := range matrix {
		for j, value := range matrix[i] {
			if value == '1' {
				dp[j]++
			} else {
				dp[j] = 0
			}
		}
		ans = max(ans, LargeRectangle(dp))
	}
	return ans
}

func LargeRectangle(list []int) int {
	stack := make([]int, 0)
	res := 0
	list = append(append([]int{-1}, list...), -1)
	for i := range list {
		for len(stack) > 0 && list[stack[len(stack)-1]] > list[i] {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = max(res, list[temp]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return res
}


