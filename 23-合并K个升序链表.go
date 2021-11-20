package goLeetCode

func mergeKLists(lists []*ListNode) *ListNode {
	var length = len(lists)
	if length == 0{
		return nil
	}
	if length == 1{
		return lists[0]
	}
	mid := length >> 1
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists1(left *ListNode, right *ListNode) *ListNode{
	if left == nil{
		return right
	}
	if right == nil{
		return left
	}
	if left.Val < right.Val{
		left.Next = mergeTwoLists(left.Next, right)
		return left
	}
	right.Next = mergeTwoLists(left, right.Next)
	return right
}
