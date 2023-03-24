package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Message string
}

var conf = &Config{
	Message: "热重载之前的值",
}

func router() {
	log.Println("开始启动")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, _ := json.Marshal(conf)
		_, _ = w.Write(res)
	})

	go func() {
		log.Fatal(http.ListenAndServe(":8989", nil))
	}()
}

func main() {
	router()
	sigCh := make(chan os.Signal, 1)
	// 参考：
	//1.https://juejin.cn/post/7197587307276533820?
	//2.https://colobu.com/2015/10/09/Linux-Signals/
	//Signal handler可以通过signal()系统调用进行设置。如果没有设置，缺省的handler会被调用，当然进程也可以设置忽略此信号。
	//有两种信号不能被拦截和处理: SIGKILL和SIGSTOP（因此这里设置了捕获SIGKILL也是没有用的）
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	for {
		multiSignalHandler(<-sigCh)
	}
}

func multiSignalHandler(signal os.Signal) {
	switch signal {
	case syscall.SIGHUP:
		log.Println("Signal:", signal.String())
		n := rand.Intn(10)
		log.Println("开始热重载，随机取值：", n)
		conf.Message = fmt.Sprintf("热重载已经完成，随机取值：%v", n)
	case syscall.SIGINT:
		log.Println("Signal:", signal.String())
		log.Println("按 Ctrl+C 中断")
		os.Exit(0)
	case syscall.SIGTERM:
		log.Println("Signal:", signal.String())
		log.Println("进程被杀死。")
		os.Exit(0)
	case syscall.SIGKILL: //捕获不到，被强制
		log.Println("Signal:", signal.String())
		log.Println("进程被强制杀死。")
		os.Exit(1)
	default:
		log.Println("未处理/未知信号")
	}
}
