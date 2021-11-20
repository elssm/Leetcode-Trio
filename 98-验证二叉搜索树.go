package goLeetCode

func isValidBST(root *TreeNode) bool {
	var res []int
	inOrder(root, &res)
	// fmt.Println(res)
	for i := 1; i < len(res); i++ {
		if res[i-1] >= res[i] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}
