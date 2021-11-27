package serial

import "reflect"

// Range represents a range with a `From-To` pair
type Range struct {
	From, To int
}

// Ranges returns Range-slice
// via an interface{} object whose elements match `match` function
// continuously more than continouse-times
func Ranges(data interface{}, continuous int, match func(interface{}) bool) (result []Range) {
	// special cases
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		return
	}
	if continuous < 2 {
		return
	}

	matches := points(data, match)
	mLen := len(matches)
	if mLen == 0 {
		return nil
	}

	begin := matches[0]
	counter := 0
	var nextMatch int
	for matchIdx, dataIdx := range matches {
		notLast := matchIdx < mLen-1
		if notLast {
			nextMatch = matches[matchIdx+1]
			if nextMatch == dataIdx+1 {
				counter++
				continue
			}
		}

		if counter >= continuous-1 {
			// found
			result = append(result, Range{
				From: begin,
				To:   dataIdx,
			})
		}

		begin = nextMatch
		counter = 0
	}
	return result
}

// matchIndex returns int slice
// whose elements in target `data` slice matches `match` function
func points(data interface{}, match func(interface{}) bool) (ints []int) {
	value := reflect.ValueOf(data)
	for i := 0; i < value.Len(); i++ {
		element := value.Index(i)
		if match(element.Interface()) {
			ints = append(ints, i)
		}
	}
	return
}
