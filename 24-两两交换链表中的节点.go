package goLeetCode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var lst *ListNode = new(ListNode)
	var res = lst

	for head != nil && head.Next != nil {
		ne := head.Next
		temp := ne.Next
		head.Next = temp
		ne.Next = head
		lst.Next = ne
		lst = head
		head = temp
	}
	return res.Next
}
