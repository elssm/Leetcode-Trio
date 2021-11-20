package goLeetCode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return findIsSymmetric(root.Left, root.Right)

}

func findIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {
		return left.Val == right.Val && findIsSymmetric(left.Left, right.Right) && findIsSymmetric(left.Right, right.Left)
	}
	return false
}
