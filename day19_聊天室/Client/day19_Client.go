//chatroom client
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error :%s", err.Error())
		os.Exit(1)
	}
}

func MessageSend(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)
		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			fmt.Println("client connect failure:" + err.Error())
			break
		}
	}

}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer conn.Close()
	go MessageSend(conn)
	//conn.Write([]byte("Hello 服务器"))
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		CheckError(err)
		fmt.Println("Receive Server Message Content" + string(buf))
	}
	fmt.Println("Client Programma End!")
}
