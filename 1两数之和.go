package main

func twoSum(nums []int, target int) []int {
	res := []int{}
	sum := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		v, ok := sum[nums[i]]
		if ok {
			res = append(res, i)
			res = append(res, v)
			break
		} else {
			sum[target-nums[i]] = i
		}
	}

	return res
}
