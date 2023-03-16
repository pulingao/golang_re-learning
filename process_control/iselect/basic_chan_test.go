package iselect

import "testing"

func TestT_basic_chan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "第一个测试"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_basic_chan()
		})
	}
}

func TestT_basic_chan2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "第二个测试"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_basic_chan2()
		})
	}
}
