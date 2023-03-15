package istruct

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name           string `json:"name"`
	City           string `json:"city"`
	Age            int
	Height, Weight float32
	int            // 匿名字段，默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个
	string         // 匿名字段
}

/**
 * 创建某个结构体的构造函数，约定俗成的一些用法
 * @see
 * @param
 */
func NewPerson(name, city string, age int) *Person {
	return &Person{
		Name: name,
		City: city,
		Age:  age,
	}
}

/**
 * 针对于某个结构体，定义其 其他属性的__set方法
 * @see
 * @param
 */
func (p *Person) SetHeight(height float32) {
	p.Height = height
}

func (p *Person) SetWeight(weight float32) {
	p.Weight = weight
}

func (p *Person) SetInt(n int) {
	p.int = n
}

func (p *Person) SetString(s string) {
	p.string = s
}

func T_Person() {
	p := NewPerson("张三", "成都", 29)
	fmt.Println("实例化后（通过构造函数处理）：", p)

	p.SetHeight(177.2)
	p.SetWeight(85.2)
	p.SetInt(33)
	p.SetString("我是一个用来设置值的字符串")
	fmt.Println("使用Set方法后：", p)

	jsonData, err := json.Marshal(p)
	fmt.Printf("json的Byte数组：%v，\n字符串Json：%v，\n错误信息：%v\n", jsonData, string(jsonData), err)
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

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
}

func T_student() {
	ce := make(map[int]student)
	ce[1] = student{1, "xiaolizi", 22}
	ce[2] = student{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)

	fmt.Println()

	ace := make([]student, 0, 2)
	ace = []student{
		{id: 3, name: "王五", age: 22},
		{id: 4, name: "赵六", age: 38},
	}
	ace = append(ace, student{
		id: 1, name: "张三", age: 18,
	}, student{
		id: 2, name: "李四", age: 20,
	})
	fmt.Println("ace", ace)
	demo(ace)
	fmt.Println("ace", ace)

	T_switch()

}

func T_switch() {
	//写法三
	var k = 0
	switch k {
	case 0:
		println("fallthrough")
		fallthrough
		/*
		   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		   而如果switch没有表达式，它会匹配true。
		   Go里面switch默认相当于每个case最后带有break，
		   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		   但是可以使用fallthrough强制执行后面紧跟的case代码，注意是后面紧跟的，不是按照规则来匹配
		*/
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	case 1:
		fmt.Println("1")
	}
}
