package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)
	f := 0
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
		if f == 0 {
			res = append(res, c)
			f = 1
		} else {
			ts := []int{}
			for i := len(c) - 1; i >= 0; i-- {
				ts = append(ts, c[i])
			}
			f = 0
			res = append(res, ts)
		}

	}

	return res
}
