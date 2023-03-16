package usual

import "testing"

func Test_GR_NoBufChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "使用无缓冲的channel实现并发控制",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GR_NoBufChan()
		})
	}
}
