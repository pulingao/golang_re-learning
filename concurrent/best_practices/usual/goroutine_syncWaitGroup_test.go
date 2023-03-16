package usual

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Goroutine 配合 sync.WaitGroup 使用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestT_wg(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "sync.WaitGroup 单独体验使用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T_wg()
		})
	}
}
