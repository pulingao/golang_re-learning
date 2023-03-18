package _interface

import "testing"

func TestT_0(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "01.概念及表述"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_0()
		})
	}
}

func TestT_1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "02.值接收者和指针接收者实现接口的区别"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_1()
		})
	}
}
