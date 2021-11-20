package goLeetCode

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append([]*TreeNode{node.Right}, stack...)
		}
		if node.Left != nil {
			stack = append([]*TreeNode{node.Left}, stack...)
		}
	}
	return res
}
