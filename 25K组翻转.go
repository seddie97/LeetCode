package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	res := &ListNode{
		Next: head,
		Val:  0,
	}
	pre := res
	cur := head
	len := 0

	for cur != nil {
		len++
		cur = cur.Next
	}
	cur = head
	for i := 0; i < len/k; i++ {
		t := k
		for t > 1 {
			t--
			nx := cur.Next
			cur.Next = nx.Next
			nx.Next = pre.Next
			pre.Next = nx
		}
		pre = cur
		cur = cur.Next
	}

	return res.Next
}
