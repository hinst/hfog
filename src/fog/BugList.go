package fog

import (
	"encoding/xml"
	"io/ioutil"
)

type TBugListCase struct {
	IxBug string `xml:"ixBug,attr"`
}

type TBugListCases struct {
	XMLName xml.Name       `xml:"cases"`
	Cases   []TBugListCase `xml:"case"`
}

type TBugList struct {
	XMLName xml.Name `xml:response`
	Cases   TBugListCases
}

func ReadBugsFromFile(filePath string) *TBugList {
	var bugs TBugList
	var data, readResult = ioutil.ReadFile(filePath)
	AssertResult(readResult)
	var decodeResult = xml.Unmarshal(data, &bugs)
	AssertResult(decodeResult)
	return &bugs
}
