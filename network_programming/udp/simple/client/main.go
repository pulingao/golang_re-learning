package main

import (
	"bufio"
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
	"net"
	"os"
	"strings"
)

/**
 * 不需要设置重连的机制，先启动服务端或者先启动客户端，其中一端关闭的情况下，另一端控制好不退出就可以了，并且在关闭的一端重启之后，再发送数据时可以正常发送和接收
 * @see
 * @param
 */
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		//IP:   net.IPv4(0, 0, 0, 0),
		IP:   net.ParseIP("127.0.0.1"),
		Port: 9090,
	})

	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
	}
	defer socket.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		s, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输出错误：", err.Error())
			continue
		}
		s = strings.Trim(s, "\r\n")
		if s == "" {
			continue
		}
		if strings.ToUpper(s) == "Q" {
			tools.Info("收到程序退出指令：%v", strings.ToUpper(s))
			return
		}

		// 发送数据
		sendData := []byte(s)
		n, err := socket.Write(sendData)
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			continue
			//return
		}
		fmt.Println("写入数据长度：", n, "，中文字符长度：", len([]rune(string(sendData))))

		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
		if err != nil {
			fmt.Println("接收数据失败，err:", err)
			//return
		}

		fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)

	}

}
