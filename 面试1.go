package main

import (
	"fmt"
	"time"
)

func main() {
	val := []int{1, 2, 3, 4}
	goRun(val)
	time.Sleep(time.Second)
}
func goRun(val []int) {
	for _, v := range val {
		go func() {
			fmt.Println(v)
		}()
	}
}
