package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		c := []int{}
		for i := 0; i < size; i++ {
			t := q[0]
			q = q[1:]
			c = append(c, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		res = append(res, c)
	}

	return res
}
