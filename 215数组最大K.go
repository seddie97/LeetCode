package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	return qSort(nums, 0, len(nums)-1, len(nums)-k)
}

func qSort(nums []int, l int, r int, k int) int {
	q := rSort(nums, l, r)
	if q == k {
		return nums[q]
	} else if q < k {
		return qSort(nums, q+1, r, k)
	} else {
		return qSort(nums, l, q-1, k)
	}
}

func rSort(nums []int, l int, r int) int {
	q := rand.Intn(r-l+1) + l
	sw(nums, q, l)
	return sort(nums, l, r)
}

func sort(nums []int, l int, r int) int {
	q := l
	for l < r {
		for l < r && nums[r] > nums[q] {
			r--
		}
		for l < r && nums[l] <= nums[q] {
			l++
		}
		if l < r {
			sw(nums, l, r)
		}
	}
	sw(nums, l, q)
	return l
}

func sw(nums []int, l int, r int) {
	t := nums[l]
	nums[l] = nums[r]
	nums[r] = t
}
