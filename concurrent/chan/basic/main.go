package main

import (
	"fmt"
	"net/http"
	"time"
)

func main2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

}

func rev(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

// *********************************************************************************************************************
//     _____             __   _____       _
//    / ____|           / /  / ____|     | |
//   | |  __  ___      / /  | |  __  ___ | | __ _ _ __   __ _
//   | | |_ |/ _ \    / /   | | |_ |/ _ \| |/ _` | '_ \ / _` |
//   | |__| | (_) |  / /    | |__| | (_) | | (_| | | | | (_| |
//    \_____|\___/  /_/      \_____|\___/|_|\__,_|_| |_|\__, |
//                                                       __/ |
//                                                      |___/
// *********************************************************************************************************************

func T_test() {
	ch := make(chan int)

	go writeChan(ch)

	for {
		val, ok := <-ch
		fmt.Println("read ch: ", val, "，OK：", ok)
		//输出
		//read ch:  0 ，OK： true
		//read ch:  1 ，OK： true
		//read ch:  2 ，OK： true
		//read ch:  3 ，OK： true

		// 这里为什么多出一个0?
		//	因为write完毕后关闭了channel，但是关闭的channel是可以读取到channel类型的零值的，但是读取状态是false（OK的值）
		//	因此int类型的零值就是0，但是ok=false
		//read ch:  0 ，OK： false

		//end
		if !ok {
			break
		}
	}

	time.Sleep(time.Second)
	fmt.Println("end")
}

func writeChan(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
	}
	close(ch)
}

// *********************************************************************************************************************
//     _____             __   _____       _
//    / ____|           / /  / ____|     | |
//   | |  __  ___      / /  | |  __  ___ | | __ _ _ __   __ _
//   | | |_ |/ _ \    / /   | | |_ |/ _ \| |/ _` | '_ \ / _` |
//   | |__| | (_) |  / /    | |__| | (_) | | (_| | | | | (_| |
//    \_____|\___/  /_/      \_____|\___/|_|\__,_|_| |_|\__, |
//                                                       __/ |
//                                                      |___/
// *********************************************************************************************************************

var ch chan int

func main() {
	ch = make(chan int)

	go HttpServer()

	SendAfterSmall()

	fmt.Println("主程序输出")
	Receive()
}

type IndexHandler struct {
	content string
}

func (ih *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ih.content)
}

type HttpHandlerFunc struct {
	Desc string
}

func (hlf *HttpHandlerFunc) WithDesc(desc string) {
	hlf.Desc = desc
}

/**
 * 定义一个结构体，并定义多个Request Handler方法
 * @see			HLF：HandlerFunc
 * @param
 */
func (hlf *HttpHandlerFunc) HelloHLF(w http.ResponseWriter, req *http.Request) {
	go SendAfterBig()
	fmt.Fprintf(w, "请求详细信息，请求方法: %s, URL: %s\n", req.Method, req.URL)
}

func HttpServer() {
	hhf := &HttpHandlerFunc{}
	// 常规方法
	http.Handle("/", &IndexHandler{content: "hello world!"})

	// 自定义函数
	http.HandleFunc("/add", hhf.HelloHLF)

	// 监听
	http.ListenAndServe(":8001", nil)
}

func SendAfterSmall() {
	time.AfterFunc(time.Second*1, func() {
		fmt.Printf("\n[%v] SendAfterSmall 开始持续执行\n", time.Now().Format("2006-01-02 15:04:05"))
		for i := 0; i <= 30; i++ {
			time.Sleep(time.Millisecond * 10)
			ch <- i // 这里需要注意，如果读一个有缓冲且无数据的channel 会panic。因此先推一条数据进入channel
		}
		//fmt.Println("开始执行关闭操作")
		//close(ch)
	})
}

func SendAfterBig() {
	time.AfterFunc(time.Millisecond*800, func() {
		fmt.Printf("\n[%v] SendAfterBig 开始持续执行\n", time.Now().Format("2006-01-02 15:04:05"))
		for i := 0; i <= 30; i++ {
			time.Sleep(time.Millisecond * 50)
			ch <- i * 10 // 这里需要注意，如果读一个有缓冲且无数据的channel 会panic。因此先推一条数据进入channel
		}
		//fmt.Println("开始执行关闭操作")
		//close(ch)
	})
}

func Receive() {
	for {
		select {
		case data, ok := <-ch:
			if ok {
				fmt.Printf("recv-data:[%v]\n", data)
			}
			//else {
			//	fmt.Println("channel-closed!")
			//	return
			//	//goto BREAK
			//}
			//default:
			//	fmt.Println("into-default")
		}
	}
}
