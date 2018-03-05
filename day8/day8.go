package main

import (
	"fmt"
)

type People struct {
	name string
	age  int
}

type Student struct {
	People
	school string
	number string
}

//func (pop People) eat(name string) {
//	pop.name = name
//	fmt.Print(pop.name + "在吃饭。。。。")
//}

func main() {
	//	pop := new(People)
	//	pop.name = "张三"
	//	pop.eat(pop.name)
	stu := new(Student)
	stu.People.name = "张三"
	stu.school = "合肥学院"
	fmt.Printf("%v", stu)
}
