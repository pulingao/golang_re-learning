package basic_type

import "fmt"

/**
 * 使用rune类型或byte类型，修改字符串的内容
 * @see		https://www.topgoer.com/go%E5%9F%BA%E7%A1%80/%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B.html
 * @see		测试：go test 使用参考：https://luckymrwang.github.io/2018/08/31/%E4%BD%BF%E7%94%A8Go-Test%E6%B5%8B%E8%AF%95%E5%8D%95%E4%B8%AA%E6%96%87%E4%BB%B6%E5%92%8C%E5%8D%95%E4%B8%AA%E6%96%B9%E6%B3%95/
 * @param
 */
func StringChange() {
	s1 := "hello"

	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)

	for i, v := range runeS2 {
		fmt.Printf("第%v个，是：%v（%c）\n", i, v, v)
	}

	runeS2[0] = '播'
	fmt.Println(string(runeS2))
}
