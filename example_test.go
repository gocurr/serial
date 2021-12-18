package serial_test

import (
	"fmt"
	"github.com/gocurr/serial"
)

type V struct {
	Val int
}

func ExampleRanges() {
	data := []interface{}{nil, 6, V{Val: 5}, V{Val: 3}}
	ranges := serial.Ranges(data, 2, func(i interface{}) bool {
		v, ok := i.(V)
		if !ok {
			return false
		}
		return v.Val > 2
	})
	fmt.Println(ranges)

	// Output: [{2 3}]
	//
}
