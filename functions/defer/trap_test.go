package _defer

import "testing"

func TestDeferTrap(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "01.defer和闭包产生的一个陷阱"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeferTrap()
		})
	}
}

func TestDeferNil(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "02.当defer定义的函数是nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeferNil()
		})
	}
}

func TestDeferCloseFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "03.defer作为关闭多个文件指针时，需要使用闭包+传参的方式"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeferCloseFile()
		})
	}
}
