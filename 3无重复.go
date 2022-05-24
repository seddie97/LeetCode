package main

func lengthOfLongestSubstring(s string) int {
	res := 0
	st := 0
	mark := [128]int{}

	for i := 0; i < 128; i++ {
		mark[i] = 0
	}

	for i, v := range s {
		st = max(st, mark[v])
		res = max(res, i-st+1)
		mark[v] = i
	}
	return res
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
