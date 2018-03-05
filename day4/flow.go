package main

import (
	"fmt"
)

func main() {
	//	var flag bool
	//	flag = false
	//	if flag {
	//		fmt.Print("输出")
	//	}

	//	sum := 0
	//	i := 1
	//	for i <= 10 {
	//		sum = i + sum
	//		i++
	//	}
	//	fmt.Print(sum)

	//	i := 3
	//	switch i {
	//	case 1:
	//		fmt.Print(1)
	//		//fallthrough
	//	case 2:
	//		fmt.Print(2)
	//	default:
	//		fmt.Print(3)
	//	}

	//	x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := make(map[string]int)
	x["张三"] = 10
	x["李四"] = 20
	for _, v := range x {
		fmt.Println(v)
	}

}
