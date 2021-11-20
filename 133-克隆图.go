package goLeetCode

type NodeList struct {
	Val  int
	Neighbors []*NodeList
}


func cloneGraph(Node *NodeList) *NodeList {
	visited := map[*NodeList]*NodeList{}
	var cg func(Node *NodeList) *NodeList
	cg = func(Node *NodeList) *NodeList {
		if Node == nil {
			return Node
		}

		// 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回
		if _, ok := visited[Node]; ok {
			return visited[Node]
		}

		// 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
		cloneNode := &NodeList{Node.Val, []*NodeList{}}
		// 哈希表存储
		visited[Node] = cloneNode

		// 遍历该节点的邻居并更新克隆节点的邻居列表
		for _, n := range Node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
		}
		return cloneNode
	}
	return cg(Node)
}
