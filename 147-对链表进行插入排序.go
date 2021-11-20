package goLeetCode

import "sort"

type mapListNode struct {
	value int
	node  *ListNode
}

type ListNodeArray []*mapListNode

func (a ListNodeArray) Len() int {
	return len(a)
}

func (a ListNodeArray) Less(i, j int) bool {
	return (a)[i].value < (a)[j].value
}

func (a ListNodeArray) Swap(i, j int) {
	(a)[i], (a)[j] = (a)[j], (a)[i]
}

func insertionSortList(head *ListNode) *ListNode {
	LArray := ListNodeArray{}
	tmphead := head
	for tmphead != nil {
		LArray = append(LArray, &mapListNode{tmphead.Val, tmphead})
		tmphead = tmphead.Next
	}
	sort.Sort(LArray)
	// fmt.Println(LArray)
	newHead := &ListNode{0, nil}
	cur := newHead
	for _, v := range LArray {
		cur.Next = (*v).node
		cur = (*v).node
	}
	cur.Next = nil
	return newHead.Next
}

func insertionSortList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	hint := &ListNode{0, head}

	sorted, tosort := head, head.Next
	for tosort != nil {
		if sorted.Val <= tosort.Val {
			sorted = sorted.Next
		} else {
			// 1 -> 2 -> 4 -> 5 - > 3
			// pre: 找到小于tosort的右边界；可见tosort不一定刚好插在sorted的前面！
			pre := hint
			for pre.Next.Val <= tosort.Val {
				pre = pre.Next
			}

			pre.Next, tosort.Next, sorted.Next = tosort, pre.Next, tosort.Next
		}

		tosort = sorted.Next // 注意；经历一轮插入后，tosort可能会到前面去 - 但总归是要在sorted后面的
	}

	return hint.Next
}
