package main

import (
	"fmt"
	"strings"
)

func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}

func reverString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str), true
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

type People struct {
	Name string
}

// 打印字符串的时候，会调用String()方法，所以此处会造成循环调用，出现问题
func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
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

func main() {

	isUni := isUniqueString("abcdefghijjklmn")
	fmt.Println(isUni)

	sa, _ := reverString("abcdef")
	fmt.Println(sa)

	// byte数组，打印的是底层的字符，中文也是会打印对应Unicode的字符
	// rune数组，打印的也是底层的字符，英文情况下和byte一致，中文或者其他语言需要用多个Unicode字符表现一个文字时，rune会按照格式（中文3个Unicode字符）来打印
	s := "abc中国人12"
	fmt.Println([]byte(s))
	for i, v := range []byte(s) {
		fmt.Printf("byte %v = %v，string：%v\n", i, v, string(v))
	}
	fmt.Println([]rune(s))
	for i, v := range []rune(s) {
		fmt.Printf("rune %v = %v，string：%v\n", i, v, string(v))
	}

	number := make(chan bool)
	letter := make(chan bool)
	done := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				if j >= 'Z' {
					done <- true
				} else {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					number <- true
				}
			}
		}
	}()

	number <- true

	for {
		select {
		case <-done:
			return
		}
	}
}
