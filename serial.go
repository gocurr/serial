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
	var nextM int
	for i, m := range matches {
		notLast := i < mLen-1
		if notLast {
			nextM = matches[i+1]
			if m+1 == nextM {
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

		begin = nextM
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
