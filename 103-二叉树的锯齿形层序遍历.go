package goLeetCode

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	dfsZigzagLevelOrder(root, -1, &res)
	return res
}

func dfsZigzagLevelOrder(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	curLevel := level + 1
	for len(*res) <= curLevel {
		*res = append(*res, []int{})
	}
	if curLevel%2 == 0 {
		(*res)[curLevel] = append((*res)[curLevel], root.Val)
	} else {
		(*res)[curLevel] = append([]int{root.Val}, (*res)[curLevel]...)
	}
	dfsZigzagLevelOrder(root.Left, curLevel, res)
	dfsZigzagLevelOrder(root.Right, curLevel, res)

}
