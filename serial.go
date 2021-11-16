package serial

type Range struct {
	From, To int
}

func Ranges(data []interface{}, continouse int, match func(interface{}) bool) (result []Range) {
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

		if counter >= continouse-1 {
			// found
			result = append(result, Range{
				From: begin,
				To:   begin + counter,
			})
		}

		begin = nextMatch
		counter = 0
	}
	return result
}

func points(data []interface{}, match func(interface{}) bool) (result []int) {
	for i, d := range data {
		if match(d) {
			result = append(result, i)
		}
	}
	return result
}
