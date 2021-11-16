package serial

import (
	"fmt"
	"testing"
)

func TestRanges(t *testing.T) {
	data := []interface{}{-2, -1, 0, 1, 2, 3, 4, 5, 6, 3, 3, 2, 3, 3, 3, 3, 4}
	ranges := Ranges(data, 2, func(v interface{}) bool {
		return v.(int) > 2 && v.(int) < 5
	})
	fmt.Println(ranges)
}
