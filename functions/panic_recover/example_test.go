package panic_recover

import "testing"

func TestT_0(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "00.概念和注意点"},
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
		{name: "01.延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_1()
		})
	}
}

func TestT_2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "02.捕获函数 recover 只有在 defer调用的函数内 直接调用 才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_2()
		})
	}
}

func TestT_3(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "03.如果在一段函数中需要完整的执行下去，避免panic的中断，可以在匿名函数中使用recover将可能出现的异常捕获并处理"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_3()
		})
	}
}

func TestT_4(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "04.实现简单的 Try Catch 的效果"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_4()
		})
	}
}
