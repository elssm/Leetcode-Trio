package goLeetCode

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1] += 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] >= 10 {
			if i == 0 {
				digits[i] -= 10
				digits = append([]int{1}, digits...)
				break
			}
			digits[i-1] += 1
			digits[i] -= 10
		}
	}
	return digits
}
