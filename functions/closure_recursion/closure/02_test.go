package closure

import "testing"

func TestT_001(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "defer函数定义时，对外部变量引用的两种方式",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_001()
		})
	}
}

func TestT_002(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "[变种]defer函数定义时，对外部变量引用的两种方式，值的修改在defer函数定义之前",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_002()
		})
	}
}

func TestT_003(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "通过Adder()返回一个匿名函数，这个匿名函数和自由变量x组成闭包，只要匿名函数的实例closure没有消亡，那么x都是引用传递",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_003()
		})
	}
}
