package goLeetCode

func partition1(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{-202, head}
	cur, last := head, newHead
	for cur.Next != nil {
		if cur.Val >= x {
			break
		}
		cur = cur.Next
		last = last.Next
	}
	if cur.Next == nil {
		return head
	}
	slow := last
	// fmt.Println(last.Val, cur.Val)
	for cur.Next != nil {
		// fmt.Println(cur.Val)
		if cur.Val < x {
			slow.Next = cur.Next
			temp := last.Next
			last.Next = cur
			cur.Next = temp
			cur = slow.Next
			last = last.Next
		} else {
			slow = slow.Next
			cur = cur.Next
		}

	}
	if cur.Val < x {
		temp := last.Next
		last.Next = cur
		cur.Next = temp
		slow.Next = nil
	}
	return newHead.Next
}
