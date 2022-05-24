package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	wg.Wait()
}
