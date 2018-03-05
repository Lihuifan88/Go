package main

import (
	"fmt"
)

func main() {
	//匿名函数
	//	var f = func(x, y int) int {
	//		return x + y
	//	}
	//	fmt.Print(f(3, 4))

	//	for i := 0; i < 5; i++ {
	//		defer fmt.Println(i)
	//	}
	//	panic("异常")

	//	defer func() {
	//		fmt.Print("After 1")
	//	}()
	//	fmt.Println("Before defer")

	type getNum func(x, y int) int
	var f getNum = func(x, y int) int {
		return x + y
	}
	fmt.Print(f(3, 4))

}
