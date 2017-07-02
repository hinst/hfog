package fog

type TBugHeaderWebStruct struct {
	Number int
	Title  string
	Rank   int
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

func (this TBugHeaderWebStruct) GetFromRankedMap(a map[int]TRankedTitle) (result []TBugHeaderWebStruct) {
	result = make([]TBugHeaderWebStruct, 0, len(a))
	for key, value := range a {
		result = append(result,
			TBugHeaderWebStruct{
				Number: key,
				Title:  value.Title,
				Rank:   value.Rank,
			})
	}
	return
}
