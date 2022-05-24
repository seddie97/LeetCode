package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Print(i)
		}
		ch <- struct{}{}
	}()

	fmt.Print("to here")

	fin := <-ch

	fmt.Print("fin", fin)
}
