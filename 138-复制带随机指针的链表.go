package goLeetCode

// NodeRanDom Definition for a Node.
type NodeRanDom struct {
	Val    int
	Next   *NodeRanDom
	Random *NodeRanDom
}

func copyRandomList(head *NodeRanDom) *NodeRanDom {
	if head == nil {
		return nil
	}
	queue := []*NodeRanDom{head}
	visited := map[*NodeRanDom]*NodeRanDom{}
	visited[head] = &NodeRanDom{head.Val, nil, nil}
	for len(queue) > 0 {
		h := queue[0]
		queue = queue[1:]
		if h.Next != nil {
			if _, ok := visited[h.Next]; !ok {
				visited[h.Next] = &NodeRanDom{h.Next.Val, nil, nil}
				queue = append(queue, h.Next)
			}
			visited[h].Next = visited[h.Next]
		}
		if h.Random != nil {
			if _, ok := visited[h.Random]; !ok {
				visited[h.Random] = &NodeRanDom{h.Random.Val, nil, nil}
				queue = append(queue, h.Random)
			}
			visited[h].Random = visited[h.Random]
		}

	}
	return visited[head]
}
