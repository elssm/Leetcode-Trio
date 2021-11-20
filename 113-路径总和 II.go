package goLeetCode

func pathSum(root *TreeNode, targetSum int) [][]int {
	tmp, res := []int{}, [][]int{}
	findPath(root, targetSum, tmp, &res)
	return res
}

func findPath(root *TreeNode, targetSum int, tmp []int, res *[][]int) {
	if root == nil {
		return
	}
	tmp = append(tmp, root.Val)
	// fmt.Println(tmp,targetSum)
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		cpt := make([]int, len(tmp))
		copy(cpt, tmp)
		*res = append(*res, cpt)
		return
	}
	findPath(root.Left, targetSum-root.Val, tmp, res)
	findPath(root.Right, targetSum-root.Val, tmp, res)
}
