package goLeetCode

func sumNumbers(root *TreeNode) int {
	res := 0
	findSumNumbers(root, 0, &res)
	return res
}

func findSumNumbers(root *TreeNode, tmp int, res *int) {
	if root.Left == nil && root.Right == nil {
		*res += tmp*10 + root.Val
		return
	}
	if root.Left != nil {
		findSumNumbers(root.Left, tmp*10+root.Val, res)
	}
	if root.Right != nil {
		findSumNumbers(root.Right, tmp*10+root.Val, res)
	}
}
