package main

import (
	"fmt"
	"regexp"
)

func main() {
	flag, _ := regexp.Match("[a-zA-Z]{3}", []byte("abC"))
	fmt.Print(flag)
}
