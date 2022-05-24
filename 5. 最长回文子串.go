package main

func longestPalindrome(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l-1 >= 0 && s[l] == s[l-1] {
			l--
		}
		for r+1 < len(s) && s[r+1] == s[r] {
			r++
		}
		for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
		}
		if (r - l + 1) > len(res) {
			res = s[l : r+1]
		}
	}
	return res
}
