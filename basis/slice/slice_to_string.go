package slice

import (
	"fmt"
	"strings"
)

type Info struct {
	infoName  string
	infoTitle string
}

/**
 * 将一个slice或者数组转化为字符串展示
 * @see		https://www.topgoer.com/go%E5%9F%BA%E7%A1%80/%E5%88%87%E7%89%87Slice.html，最后一行
 * @see		是否可以封装成为一个共有函数，倒是也没有太大的用处
 * @param
 */
func SliceToString1() {
	//slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 注意：《struct{ x string }》相当于 Info 的作用
	slice := []struct {
		name    string
		age     int
		info    Info
		content struct {
			x string
		}
	}{
		{name: "张三", age: 10, info: Info{infoName: "床前", infoTitle: "明月光"}, content: struct{ x string }{x: "再来一个结构体1"}},
		{name: "李四", age: 11, info: Info{infoName: "疑是", infoTitle: "地上霜"}, content: struct{ x string }{x: "再来一个结构体2"}},
		{name: "王五", age: 12, info: Info{infoName: "举头", infoTitle: "望明月"}, content: struct{ x string }{x: "再来一个结构体3"}},
		{name: "赵六", age: 13, info: Info{infoName: "低头", infoTitle: "思故乡"}, content: struct{ x string }{x: "再来一个结构体4"}},
	}
	tmp := fmt.Sprint(slice)
	tmp1 := strings.Trim(tmp, "[]{}")
	replaceString := []string{" ", "{", "}"}
	for _, v := range replaceString {
		tmp1 = strings.ReplaceAll(tmp1, v, "")
	}
	fmt.Printf("tmp，类型：%T，值：%v\n", tmp1, tmp1)

}

/**
 * 作为一个提示方法放在这里吧，没有什么通用性
 * @see
 * @param
 */
func StringSliceConvToInterface(args ...string) []interface{} {
	ss := make([]interface{}, len(args))
	for i, v := range args {
		ss[i] = v
	}
	return ss
}

/**
 * 封装的函数
 * @see
 * @param
 */
func SliceToString(s []interface{}) string {
	if s == nil {
		return ""
	}
	tmp := fmt.Sprint(s)
	tmp = strings.Trim(tmp, "[]{}")
	replaceString := []string{" ", "{", "}"}
	for _, v := range replaceString {
		tmp = strings.ReplaceAll(tmp, v, "")
	}
	return strings.Trim(tmp, " []{}")
}
