package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	var s []int

	for i := 0; i < n; i++ {
		for len(s) > 0 && temperatures[i] > temperatures[s[len(s)-1]] {
			res[s[len(s)-1]] = i - s[len(s)-1]
			s = s[:len(s)-1]
		}

		s = append(s, i)
	}

	return res
}
