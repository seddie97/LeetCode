package main

type DNode struct {
	Key, Val  int
	Next, Pre *DNode
}

type LRUCache struct {
	size, capacity int
	head           *DNode
	cache          map[int]*DNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		size:     0,
		capacity: capacity,
		head:     &DNode{},
		cache:    map[int]*DNode{},
	}
	lru.head.Pre = lru.head
	lru.head.Next = lru.head
	return lru
}

func (this *DNode) add(node *DNode) {
	node.Next = this.Next
	node.Pre = this
	this.Next.Pre = node
	this.Next = node
}

func (this *DNode) remove() *DNode {
	this.Pre.Next = this.Next
	this.Next.Pre = this.Pre
	return this
}

func (this *LRUCache) Get(key int) int {
	v := 0
	t := this.cache[key]
	if t == nil {
		v = -1
	} else {
		v = t.Val
		t = t.remove()
		this.head.add(t)
	}
	return v
}

func (this *LRUCache) Put(key int, value int) {
	t := this.cache[key]
	if t == nil {
		t = &DNode{
			Key: key,
			Val: value,
		}
		this.size++
	} else {
		t = t.remove()
		delete(this.cache, key)
		t.Val = value
	}
	this.head.add(t)
	this.cache[key] = t
	if this.size > this.capacity {
		this.size--
		rem := this.head.Pre.remove()
		delete(this.cache, rem.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
