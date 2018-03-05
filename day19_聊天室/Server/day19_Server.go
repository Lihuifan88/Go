//chatroom server
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var onlineConns = make(map[string]net.Conn)
var messaegQueue = make(chan string, 1000)
var quitChan = make(chan bool)

func ConsumeMessage() {
	for {
		select {
		case message := <-messaegQueue:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChan:
			break
		}
	}
}

func doProcessMessage(message string) {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := contents[0]
		sendMessage := contents[1]
		addr = strings.Trim(addr, " ")
		if conn, ok := onlineConns[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("Online sendMessage Failure!")
			}
		}
	}

}

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error :%s", err.Error())
		os.Exit(1)
	}
}

func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024) //建立一个缓冲区
	defer conn.Close()
	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}
		if numOfBytes != 0 {
			message := string(buf[0:numOfBytes])
			messaegQueue <- message
		}
	}
}

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer listen_socket.Close()
	fmt.Println("Server Has Waiting.....")
	//	for {
	//		conn, err := listen_socket.Accept()
	//		CheckError(err)
	//		go ProcessInfo(conn)
	//	}
	go ConsumeMessage()
	for {
		conn, err := listen_socket.Accept()
		CheckError(err)
		//将conn存储到onlineConnes映射表中
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns[addr] = conn
		for i := range onlineConns {
			fmt.Println(i)
		}
		go ProcessInfo(conn)
	}

}
