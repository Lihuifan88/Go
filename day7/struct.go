package main

import (
	"fmt"
)

//定义了一个struct
type Student struct {
	id   int
	name string
	age  int
}

func main() {
	//var s *Student = new(Student)
	s := new(Student)
	s.id = 1
	s.name = "张三"
	s.age = 20
	fmt.Printf("%v", s)
}
