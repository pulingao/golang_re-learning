package usual

import "testing"

func TestGR_SyncCond(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "Goroutine 配合 sync.Cond 使用",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GR_SyncCond()
		})
	}
}
