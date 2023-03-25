package main

import (
	"bufio"
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
	"math/rand"
	"net"
	"strings"
	"time"
)

// 使用 nc 读写 TCP/UDP 的连接，参考：https://www.ifmicro.com/%E8%AE%B0%E5%BD%95/2017/12/12/netcat-usage/

var connMap map[string]*net.TCPConn

func main() {
	var tcpAddr *net.TCPAddr
	connMap = make(map[string]*net.TCPConn)

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "0.0.0.0:9999")
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err == nil {
		tools.Success("服务启动成功：%v", tcpAddr)
		fmt.Println(`
**********************************************************************************************************
使用 nc 读写 TCP/UDP 的连接，参考：https://www.ifmicro.com/%E8%AE%B0%E5%BD%95/2017/12/12/netcat-usage/
  1.与 www.baidu.com 的 80 端口建立一个 TCP 连接(本地端口随机)
    nc www.baidu.com 80
  2.使用本地 1234 端口与 www.baidu.com 的 80 端口建立一个 TCP 连接
    nc -p 1234 www.baidu.com 80
  3.与 www.baidu.com 的 80 端口建立一个 TCP 连接, 超过 5s 自动断开
    nc -w 5 www.baidu.com 80
  4.与 host.example.com 的 53 端口建立一个 UDP 连接
    nc -u host.example.com 53
  5.使用 10.1.2.3 作为本地 IP 与 host.example.com 的 42 端口建立一个 TCP 连接
    nc -s 10.1.2.3 host.example.com 42
  6.详细展示 127.0.0.1 的 TCP 端口 9997-9999 的开启状况, 使用 -v 选项
    nc -v -w 3 -s 127.0.0.1 192.168.31.178 -z 9997-9999
**********************************************************************************************************
`)
	}
	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		// 加入tcp连接到连接中心
		tcpConnIfy := tcpConn.RemoteAddr().String()
		tools.Info("检测到客户端连接：%v", tcpConnIfy)
		connMap[tcpConnIfy] = tcpConn

		go tcpPipe(tcpConn)
	}

}

func tcpPipe(curConn *net.TCPConn) {
	ipStr := curConn.RemoteAddr().String()
	defer func() {
		tools.Warning("客户端：%v，异常或结束，服务端主动断开连接", ipStr)
		if _, ok := connMap[ipStr]; ok {
			delete(connMap, ipStr)
		}
		curConn.Close()
	}()

	// 从传入的当前连接中读取
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(curConn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			tools.Error("读取客户端：%v，出现错误或结束：%v", ipStr, err)
			return
		}
		message = strings.Trim(message, "\r\n")
		if message == "" {
			continue
		}
		tools.Success("Receive From [%v]，内容：%v", ipStr, message)

		// 返回给客户端（当前连接响应）
		randInt := rand.Intn(65535)
		msg := fmt.Sprintf("服务端响应：%v，原接收消息体：%v\n", randInt, message)
		b := []byte(msg)
		curConn.Write(b)

		// 广播
		for i, v := range connMap {
			if v == curConn {
				tools.Warning("连接中心：%v，当前连接：%v，已经发送过数据，广播操作跳过不处理", v, curConn)
				continue
			}
			msg := fmt.Sprintf("广播[%v]响应：%v，> %v，原接收消息体：%v\n", len(connMap), randInt, i, message)
			v.Write([]byte(msg))
		}
	}
}
