package main

import (
	"fmt"
)

func main() {
	//	y := make([]int, 3, 5)
	//	y = append(y, 5, 6, 7)
	//	var slice []int
	//	slice = append(slice, 1, 2, 3)
	student := make(map[string]float32)
	student["张三"] = 18.2
	fmt.Printf("%v", student)
}
