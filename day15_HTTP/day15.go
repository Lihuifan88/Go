package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//		w.Write([]byte("Hello World!"))
	//	})
	//	http.ListenAndServe("127.0.0.1:8080", nil)

	//	resp, err := http.Get("http://www.baidu.com")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)
	//	fmt.Print(string(body))

	resp, err := http.Post("http://www.sto.cn/Track/Add", "application/x-www-form-urlencoded", strings.NewReader("waybillNo=3349221675141"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))

}
