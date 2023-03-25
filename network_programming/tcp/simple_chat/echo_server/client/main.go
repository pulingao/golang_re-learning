package main

import (
	"bufio"
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
	"net"
	"strings"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()

	tools.Info("客户端连接成功：%v，当前连接 %v:{%v}", conn.RemoteAddr().String(), conn.LocalAddr().Network(), conn.LocalAddr().String())

	go onMessageReceived(conn)

	go onMessageSend(conn)

	<-quitSemaphore
}

func onMessageReceived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	clientIdy := conn.LocalAddr().String()
	for {
		msg, err := reader.ReadString('\n')
		msg = strings.Trim(msg, "\r\n")
		tools.Info("Client：%v，收到Server消息：%v", clientIdy, msg)
		if err != nil {
			quitSemaphore <- true
			break
		}

	}
}

func onMessageSend(conn *net.TCPConn) {
	count := 1
	clientIdy := conn.LocalAddr().String()
	for {
		msg := "" // 不在全局定义，防止第一次输出值后，再回车时拿到第一次的值，这里每次都赋值为空，则能保证不拿到上次的值
		fmt.Scanln(&msg)
		if msg == "" {
			continue
		}
		clientSendMsg := fmt.Sprintf("Client：%v，(%v)%v\n", clientIdy, count, msg)
		b := []byte(clientSendMsg)
		count++
		conn.Write(b)
	}

}
