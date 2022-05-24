package main

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum = max(nums[i], nums[i]+sum)
		res = max(res, sum)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
