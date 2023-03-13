package basic_type

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringChange(t *testing.T) {
	StringChange()

	s := " 1A2B3C9Y4D5E6F"
	fmt.Println(strings.Trim(s, " []{}"))
}
