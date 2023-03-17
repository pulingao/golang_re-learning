package recursion

import "testing"

func TestT_00(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "00.概念"},
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
		{name: "01.阶乘实现"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_01()
		})
	}
}
