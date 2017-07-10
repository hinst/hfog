package fog

import (
	"strconv"
	"strings"
)

func IntToStr(x int) string {
	return strconv.Itoa(x)
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
