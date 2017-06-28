package fog

import (
	"strconv"
)

func StrToInt0(text string) (result int) {
	var x, parseResult = strconv.ParseInt(text, 10, 32)
	if parseResult == nil {
		result = int(x)
	}
	return
}
