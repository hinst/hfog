package fog

import (
	"strconv"
	"strings"
)

func StrToInt0(text string) (result int) {
	var x, parseResult = strconv.ParseInt(text, 10, 32)
	if parseResult == nil {
		result = int(x)
	}
	return
}

func IntToStr(x int) string {
	return strconv.Itoa(x)
}

func StringContainsAnyFromArray(text string, words []string) (result bool) {
	for _, word := range words {
		if strings.Contains(text, word) {
			result = true
			break
		}
	}
	return
}
