package serial

import (
	"fmt"
	"testing"
)

func TestRanges(t *testing.T) {
	continuous := 2

	data := []int{-2, -1, 0, 1, 2, 3, 4, 5, 5, 1, 5, 5, 3, 3, 3, 3, 4}
	ranges := Ranges(data, continuous, func(v interface{}) bool {
		return v.(int) > 4 && v.(int) < 6
	})
	fmt.Println(ranges)
}
