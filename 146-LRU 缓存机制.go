package goLeetCode

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	key   int
	value int
	pre   *DLinkNode
	next  *DLinkNode
}

func initDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	L := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkNode{},
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	L.head.next = L.tail
	L.tail.pre = L.head
	return L
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.cache[key]; !ok {
		return -1
	}
	node := c.cache[key]
	c.moveToHead(node)
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cache[key]; !ok {
		node := initDLinkNode(key, value)
		c.cache[key] = node
		c.addToHead(node)
		c.size++
		if c.size > c.capacity {
			removed := c.removeTail()
			delete(c.cache, removed.key)
			c.size--
		}
	} else {
		node := c.cache[key]
		node.value = value
		c.moveToHead(node)
	}
}

func (c *LRUCache) addToHead(node *DLinkNode) {
	node.pre = c.head
	node.next = c.head.next
	c.head.next.pre = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *DLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (c *LRUCache) moveToHead(node *DLinkNode) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() *DLinkNode {
	node := c.tail.pre
	c.removeNode(node)
	return node
}
