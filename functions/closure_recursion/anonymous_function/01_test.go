package anonymous_function

import "testing"

func TestT_01(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "定义一个 接收匿名函数作为参数 的 函数，实现回调效果",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_01()
		})
	}
}
