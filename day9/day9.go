package main

import (
	"fmt"
)

//type Animal interface {
//	Run()
//	Fly()
//}
//type Animal2 interface {
//	Fly()
//}
//type Bird struct {
//}

//func (bird Bird) Run() {
//	fmt.Print("Bird is Running!!!")

//}
//func (bird Bird) Fly() {
//	fmt.Print("Bird is Flyying!!!")

//}

func main() {
	//	var animal Animal
	//	bird := new(Bird)
	//	animal = bird
	//	animal.Fly()

	//	var animal2 Animal2
	//	animal2 = animal
	//	animal2.Fly()
	var v interface{}
	v = "空接口"
	//	if v, ok := v.(float64); ok {
	//		fmt.Print(v, ok)
	//	} else {
	//		fmt.Print(v, ok)
	//	}
	switch v.(type) {
	case string:
		fmt.Print("这是一个字符串")
	}
}
