package fog

import "encoding/xml"

type TBugListCase struct {
	IxBug string `xml:"ixBug,attr"`
}

type TBugListCaseCtnr struct {
}

type TBugListCases struct {
	XMLName xml.Name       `xml:"cases"`
	Cases   []TBugListCase `xml:"case"`
}

type TBugList struct {
	XMLName xml.Name `xml:response`
	Cases   TBugListCases
}
