package main

import (
	"fmt"
	//"time"
)

//func add(x, y int) {
//	fmt.Println(x + y)
//}

func channel(x int, quite chan int) {
	fmt.Println(x)
	quite <- 1

}
func main() {
	//	for i := 0; i < 10; i++ {
	//		go add(i, i)
	//	}
	//	time.Sleep(1)
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go channel(i, chs[i])
	}

	for _, v := range chs {
		<-v
	}

}
