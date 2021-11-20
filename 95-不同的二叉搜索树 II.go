package goLeetCode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateBSTrees(1, n)
}

func generateBSTrees(start, end int) []*TreeNode {
	var tree []*TreeNode
	if start > end {
		tree = append(tree, nil)
		return tree
	}
	var left []*TreeNode
	var right []*TreeNode
	for i := start; i <= end; i++ {
		left = generateBSTrees(start, i-1)
		right = generateBSTrees(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{i, l, r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}
