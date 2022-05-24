package main

import (
	"fmt"
	"sync"
)

type Task struct {
	Data string
}

var wg sync.WaitGroup

//生产逻辑
func Producer(tasks chan Task) {
	t := Task{}

	for i := 62; i < 72; i++ {
		t.Data = string(i)
		tasks <- t
	}
}

func ProducerDispatch(tasks chan Task) {
	defer close(tasks)

	Producer(tasks)
}

//消费数据处理逻辑
func Consumer(task Task) {
	fmt.Printf("consum task:%v\n", task)
}

func ConsumerDispatch(tasks chan Task) {
	defer wg.Done()

	for task := range tasks {
		Consumer(task)
	}
}
