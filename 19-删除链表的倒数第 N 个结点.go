package goLeetCode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var counter, flag = head, head
	var count int = 0
	for counter != nil {
		count++
		counter = counter.Next
	}
	// fmt.Println(count)
	if count-n-1 < 0 {
		return head.Next
	}
	for i := 0; i < count-n-1; i++ {
		flag = flag.Next
		if flag.Next == nil {
			head = head.Next
			return head
		}
	}
	flag.Next = flag.Next.Next
	return head
}
