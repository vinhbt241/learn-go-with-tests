package iteration

import "strings"

func Repeat(input string, repeatTimes int) string {
	var repeated strings.Builder

	for range repeatTimes {
		repeated.WriteString(input)
	}

	return repeated.String()
}
