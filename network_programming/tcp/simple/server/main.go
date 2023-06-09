package main

import (
	"bufio"
	"fmt"
	"net"
)

func T_0() {
	s := `
***************************************************************************************************************************************************************************************
Socket介绍
	Socket是BSD UNIX的进程通信机制，通常也称作”套接字”，用于描述IP地址和端口，是一个通信链的句柄。
	Socket可以理解为TCP/IP网络的API，它定义了许多函数或例程，程序员可以用它们来开发TCP/IP网络上的应用程序。
	电脑上运行的应用程序通常通过”套接字”向网络发出请求或者应答网络请求。

Socket是应用层与TCP/IP协议族通信的中间软件抽象层。
	在设计模式中，Socket其实就是一个门面模式，它把复杂的TCP/IP协议族隐藏在Socket后面，对用户来说只需要调用Socket规定的相关函数，让Socket去组织符合指定的协议数据然后进行通信

常用的Socket类型有两种：
	流式Socket和数据报式Socket，
		流式是一种面向连接的Socket，针对于面向连接的TCP服务应用
			TCP：比较靠谱，面向连接，比较慢
		数据报式Socket是一种无连接的Socket，针对于无连接的UDP服务应用
			UDP：不是太靠谱，比较快，一般直播使用UDP
***************************************************************************************************************************************************************************************
`
	fmt.Println(s)
}

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("从客户端读取失败, 错误:", err)
			break
		}
		revStr := string(buf[:n])
		responseStr := fmt.Sprintf("response{ %v }", revStr)
		fmt.Println("收到客户端发来的数据：", revStr, "，准备响应：", responseStr)
		conn.Write([]byte(responseStr)) // 发送（响应）数据
	}
}

func main() {
	T_0()
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("监听失败，错误:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("建立链接失败, 错误信息:", err)
			continue
		}
		fmt.Println(conn)
		go process(conn) // 启动一个goroutine处理连接
	}
}
