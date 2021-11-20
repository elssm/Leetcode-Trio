package goLeetCode

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	var cur *TreeNode = root
	var p *TreeNode = nil
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append([]*TreeNode{cur}, stack...)
			cur = cur.Left
		}
		cur = stack[0]
		if cur.Right == nil || cur.Right == p {
			res = append(res, cur.Val)
			stack = stack[1:]
			p = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return res
}
