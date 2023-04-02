package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := `你好，你好，为什么发生这个事情？`
		conn.Write([]byte(msg))
	}
}
