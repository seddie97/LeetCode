package main

import (
	"fmt"
	"sync"
)

var num int
var mtx sync.Mutex
var wg1 sync.WaitGroup

func add() {
	mtx.Lock()
	defer mtx.Unlock()
	defer wg1.Done()
	num += 1
}
func main() {
	for i := 0; i < 100; i++ {
		wg1.Add(1)
		go add()
	}
	wg1.Wait()
	fmt.Println("num:", num)
}
