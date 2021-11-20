package goLeetCode

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next != nil && head.Next.Val == head.Val {
		for head.Next != nil && head.Next.Val == head.Val {
			head.Next = head.Next.Next
		}
	}
	head.Next = deleteDuplicates1(head.Next)
	return head
}
