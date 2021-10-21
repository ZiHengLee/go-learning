package bilinknode

type node struct {
	key,val    int
	prev, next *node
}


type linkedLRUCache struct {
	capacity int
	head,tail *node
	cache map[int]*node
}



func LRUCache(capacity int) LRU {
	return &linkedLRUCache{capacity:capacity,cache:make(map[int]*node)}
}


type LRU interface {
	Get(key int) int
	Put(key, value int)
}


func (c *linkedLRUCache) Get(key int) int {
	if node,ok := c.cache[key]; ok {
		c.remove(node)
		c.add(node)
		return node.val
	}
	return -1
}


func (c *linkedLRUCache) Put(key, val int) {
	if n,ok := c.cache[key]; ok {
		c.remove(n)
		n.val = val
		c.add(n)
		return
	}else {
		n = &node{key: key,val:val}
		c.cache[key] = n
		c.add(n)
	}

	if len(c.cache) > c.capacity {
		delete(c.cache,c.tail.key)
		c.remove(c.tail)
	}
}

//add node to head
func (c *linkedLRUCache) add(n *node) {
	if c.head != nil {
		c.head.prev = n
		n.next = c.head
	}

	c.head = n
	if c.tail == nil {
		c.tail = n
		c.tail.prev = n
		c.tail.next = nil
	}
}

//remove node
func (c *linkedLRUCache) remove(n *node) {
	//remove head node
	if c.head == n {
		if n.next != nil {
			n.next.prev = nil
		}
		c.head = n.next
		return
	}
	//remove tail node
	if c.tail == n {
		c.tail = c.tail.prev
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}