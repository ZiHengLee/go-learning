package bilinknode

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode // 指向哈希表的指针
	cap        int               // 长度
	head, tail *LinkNode         // 两个哨兵
}

func Constructor(capacity int) LRUCache {
	//采用虚拟头尾节点，这两个节点始终处于首位，但不计入map中
	head := &LinkNode{-1, 0, nil, nil}
	tail := &LinkNode{-1, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) Get(key int) int {
	m := this.m
	if node, ok := m[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) moveToHead(node *LinkNode) {
	if node.pre != nil && node.next != nil{
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) Put(key int, value int) {
	m := this.m
	cap := this.cap
	tail := this.tail
	if node, ok := m[key]; ok {
		//插入新的值
		node.val = value
		this.moveToHead(node)
	} else {
		newNode := &LinkNode{key, value, nil, nil}
		if len(m) == cap {
			rmTail := tail.pre
			rmTail.pre.next = tail
			tail.pre =rmTail.pre
			rmTail.next = nil
			delete(m, rmTail.key)
		}
		this.moveToHead(newNode)
		m[key] = newNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */