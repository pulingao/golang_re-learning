package usual

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pulingao/golang_re-learning/tools"
	"io"
	"net/http"
	"sync"
	"time"
)

func A(i int) {
	fmt.Println("我是A", i)
}

func GR_SyncWaitGroup() {

	//使用方式参考：http://c.biancheng.net/view/108.html
	var wg sync.WaitGroup

	fmt.Println("我是main，时间：", tools.MySysTime())

	wg.Add(1)
	go func(i int) {
		defer wg.Done()
		A(i)
	}(1)

	wg.Add(1)
	go func() {
		defer wg.Done() // 注意defer中的Done方法的使用
		for i := 1; i <= 3; i++ {
			fmt.Printf("等待：%v s\n", i)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

	fmt.Println("执行完了，时间：", tools.MySysTime())
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

func T_wg() {
	// 声明一个等待组，参考：http://c.biancheng.net/view/108.html
	var wg sync.WaitGroup

	// 准备一系列的网站地址
	var urls = []string{
		"https://github.com/",
		"https://www.qiniu.com/",
		"https://www.baidu.com",
		"http://www.baidu.com",
	}

	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时, 将等待组增加1
		wg.Add(1)
		// 开启一个并发
		go func(url string) {
			// 使用defer, 表示函数完成时将等待组值减1
			defer wg.Done()
			// 使用http访问提供的地址
			res, err := http.Get(url)
			defer res.Body.Close()
			if err != nil {
				fmt.Println(tools.MySysTime(), url, "，请求错误：", err)
				return
			}

			// res.Body是一个io.ReadCloser类型，是一个流，因此不能读取多次，以此读取完成后就被关闭了，因为不能将一个传入的TCP连接倒回去再读取
			// 因此
			//	1.不能使用 io.ReadSeeker 的方式来倒回到流初始的地方进行读取，因为没有实现 io.Seeker 接口
			//	2.所以只能使用 io.TeeReader 的方式来复制流
			// 参考：
			//	1.https://www.zadmei.com/rhcgzdir.html
			//	2.https://www.golang-tech-stack.com/qa/3150
			buf := &bytes.Buffer{}
			tee := io.TeeReader(res.Body, buf)
			bds, err := io.ReadAll(tee)
			if err != nil {
				fmt.Println("复制流错误：", err)
			}
			tools.Warning("获取URL：%v，Body长度：%v", url, len(string(bds)))

			// 使用goquery获取一些信息
			doc, err := goquery.NewDocumentFromReader(buf)
			if err != nil {
				fmt.Println(tools.MySysTime(), url, "，html解析错误：", err)
				return
			}

			title := doc.Find("title").Text()

			// 访问完成后, 打印地址和可能发生的错误
			fmt.Println(tools.MySysTime(), url, "，标题内容：", title)

		}(url) // 通过参数传递url地址
	}
	// 等待所有的任务完成
	wg.Wait()

	fmt.Println("over", tools.MySysTime())
}
