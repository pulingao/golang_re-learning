package main

import (
	"fmt"
)

var ch chan int = make(chan int, 2)

func foo(id int) { //id: 这个routine的标号
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	//for i := 0; i < 5; i++ {
	//	fmt.Print(<-ch)
	//}
	go func() {
		for {
			select {
			case data, ok := <-ch:
				fmt.Printf("ok: %v - ", ok)
				if ok {
					fmt.Println("data", data)
				} else {
					fmt.Printf("*")
				}
			}
		}
	}()
	// 这里虽然不会退出，但是确实很好CPU的资源，可以参考 basic/main.go 的实现
	for {

	}
}
