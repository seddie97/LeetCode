package main

import (
	"fmt"
)

func main() {
	s := ""
	fmt.Sscan(s)
	fmt.Print(s)
	l := len(s) - 1
	res := 0

	for i := l; i >= 0; i-- {
		fmt.Printf("----------:" + string(s[i]))
		if s[i] != ' ' {
			res++
		}
	}

	fmt.Print(res)

}
