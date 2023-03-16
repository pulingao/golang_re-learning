package iselect

import (
	"fmt"
	"time"
)

func T_basic_chan() {
	ch := make(chan int)
	go func() {
		time.AfterFunc(time.Second*3, func() {
			fmt.Println("开始持续执行")
			for i := 0; i <= 30; i++ {
				time.Sleep(time.Millisecond * 200)
				ch <- i // 这里需要注意，如果读一个有缓冲且无数据的channel 会panic。因此先推一条数据进入channel
			}
			//fmt.Println("开始执行关闭操作")
			//close(ch)
		})
	}()

	go func() {
		time.AfterFunc(time.Second*15, func() {
			fmt.Println("开始持续执行")
			for i := 0; i <= 30; i++ {
				time.Sleep(time.Millisecond * 200)
				ch <- i * 10 // 这里需要注意，如果读一个有缓冲且无数据的channel 会panic。因此先推一条数据进入channel
			}
			//fmt.Println("开始执行关闭操作")
			//close(ch)
		})
	}()

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

	fmt.Println("主程序输出")
}

func T_basic_chan2() {
	ch := make(chan int)

	go writeChan(ch)

	for {
		val, ok := <-ch
		fmt.Println("read ch: ", val)
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
}
