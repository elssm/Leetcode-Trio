package goLeetCode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	leftVal, rightVal := maxNormSum(root.Left), maxNormSum(root.Right)
	// fmt.Println(root.Val, leftVal, rightVal)
	// fmt.Println(leftVal, rightVal)
	if root.Right == nil && root.Left != nil {
		return max(max(root.Val, root.Val+leftVal), maxPathSum(root.Left))
	}
	if root.Left == nil && root.Right != nil {
		return max(max(root.Val, root.Val+rightVal), maxPathSum(root.Right))
	}
	if root.Left != nil && root.Right != nil {
		return maxOfThree(max(root.Val, maxOfThree(root.Val+leftVal, root.Val+rightVal, root.Val+leftVal+rightVal)), maxPathSum(root.Left), maxPathSum(root.Right))
	}
	return 0
}

func maxNormSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(root.Val, root.Val+max(maxNormSum(root.Left), maxNormSum(root.Right)))
}

func maxOfThree(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
