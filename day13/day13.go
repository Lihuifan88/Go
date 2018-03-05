package main

import (
	"fmt"
	//	"runtime"
	"strconv"
	"time"
)

var ch chan int

func main() {
	ch = make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			if i == 10 {
				fmt.Println("开始让出CPU")
				//				runtime.Gosched() //主动让出CPU
				<-ch //阻塞 让出CPU
			}
			fmt.Println("routime 1:" + strconv.Itoa(i))
		}
	}()
	go func() {
		for i := 100; i < 200; i++ {
			fmt.Println("routime 2:" + strconv.Itoa(i))
		}
	}()
	time.Sleep(time.Second)
}
