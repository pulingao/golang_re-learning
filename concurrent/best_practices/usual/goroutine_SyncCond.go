package usual

import (
	"github.com/pulingao/golang_re-learning/tools"
	"sync"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	tools.Info("%v 开始读取", name)
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	tools.Info("开始写入")
	tools.WaitInSecondsOutput(3, "模拟处理过程")

	c.L.Lock()
	done = true
	// 处理完成，释放锁
	c.L.Unlock()

	tools.Warning("%v 唤醒所有执行到Wait()的逻辑", name)

	// 广播给其他的wait()的逻辑
	c.Broadcast()
}

/************************************************************************************************************************
sync.Cond 介绍
	sync.Cond 条件变量用来协调想要访问共享资源的那些 goroutine，当共享资源的状态发生变化的时候，它可以用来通知被互斥锁阻塞的 goroutine。
	cond.L.Lock()和cond.L.Unlock()：也可以使用lock.Lock()和lock.Unlock()，完全一样，因为是指针转递
	cond.Wait()：Unlock()->*阻塞等待通知(即等待Signal()或Broadcast()的通知)->收到通知*->Lock()
	cond.Signal()：Signal 唤醒一个协程，若没有Wait()，也不会报错。Signal()通知的顺序是根据原来加入通知列表(Wait())的先入先出
	cond.Broadcast(): 通知所有Wait()了的，若没有Wait()，也不会报错
*/

/*
*
  - 代码逻辑释义
    1.done 即互斥锁需要保护的条件变量。
    2.read() 调用 Wait() 等待通知，直到 done 为 true。
    3.write() 接收数据，接收完成后，将 done 置为 true，调用 Broadcast() 通知所有等待的协程。
    4.write() 中的暂停了3s，一方面是模拟耗时，另一方面是确保前面的 3 个 read 协程都执行到 Wait()，处于等待状态。main 函数最后暂停了 2s，确保所有操作执行完毕。
*/
func GR_SyncCond() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	tools.WaitInSecondsOutput(2, "等待全部执行完毕")
}
