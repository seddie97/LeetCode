package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	isValidBST(&TreeNode{
		Val: 1,
	})
}

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	cur := root
	last := ^int(^uint(0) >> 1)
	fmt.Println(last)
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Val <= last {
			return false
		}
		last = cur.Val
		cur = cur.Right
	}

	return true
}
