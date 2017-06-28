package fog

import "encoding/xml"

type TBugData struct {
	XMLName xml.Name `xml:response`
	Cases   TBugCasesData
}

type TBugCasesData struct {
	XMLName xml.Name       `xml:"cases"`
	Cases   []TBugCaseData `xml:"case"`
}

type TBugCaseData struct {
	IxBug  string             `xml:"ixBug,attr"`
	STitle string             `xml:"sTitle,cdata"`
	Events TBugCaseEventsData `xml:"events"`
}

type TBugCaseEventsData struct {
	Events []TBugCaseEventData `xml:"event"`
}

type TBugCaseEventData struct {
	Dt             string `xml:"dt"`
	SVerb          string `xml:"sVerb,cdata"`
	EvtDescription string `xml:"evtDescription,cdata"`
	SPerson        string `xml:"sPerson,cdata"`
	S              string `xml:"s,cdata"`
	SHTML          string `xml:"sHtml,cdata"`
}
