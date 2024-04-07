package bilinknode

// LRUCache Least Recently Used,最近最少使用,这里用双向链表是方便删除最后一个节点
type LRUCache struct {
	m   map[int]*LinkNode
	cap int
	//前后两个指针不存储实际数据
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{-1, 0, nil, nil}
	tail := &LinkNode{-1, 0, nil, nil}
	head.next = tail
	tail.pre = head
	lru := LRUCache{
		make(map[int]*LinkNode, capacity), capacity, head, tail,
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	m := this.m
	if node, ok := m[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	m := this.m
	cap := this.cap
	tail := this.tail
	if node, ok := m[key]; ok {
		node.val = value
		this.moveToHead(node)
	} else {
		newNode := &LinkNode{key, value, nil, nil}
		if len(m) == cap {
			remNode := tail.pre
			remNode.pre.next = tail
			tail.pre = remNode.pre
			remNode.next = nil
			remNode.pre = nil
			delete(m, remNode.key)
		}
		m[key] = newNode
		this.moveToHead(newNode)
	}
}

func (this *LRUCache) moveToHead(node *LinkNode) {
	if node.next != nil && node.pre != nil {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}
