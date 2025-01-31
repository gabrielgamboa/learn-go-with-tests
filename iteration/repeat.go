package iteration

import "strings"

const REPEAT_COUNT = 5

func Repeat(charactere string, repeatTimes ...int) string {

	//Strings are immutable in Go, and use strings.Builder to create long streams makes performance better
	var response strings.Builder //Here we are declaring a string.Builder variable only
	var count int

	if len(repeatTimes) == 1 {
		count = repeatTimes[0]
	} else {
		count = REPEAT_COUNT
	}

	for i := 0; i < count; i++ {
		response.WriteString(charactere)
	}

	return response.String()
}
