package goLeetCode

func connect1(root *Node) *Node {
	travel(root)
	return root
}

func travel1(root *Node) {
	if root == nil {
		return
	}
	var pre, cur *Node = root, nil
	for getChild(pre) != nil {
		cur = pre
		for cur != nil {
			child := cur
			if cur.Right != nil && cur.Left != nil {
				cur.Left.Next = cur.Right
				child = cur.Right
			} else {
				child = getChild(cur)
			}
			if child != cur && child != nil {
				child.Next = getChild(cur.Next)
			}
			cur = cur.Next
		}
		pre = getChild(pre)
	}

}

func getChild(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return root.Left
	} else if root.Right != nil {
		return root.Right
	} else {
		return getChild(root.Next)
	}

}
