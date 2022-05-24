package main

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		n := i
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[n] + nums[l] + nums[r]
			if sum == 0 {
				t := []int{nums[n], nums[l], nums[r]}
				res = append(res, t)
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
