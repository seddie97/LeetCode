package main

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	c := 0

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res[c][0] = nums[i]
				res[c][1] = nums[l]
				res[c][2] = nums[r]
				c++
			} else if sum < 0 {
				l++
			} else {
				r--
			}
			l++
			r--
		}

	}

	return res
}
