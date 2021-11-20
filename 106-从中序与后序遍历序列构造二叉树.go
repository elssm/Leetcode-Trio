package goLeetCode

func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	value := postorder[len(postorder)-1]
	root := &TreeNode{value, nil, nil}
	pos := findPos(value, inorder)
	if pos != -1 {
		root.Left = buildTree1(inorder[:pos], postorder[:pos])
		root.Right = buildTree1(inorder[pos+1:], postorder[pos:len(postorder)-1])
	}
	return root
}

func findPos(a int, inorder []int) int {
	for i, v := range inorder {
		if a == v {
			return i
		}
	}
	return -1
}
