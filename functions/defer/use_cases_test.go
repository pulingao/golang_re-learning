package _defer

import (
	"fmt"
	"github.com/pulingao/golang_re-learning/tools"
	"testing"
)

func TestCounterExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "01.使用Defer的一个反例",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CounterExample()

			tools.NewLine()
			fmt.Println("*********************************************************************************")
			FirstToSolveCE()
			SecondToSolveCE()
			ThirdToSolveCE()
		})
	}
}

func TestFirstToSolveCE(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "02.第一个解决办法"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FirstToSolveCE()
		})
	}
}

func TestSecondToSolveCE(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "03.第二个解决办法"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SecondToSolveCE()
		})
	}
}
func TestThirdToSolveCE(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "04.第三个解决办法"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ThirdToSolveCE()
		})
	}
}

func TestDeferError(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "05.延迟调用发生错误，这些调用依旧会被执行"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeferError()
		})
	}
}

func TestDeferPerformance(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "06.滥用 defer 可能会导致性能问题，尤其是在一个 \"大循环\" 里。"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeferPerformance()
		})
	}
}
