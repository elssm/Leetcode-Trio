package goLeetCode

func candy(ratings []int) int {
	res := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res[i] = max(res[i], res[i+1]+1)
		}
	}
	count := len(res)
	for i := 0; i < len(res); i++ {
		count += res[i]
	}
	return count
}
