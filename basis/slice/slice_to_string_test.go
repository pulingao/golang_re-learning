package slice

import "testing"

func TestSliceToString(t *testing.T) {
	SliceToString1()

	ss := []string{"1", "2"}
	StringSliceConvToInterface(ss...)
}
