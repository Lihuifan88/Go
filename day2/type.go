package main

//包后面紧跟import
import (
	"fmt"
)

func main() {
	//	var x int
	//	x = 2
	//简短声明，无法用于函数体外部
	//	y, z := "你好", "GOLANG"
	//	fmt.Printf("%d %s%s", x, y, z)
	//	var t [10]int
	//	t[1] = 2
	t := [10]int8{2, 3}
	s := [2]string{"Hello", "GOLANG"}
	fmt.Print("%v,%v", t, s)
}
