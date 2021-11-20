package goLeetCode

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for slow != nil && fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			return true
		}
	}
	return false
}
