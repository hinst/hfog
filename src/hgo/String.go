package hgo

import "strconv"

func BoolToStr(x bool) string {
	if x {
		return "true"
	} else {
		return "false"
	}
}

func Int64ToStr(x int64) string {
	return strconv.FormatInt(x, 10)
}
