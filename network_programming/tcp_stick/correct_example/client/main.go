package main

import (
	"fmt"
	"github.com/pulingao/golang_re-learning/network_programming/tcp_stick/correct_example/proto"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	//如果内容很长，发送不成功
	//猜想：如果超出了buffer的限制，应该会出现收取不到的情况
	for i := 0; i < 20; i++ {
		msg := `来一个很长的内容，还有回车字符等，看看是否也是一致的？我们先说一下UDP的概念和作用 UDP是用户数据报协议，是一个简单的面向数据报的运输层协议。UDP不提供可靠性，它只是把应用程序传给IP层的数据报发送出去，但是并不能保证它们能到达目的地。由于UDP在传输数据报前不用在客户和服务器之间建立一个连接，且没有超时重发等机制，故而传输速度很快。那么下面我们就说一下在go语言环境先如何实现客户端与服务器端UDP协议连接通讯`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		fmt.Println(len(data))

		conn.Write(data)
		if i == 9 {
			os.Exit(1)
		}
	}
}
