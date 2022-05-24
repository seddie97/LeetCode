package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func producter(name string, ch chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println(name, "product :", i)
		ch <- i
	}
	wg.Done()
}

func consumer(name string, ch chan int) {

	for i := 0; i < 4; i++ {
		h := <-ch
		fmt.Println(name, "consum :", h)
	}
	fmt.Println(name, " is Done")
	wg.Done()
}

func main() {

	ch := make(chan int, 10)
	defer close(ch)
	wg.Add(4)
	go producter("生产者A", ch)
	go producter("生产者B", ch)
	go consumer("消费者A", ch)
	go consumer("消费者B", ch)
	wg.Wait()

	//time.Sleep(10 * time.Second)
}
