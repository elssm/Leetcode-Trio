package goLeetCode

func largestRectangleArea(heights []int) int {
	var stack []int
	heights = append(heights, 0)
	maxValue := 0
	for index, value := range heights {
		for len(stack) != 0 && heights[stack[len(stack)-1]] >= value {
			cur := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				maxValue = max(maxValue, index*cur)
			} else {
				maxValue = max(maxValue, (index-stack[len(stack)-1]-1)*cur)
			}
		}
		stack = append(stack, index)
	}
	return maxValue
}
