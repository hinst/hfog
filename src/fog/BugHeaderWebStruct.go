package fog

type TBugHeaderWebStruct struct {
	Number int
	Title  string
}

func (this TBugHeaderWebStruct) GetFromMap(a map[int]string) (result []TBugHeaderWebStruct) {
	result = make([]TBugHeaderWebStruct, 0, len(a))
	for key, value := range a {
		result = append(result,
			TBugHeaderWebStruct{
				Number: key,
				Title:  value,
			})
	}
	return
}
