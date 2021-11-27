package serial

import (
	"fmt"
	"testing"
)

type V struct {
	Val int
}

func TestRanges(t *testing.T) {
	data := []interface{}{nil, 1, V{Val: 5}, V{Val: 3}}
	ranges := Ranges(data, 2, func(i interface{}) bool {
		v, ok := i.(V)
		if !ok {
			return false
		}
		return v.Val > 2
	})
	fmt.Println(ranges)
}
