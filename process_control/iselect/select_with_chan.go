package iselect

import (
	"fmt"
	"time"
)

/**
 * 一个反例
 * @see
 * @param
 */
func T_select_counterexample() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}

	fmt.Println("以下是原因：")
	s := `
	select { //不停的在这里检测
	case <-chanl : //检测有没有数据可以读
	//如果chanl成功读取到数据，则进行该case处理语句
	case chan2 <- 1 : //检测有没有可以写
	//如果成功向chan2写入数据，则进行该case处理语句
	
	
	//假如没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞//一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源
	default:
		//如果以上都没有符合条件，那么则进行default处理流程
	}
`
	fmt.Println(s)

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

	fmt.Println()
	fmt.Println("详细的解释：")

	s2 := `
	在一个select语句中，Go会按顺序从头到尾评估每一个发送和接收的语句。
	
	如果其中的任意一个语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。
	
	如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有两种可能的情况： 
		①如果给出了default语句，那么就会执行default的流程，同时程序的执行会从select语句后的语句中恢复。 
		②如果没有default语句，那么select语句将被阻塞，直到至少有一个case可以进行下去。
`
	fmt.Println(s2)

}

/**
 * 字符分割线（character dividing line）
 */
func NewLineWithCDL() {
	fmt.Println()
	fmt.Println("// ---------------------------------------------------------------------------------------------------------------------")
	fmt.Println("//     _____             __   _____       _")
	fmt.Println("//    / ____|           / /  / ____|     | |")
	fmt.Println("//   | |  __  ___      / /  | |  __  ___ | | __ _ _ __   __ _")
	fmt.Println("//   | | |_ |/ _ \\    / /   | | |_ |/ _ \\| |/ _` | '_ \\ / _` |")
	fmt.Println("//   | |__| | (_) |  / /    | |__| | (_) | | (_| | | | | (_| |")
	fmt.Println("//    \\_____|\\___/  /_/      \\_____|\\___/|_|\\__,_|_| |_|\\__, |")
	fmt.Println("//                                                       __/ |")
	fmt.Println("//                                                      |___/")
	fmt.Println("// ---------------------------------------------------------------------------------------------------------------------")
	fmt.Println()
}

/**
 * 基本使用
 * @see 		概念：https://www.topgoer.com/%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6/%E6%9D%A1%E4%BB%B6%E8%AF%AD%E5%8F%A5select.html#golang-select%E7%9A%84%E4%BD%BF%E7%94%A8%E5%8F%8A%E5%85%B8%E5%9E%8B%E7%94%A8%E6%B3%95
 * @param
 */
func T_concept() {
	NewLineWithCDL()
	fmt.Println("基本使用原则：")
	s := `
	select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。

	select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。 

	select中的case语句必须是一个channel操作

	select中的default子句总是可运行的。
	
	如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
	
	如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
	
	如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行
`
	fmt.Println(s)
}

func T_demo() {
	ch := make(chan int)
	go func() {
		time.AfterFunc(time.Second*3, func() {
			fmt.Println("开始持续执行")
			for i := 0; i <= 30; i++ {
				time.Sleep(time.Millisecond * 200)
				ch <- i // 这里需要注意，如果读一个有缓冲且无数据的channel 会panic。因此先推一条数据进入channel
			}
			fmt.Println("开始执行关闭操作")

			close(ch)
		})
	}()

	for {
		select {
		case data, ok := <-ch:
			if ok {
				fmt.Printf("recv-data:[%v]\n", data)
			} else {
				fmt.Println("channel-closed!")
				return
				//goto BREAK
			}
			//default:
			//	fmt.Println("into-default")
		}
	}

	fmt.Println("主程序输出")
}
