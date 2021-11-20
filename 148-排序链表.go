package goLeetCode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	l2 := slow.Next
	slow.Next = nil
	return merge2(sortList(head), sortList(l2))
}

// 合并有序链表
func merge2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dumyHead := &ListNode{Val: 0}
	lastSorted := dumyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			lastSorted.Next = l1
			l1 = l1.Next
		} else {
			lastSorted.Next = l2
			l2 = l2.Next
		}
		lastSorted = lastSorted.Next
	}
	if l1 != nil {
		lastSorted.Next = l1
	}
	if l2 != nil {
		lastSorted.Next = l2
	}
	return dumyHead.Next
}

func sortList2(head *ListNode) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}

	midNode := middleNode(head)
	cur = midNode.Next
	midNode.Next = nil
	midNode = cur

	left := sortList(head)
	right := sortList(midNode)
	return mergeTwoList(left, right)
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

func mergeTwoList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeTwoList(left.Next, right)
		return left
	}
	right.Next = mergeTwoList(left, right.Next)
	return right
}
