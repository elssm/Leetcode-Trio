package goLeetCode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid, last := findListMid(head)
	if mid == head {
		head = nil
	}
	root := &TreeNode{mid.Val, nil, nil}
	nextHead := mid.Next
	last.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(nextHead)
	return root
}

func findListMid(head *ListNode) (*ListNode, *ListNode) {
	last := &ListNode{0, head}
	cur, fast := head, head
	for fast.Next != nil {
		cur = cur.Next
		last = last.Next
		fast = fast.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next
	}
	return cur, last
}
