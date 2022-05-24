package main

type DNode struct {
	key, val  int
	pre, next *DNode
}

type LRUCache struct {
	size, cap int
	cache     map[int]*DNode
	head      *DNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:  0,
		cap:   capacity,
		cache: map[int]*DNode{},
		head:  &DNode{},
	}
	l.head.pre = l.head
	l.head.next = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}

}

func (this *LRUCache) Put(key int, value int) {

}

func (this *LRUCache) addToHead(node *DNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DNode) {

}

func main() {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
