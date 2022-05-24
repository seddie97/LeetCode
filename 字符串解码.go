package main

import "fmt"

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}

func decodeString(s string) string {
	res := ""
	var sS []string
	var sI []int
	cur := ""
	n := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]) - '0'
		} else if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			cur += string(s[i])
		} else if s[i] == ']' {
			t := sI[len(sI)-1]
			sI = sI[:len(sI)-1]
			temp := sS[len(sS)-1]
			sS = sS[:len(sS)-1]
			for t > 0 {
				temp += cur
				t--
			}
			cur = temp
		} else {
			sI = append(sI, n)
			sS = append(sS, cur)
			n = 0
			cur = ""
		}
		res = cur
	}

	return res
}
