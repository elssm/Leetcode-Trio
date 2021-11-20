package goLeetCode

import "math"

func minimumTotal(tri [][]int) int {
	if len(tri) == 0 {
		return 0
	}
	dp, minNum := make([]int, len(tri)), math.MaxInt32
	dp[0] = tri[0][0]
	for i := 1; i < len(tri); i++ {
		for j := len(tri[i]) - 1; j >= 0; j-- {
			// fmt.Println(dp)
			if j == 0 {
				dp[j] += tri[i][0]
			} else if j == len(tri[i])-1 {
				dp[j] = dp[j-1] + tri[i][j]
			} else {
				dp[j] = min(dp[j-1]+tri[i][j], dp[j]+tri[i][j])
			}
		}
	}
	// fmt.Println(dp)
	for _, v := range dp {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}
