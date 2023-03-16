package main

import (
	"testing"
)

func Test_main2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试使用range来遍历channel"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main2()
		})
	}
}

func TestT_test(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "for获取值，取不到后跳出"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_test()
		})
	}
}
