package main

import (
	"encoding/json"
	"fmt"
)

//json.Marshal(v interface{})(byte[],error)
func main() {
	//map to json
	m := make(map[string]float64)
	m["张三"] = 100.4
	s, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

}
