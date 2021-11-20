package goLeetCode

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	number := make([]byte, n)
	var ans []byte
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
		number[i-1] = '0' + byte(i)
	}
	k--
	// 从高位到低位进行递归
	for i := 1; i <= n; i++ {
		index := k / factorial[n-i]
		ans = append(ans, number[index])
		k -= index * factorial[n-i]
		number = append(number[:index], number[index+1:]...)
	}
	return string(ans)
}
