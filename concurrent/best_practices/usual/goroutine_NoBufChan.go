package usual

import (
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
)

func AA(i int) {
	fmt.Println("我是AA", i)
}

/**
 * 无buf的channel控制
 * @see
 * @param
 */
func GR_NoBufChan() {
	ch := make(chan bool, 1)

	tools.Info("Main函数开始")

	go func(i int, chp chan<- bool) {
		defer close(chp)

		AA(i)

		tools.WaitInSecondsOutput(3, "例行等待")

		fmt.Println("finish")
		chp <- true

	}(1, ch)

	tools.Info("wait")

	<-ch

	tools.Success("执行完了")

}
