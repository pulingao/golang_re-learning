package anonymous_function

import (
	"github.com/pulingao/golang_re-learning/tools"
)

/**
 * 求和并调用callback函数对结果进行特殊处理
 */
func sumWorker(data []int, callback func(int)) {
	sum := 0
	for _, num := range data {
		sum += num
	}

	callback(sum)
}

/**
 * 单独放出来，认为可能是一个用法，方便开拓思维
 * @see
 * @param
 */
func T_01() {

	// 打印出求和结果
	sumWorker([]int{1, 2, 3, 4}, func(a int) {
		tools.Error("结果是：%v", a)
	})

	// 判断求和结果是否大于100
	sumWorker([]int{1, 2, 3, 4}, func(a int) {
		if a > 100 {
			tools.Info("sum > 100")
		} else {
			tools.Success("sum <= 100")
		}
	})
}
