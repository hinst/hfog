package fog

func BoolToData(x bool) []byte {
	if x {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

func BoolFromData(data []byte) bool {
	return data != nil && len(data) > 0 && data[0] != 0
}
