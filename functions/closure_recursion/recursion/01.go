package recursion

import (
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
)

func T_00() {
	s := `
************************************************************************************************************************
递归，就是在运行的过程中调用自己。 一个函数调用自己，就叫做递归函数。

构成递归需具备的条件：

    1.子问题须与原始问题为同样的事，且更为简单。
    2.不能无限制地调用本身，须有个出口，化简为非递归状况处理。
************************************************************************************************************************
`
	fmt.Println(s)
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func T_01() {
	var n int = 9
	tools.Info("%v 的阶乘是：%v", n, factorial(n))

	tools.NewLineWithCDL()

	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}
