package goLeetCode

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	preMid := slow
	preCur := slow.Next
	for preCur.Next != nil {
		cur := preCur.Next
		preCur.Next = cur.Next
		cur.Next = preMid.Next
		preMid.Next = cur
	}
	slow = head
	fast = preMid.Next
	for slow != preMid {
		preMid.Next = fast.Next
		fast.Next = slow.Next
		slow.Next = fast
		slow = fast.Next
		fast = preMid.Next
	}
}
