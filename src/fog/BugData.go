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
	STitle TCDATA             `xml:"sTitle"`
	Events TBugCaseEventsData `xml:"events"`
}

type TBugCaseEventsData struct {
	Events []TBugCaseEventData `xml:"event"`
}

type TBugCaseEventData struct {
	Dt             string `xml:"dt"`
	SVerb          TCDATA `xml:"sVerb"`
	EvtDescription TCDATA `xml:"evtDescription"`
	SPerson        TCDATA `xml:"sPerson"`
	S              TCDATA `xml:"s"`
	SHTML          TCDATA `xml:"sHtml"`
}

func (this *TBugCaseEventData) ToDBStruct() TDBFlatStructArray {
	return TDBFlatStructArray{}.SaveStrings(
		map[string]string{
			"Dt":             this.Dt,
			"SVerb":          this.SVerb.Text,
			"EvtDescription": this.EvtDescription.Text,
			"SPerson":        this.SPerson.Text,
			"S":              this.S.Text,
			"SHTML":          this.SHTML.Text,
		})
}
