package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	cur := root
	p := &TreeNode{}
	q := &TreeNode{}
	for i := 2; i <= 10; i++ {
		cur.Left = &TreeNode{Val: i}
		cur.Right = &TreeNode{Val: i + 1}
		cur = cur.Left
		if i == 9 {
			p = cur
		}
	}
	cur = root.Right
	for i := 11; i <= 20; i++ {
		cur.Left = &TreeNode{Val: i}
		cur.Right = &TreeNode{Val: i + 1}
		cur = cur.Right
		if i == 18 {
			q = cur
		}
	}
	fmt.Println(findFN(root, p, q).Val)
}

func findFN(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := findFN(root.Left, p, q)
	right := findFN(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}

	return right
}
