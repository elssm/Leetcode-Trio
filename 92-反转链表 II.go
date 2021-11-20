package goLeetCode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{0, head}
	pre := newHead
	for i := 0; pre.Next != nil && i < left-1; i++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		temp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = temp
	}
	return newHead.Next
}
