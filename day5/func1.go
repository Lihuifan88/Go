//func1
package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {
	return b, a
}
func main() {
	//	a := 1
	//	b := 2
	//	a, b = swap(a, b)
	//	fmt.Printf("%d %d", a, b)
	a := 1
	test(&a)
	fmt.Printf("%d", a)

}

func test(a *int) *int {
	*a = *a + 1
	return a
}
