package main

var res [][]int
var v []int
var t []int

func permute(nums []int) [][]int {
	res = [][]int{}
	for i := 0; i < len(nums); i++ {
		v = append(v, 0)
	}

	findRes(nums)
	return res
}

func findRes(nums []int) {
	if len(t) == len(nums) {
		p := make([]int, len(t))
		copy(p, t)
		res = append(res, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		if v[i] == 1 {
			continue
		}
		v[i] = 1
		t = append(t, nums[i])
		findRes(nums)
		t = t[:len(t)-1]
		v[i] = 0
	}
}
