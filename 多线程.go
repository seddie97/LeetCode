package main

import (
	"fmt"
	"sync"
)

var wg12 sync.WaitGroup

func hello(i int) {
	defer wg12.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {

	for i := 0; i < 10; i++ {
		wg12.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg12.Wait() // 等待所有登记的goroutine都结束
}
