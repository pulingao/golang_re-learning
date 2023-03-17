package _defer

import (
	"errors"
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
	"io"
	"os"
)

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	fmt.Println("下面的内容是否执行?")
	i = a / b
	return
}

func DeferTrap() {
	foo(2, 0)

	// 输出
	//third defer err divided by zero!
	//second defer err <nil>
	//first defer err <nil>

	s := `
========================================================================================================================
解释：
	1.defer按照注册顺序执行，因此先输出最后一个定义的defer内容
	2.第二个是作为值传递的形式来执行defer，当执行到这里时，值传递的形式只是保存当时声明时的err的值，声明时err=nil，因此就是nil
	3.第一个也是作为值传递，结果和第二个一样
	4.为什么后面的内容没有输出，因为判断b==0后就直接return了，没有机会执行后面的代码
`
	fmt.Println(s)
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

func testNil() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

/**
 * 当defer定义的函数是nil
 * @see
 * @param
 */
func DeferNil() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

			fmt.Println("解释：名为 testNil 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。然而值得注意的是，run() 的声明是没有问题，因为在test函数运行完成后它才会被调用。")
		}
	}()
	testNil()

}

// *********************************************************************************************************************
//
//	  _____             __   _____       _
//	 / ____|           / /  / ____|     | |
//	| |  __  ___      / /  | |  __  ___ | | __ _ _ __   __ _
//	| | |_ |/ _ \    / /   | | |_ |/ _ \| |/ _` | '_ \ / _` |
//	| |__| | (_) |  / /    | |__| | (_) | | (_| | | | | (_| |
//	 \_____|\___/  /_/      \_____|\___/|_|\__,_|_| |_|\__, |
//	                                                    __/ |
//	                                                   |___/
//
// *********************************************************************************************************************
func do() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			err := f.Close()
			if err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			} else {
				tools.Info("文件：book.txt，正确关闭")
			}
		}(f)
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		// 注意这里的使用，使用了闭包+传参的方式，确保多个文件的句柄被正确的拿到和关闭
		// 因为文件的Close是一个接收者是指针的方法：func (f *File) Close() error {}
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			} else {
				tools.Info("文件：another-book.txt，正确关闭")
			}
		}(f)
	}

	return nil
}

func DeferCloseFile() {
	if err := do(); err != nil {
		tools.Error("错误类型：%T，信息：%v", err, err)
	}
}
