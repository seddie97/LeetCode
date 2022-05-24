package main

type MyQueue struct {
	Length int
	Pre    *Value
	Suf    *Value
}

type Value struct {
	Val  int
	Next *Value
}

/** Initialize your data structure here. */
func Constructor3() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	val := &Value{Val: x}

	if this.Suf != nil {
		this.Suf.Next = val
		this.Suf = val
	} else {
		this.Suf = val
	}

	if this.Pre == nil {
		this.Pre = val
	}

	this.Length++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	val := this.Pre
	this.Pre = val.Next
	this.Length--

	return val.Val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Pre.Val
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.Length == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
