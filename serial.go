package serial

import "reflect"

// Range represents a range with a "From-To" pair
type Range struct {
	From, To int
}

// Ranges returns Range-slice
// whose elements match "match" function
// continuously more-than or equal-to continuous times
func Ranges(data interface{}, continuous int, match func(interface{}) bool) (ranges []Range) {
	// special cases
	if data == nil {
		return
	}
	if continuous < 2 {
		return
	}

	// only Slice and Array can pass through
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return
	}

	// indexes of matched elements
	indexes := matchIndexes(data, match)
	l := len(indexes)
	if l == 0 {
		return
	}

	begin := indexes[0]
	counter := 0
	var nextMatch int
	for matchIdx, dataIdx := range indexes {
		notLast := matchIdx < l-1
		if notLast {
			nextMatch = indexes[matchIdx+1]
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

// matchIndexes returns int slice
// whose elements in data match "match" function
func matchIndexes(data interface{}, match func(interface{}) bool) (indexes []int) {
	value := reflect.ValueOf(data)
	for i := 0; i < value.Len(); i++ {
		element := value.Index(i)
		if match(element.Interface()) {
			indexes = append(indexes, i)
		}
	}
	return
}
