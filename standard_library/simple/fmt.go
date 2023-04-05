package main

import (
	"bufio"
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
	"os"
	"strings"
)

/**
 * 一些常用包的，暂时未知的处理方式
 * @see
 * @param
 */
func main() {
	err := fmt.Errorf("发生了自定义错误：%v", "错误信息")
	fmt.Println(err)

	tools.NewLineWithCDL()

	//PlaceHolder()
	//UScan()
	//bufIoDemo()
	UFscan()
}

/**
 * 这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据
 * @see
 * @param
 */
func UFscan() {
	var (
		i int
		b bool
		s string
	)
	//fmt.Fscan()
	//r := strings.NewReader("10 false GFG")
	//n, err := fmt.Fscan(r, &i, &b, &s)

	//fmt.Fscanf
	//r := strings.NewReader("i:10 b:false s:GFG")
	//n, err := fmt.Fscanf(r, "i:%d b:%t s:%s", &i, &b, &s)

	// 需要将字符串转换为流，即要是一个 io.Reader 的对象
	r := strings.NewReader("10 false GFG")
	n, err := fmt.Fscanln(r, &i, &b, &s)

	// If the above function returns an error then
	// below statement will be executed
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}

	// Printing each type of scanned texts
	fmt.Println(i, b, s)

	// It returns the number of items
	// successfully scanned
	fmt.Println(n)

	fmt.Println()

	USscan()

}

func USscan() {
	// Declaring some variables
	var name string
	var alphabet_count int
	var float_value float32
	var boolean_value bool

	// Calling the Sscan() function which returns the number of elements successfully scanned and error if it persists
	// 直接从字符串中获取，不需要转换为流
	n, err := fmt.Sscan("GeeksforGeeks 13 6.7 true", &name, &alphabet_count, &float_value, &boolean_value)

	// Below statements get
	// executed if there is any error
	if err != nil {
		panic(err)
	}

	// Printing the number of
	// elements and each elements also
	fmt.Printf("找到了%d个变量: %s, %d, %g, %t", n, name, alphabet_count, float_value, boolean_value)
}

/**
 * 有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现
 * @see
 * @param
 */
func bufIoDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%%#v：%#v\n", text)
	fmt.Printf("%%v：%v\n", text)
}

/**
 * 占位符相关
 * @see
 * @param
 */
func PlaceHolder() {
	fmt.Println(`
占位符	说明
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	打印值的类型
%%	百分号
`)

	type Data struct {
		Ip       string   `json:"ip"`
		User     string   `json:"user"`
		From     string   `json:"from"`
		Type     string   `json:"type"`
		Content  string   `json:"content"`
		UserList []string `json:"user_list"`
	}

	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%%v：%v\n", o)
	fmt.Printf("%%+v：%+v\n", o)
	fmt.Printf("%%#v：%#v\n", o)
	fmt.Printf("%%T：%T\n", o)
	fmt.Printf("100%%\n")

	fmt.Println()

	d := Data{
		Ip:       "127.0.0.1",
		User:     "u1123",
		From:     "成都",
		Type:     "类型",
		Content:  "内容",
		UserList: []string{"User1", "User2", "User3"},
	}
	fmt.Printf("%%v：%v\n", d)
	fmt.Printf("%%+v：%+v\n", d)
	fmt.Printf("%%#v：%#v\n", d)

}

/**
 * scan相关
 * @see
 * @param
 */
func UScan() {
	var (
		name    string
		age     int
		married bool
	)
	//fmt.Scan(&name, &age, &married)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	// 获取变量时，需要严格按照指定的格式来书写
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

}
