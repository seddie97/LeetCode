package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var msgChan = make(chan string)
var bufChan = make(chan int, 100)

func Producter(wgPd *sync.WaitGroup, max int) {
	for i := 0; i < max; i++ {
		bufChan <- i
		fmt.Fprintf(os.Stderr, "%v put ppppppppp %d\n", os.Getpid(), i)
		time.Sleep(1 * time.Second)
	}
	wgPd.Done()
}

func Customer(wgCs *sync.WaitGroup, name string) {
	for data := range msgChan {

		fmt.Printf("%s read rrr %d\n", name, data)
	}
	wgCs.Done()
	//msgChan <- name
}

func main() {
	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go Producter(wgPd, 5)
	}

	for i := 0; i < 2; i++ {
		wgCs.Add(1)
		go Customer(wgCs, "read"+string(i))

	}
	wgPd.Wait()
	close(msgChan)
	wgCs.Wait()
	//fmt.Printf("%s %s is done!\n", <-msgChan, <-msgChan)
}
