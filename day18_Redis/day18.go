package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("test", []byte("Hello Wolrd"))
	if err != nil {
		panic(err)
	}
	res, err := client.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(res))

	//test hmset
	f := make(map[string]interface{})
	f["name"] = "张三"
	f["age"] = 24
	err = client.Hmset("test_hash", f)
	if err != nil {
		panic(err)
	}

	//test_zset
	_, err = client.Zadd("test_zset", []byte("HELLO"), 100)
	if err != nil {
		panic(err)
	}
}
