package tools

import (
	"fmt"
	"time"
)

/**
 * 一些怪想法的实现
 */

/**
 * 按秒等待并输出（不设计毫秒等输出，此方法的初衷就是为了等待观察效果）
 * @see
 * @param
 */
func WaitInSecondsOutput(seconds int, s string) {
	for i := 1; i <= seconds; i++ {
		fmt.Printf("[%v] %v [%vs]\n", MySysTime(), s, i)
		time.Sleep(time.Second)
	}
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
 * 换行
 */
func NewLine() {
	fmt.Println()
}
