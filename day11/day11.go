package main

import (
	"fmt"
	"time"
)

var ch chan int

func test_channel() {
	ch <- 1
	fmt.Print("goruntime 1 end !")
}

func main() {
	ch = make(chan int, 1)
	go test_channel()
	time.Sleep(2 * time.Second)
	fmt.Print("running end!")
	<-ch
	time.Sleep(time.Second)
}
