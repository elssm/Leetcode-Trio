package goLeetCode

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	dfsLevelOrder(root, -1, &res)
	return res
}

func dfsLevelOrder(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	curLevel := level + 1
	for len(*res) <= curLevel {
		*res = append(*res, []int{})
	}
	(*res)[curLevel] = append((*res)[curLevel], root.Val)
	dfsLevelOrder(root.Left, curLevel, res)
	dfsLevelOrder(root.Right, curLevel, res)
}
