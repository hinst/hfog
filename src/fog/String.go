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

// Case insensitive
func CountStringContainsFromArray(text string, words []string) (result int) {
	text = strings.ToUpper(text)
	for _, word := range words {
		if strings.Contains(text, strings.ToUpper(word)) {
			result++
		}
	}
	return
}

func CheckStringHasSuffixes(text string, suffixes []string) (result bool) {
	for _, suffix := range suffixes {
		if strings.HasSuffix(text, suffix) {
			result = true
			break
		}
	}
	return
}
