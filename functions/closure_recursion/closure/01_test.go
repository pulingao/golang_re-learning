package closure

import (
	"testing"
)

func TestT_00(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "00.概念",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_00()
		})
	}
}

func TestT_01(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "01.闭包的基础演示",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_01()
		})
	}
}

func TestT_02(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "02.闭包复制的是原对象指针（作为引用传递）",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_02()
		})
	}
}

func TestT_03(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "03.外部引用函数参数",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_03()
		})
	}
}

func TestT_04(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "04.返回两个闭包",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_04()
		})
	}
}
