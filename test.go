package main

import (
	"fmt"
	"sync"
)

// 实现顺序输出123
func TestPrintNum() {
	w := sync.WaitGroup{}
	a := make(chan struct{}, 1)
	b := make(chan struct{})
	w.Add(2)
	go func() {
		defer w.Done()
		for i := 1; i < 21; i++ {
			<-a
			if i%2 != 0 {
				fmt.Println(i)
			}
			b <- struct{}{}
		}
	}()
	go func() {
		defer w.Done()
		for i := 1; i < 21; i++ {
			<-b
			if i%2 == 0 {
				fmt.Println(i)
			}
			a <- struct{}{}
		}
	}()
	a <- struct{}{}
	w.Wait()
}
