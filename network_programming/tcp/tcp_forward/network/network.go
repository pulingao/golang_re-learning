package network

import (
	"io"
	"log"
	"net"
)

const (
	KeepAlive     = "KEEP_ALIVE"
	NewConnection = "NEW_CONNECTION"
)

func CreateTCPListener(addr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpListener, nil
}

func CreateTCPConn(addr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpListener, nil
}

func Join2Conn(local *net.TCPConn, remote *net.TCPConn) {
	go joinConn(local, remote)
	go joinConn(remote, local)
}

func joinConn(local *net.TCPConn, remote *net.TCPConn) {
	defer local.Close()
	defer remote.Close()

	_, err := io.Copy(local, remote)
	if err != nil {
		log.Println("copy failed ", err.Error())
		return
	}

	//可以打印发送的内容，不过都是乱码
	//var buffer = make([]byte, 4096000)
	//for {
	//	n, err := local.Read(buffer)
	//	if err != nil {
	//		fmt.Printf("Unable to read from input, error: %s\n", err.Error())
	//		break
	//	}
	//	s := buffer[:n]
	//	tools.Info("转发的内容：%v", string(s))
	//	n, err = remote.Write(s)
	//	if err != nil {
	//		fmt.Printf("Unable to write to output, error: %s\n", err.Error())
	//		break
	//	}
	//}

}
