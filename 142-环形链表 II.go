package goLeetCode

func detectCycle(head *ListNode) *ListNode {
	meetPoint := hasCycle2(head)
	if meetPoint == nil {
		return nil
	}
	for head != nil && meetPoint != nil {
		if head == meetPoint {
			return meetPoint
		}
		head = head.Next
		meetPoint = meetPoint.Next
	}
	return nil
}

func hasCycle2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return fast
		}
	}
	return nil

}
