package serial

import "reflect"

// Range represents a range with a `From-To` pair
type Range struct {
	From, To int
}

// Ranges returns Range-slice
// via an interface{} object whose elements match `match` function
// continuously more-than or equal-to continouse times
func Ranges(data interface{}, continuous int, match func(interface{}) bool) (ranges []Range) {
	// special cases
	if data == nil {
		return
	}
	if continuous < 2 {
		return
	}

	// only Slice and Array pass through
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return
	}

	// indexs in data matched `match` function
	indexs := matchIndexs(data, match)
	l := len(indexs)
	if l == 0 {
		return
	}

	begin := indexs[0]
	counter := 0
	var nextMatch int
	for matchIdx, dataIdx := range indexs {
		notLast := matchIdx < l-1
		if notLast {
			nextMatch = indexs[matchIdx+1]
			if nextMatch == dataIdx+1 {
				counter++
				continue
			}
		}

		if counter >= continuous-1 {
			// found
			ranges = append(ranges, Range{
				From: begin,
				To:   dataIdx,
			})
		}

		begin = nextMatch
		counter = 0
	}
	return
}

// matchIndexs returns int slice
// whose elements in target `data` slice matches `match` function
func matchIndexs(data interface{}, match func(interface{}) bool) (ints []int) {
	value := reflect.ValueOf(data)
	for i := 0; i < value.Len(); i++ {
		element := value.Index(i)
		if match(element.Interface()) {
			ints = append(ints, i)
		}
	}
	return
}
