package main

import (
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
	"strings"
	"time"
)

func main() {
	fmt.Println(`
************************************************************************************************************************
time的一些使用，参考：
1. https://www.jianshu.com/p/2b4686b8de4a
2. https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter04/04.4.md
	timer：相当于js中的setTimeout，指定的时间后执行（执行的前提是需要从timer返回的channel中取出值才能执行）
	ticker：相当于js中的setInterval，以指定的时间重复执行
	Ticker 和 Timer 类似，区别是：Ticker 中的runtimeTimer字段的 period 字段会赋值为 NewTicker(d Duration) 中的d，表示每间隔d纳秒，定时器就会触发一次。

	time.Tick 和 time.NewTicker
		除非程序终止前定时器一直需要触发，否则，不需要时应该调用 Ticker.Stop 来释放相关资源。
		如果程序终止前需要定时器一直触发，可以使用更简单方便的 time.Tick 函数，因为 Ticker 实例隐藏起来了，因此，该函数启动的定时器无法停止。

如果在项目中使用定时器，可以参考：https://github.com/robfig/cron
************************************************************************************************************************
`)

	TickerTest()

}

/**
 * 一个示例
 * @see		https://zhuanlan.zhihu.com/p/225999835
 * @param
 */
func TickerTest() {
	ticker := time.NewTicker(time.Second)
	defer func() {
		ticker.Stop()
	}()

	done := make(chan bool)
	go func() {
		time.Sleep(20 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("达到超时时间，退出")
			count := 0
			// 验证处
			for i := range ticker.C {
				count++
				fmt.Printf("再从定时器中获取：%v\n", i)
				if count >= 2 {
					return
				}
			}
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)

			// 可以认为这里的Sleep确实阻塞了程序2s，同时导致从定时器中获取的代码也是被阻塞了
			time.Sleep(2 * time.Second)
		}
	}

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

func ControlLove() {
	go Love() // 起一个协程去执行定时任务

	TAfter()

	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Println(len(m), m)

	s := "中国"
	for i, v := range []byte(s) {
		fmt.Printf("%v类型：%T，值：%v\n", i, v, v)
	}
	fmt.Println()
	for i, v := range []rune(s) {
		fmt.Printf("%v类型：%T，值：%v\n", i, v, v)
	}

	stop := ""
	for {
		fmt.Scan(&stop)
		if strings.ToUpper(stop) == "Q" {
			break
		}
	}
}

func TAfter() {
	tools.Info("刚开始进入，打印一下")
	//time.After(time.Second * 3) // 单纯使用，并没有什么效果，也不能达到阻塞程序的目的，只有从返回的channel中获取值，才能达到阻塞的状态，对比此行下面的例子
	firstAf := <-(time.After(time.Second * 3))
	tools.Info("aa: %T，%v，%v", firstAf, firstAf, "第一次输出的东西")
	_ = time.AfterFunc(time.Second*2, func() {
		tools.InfoDo("第二次输出一个什么东西")
	})
}

func Love() {
	timer := time.NewTimer(2 * time.Second) // 新建一个Timer
	//for {
	//	select {
	//	case <-timer.C:
	//		fmt.Println("I Love You!")
	//		//timer.Reset(2 * time.Second) // 上一个when执行完毕重新设置
	//		//timer.Reset(500 * time.Millisecond) // 如果不使用Reset来处理，则只运行一次，因此可以简写为下面的样子
	//	}
	//}
	//return

	// 不使用 Reset 时，其实就相当于执行一次，使用select就可以了
	//select {
	//case <-timer.C:
	//	fmt.Println("I Love You!")
	//}

	// 直接使用
	aa := <-timer.C
	tools.Success("aa: %T，%v，%v", aa, aa, "I Love You!")
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

func tickDemo() {
	ticker := time.Tick(time.Millisecond * 800) //定义一个3秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
