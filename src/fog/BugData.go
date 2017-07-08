package fog

import (
	"encoding/xml"
)

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

	RGAttachments TRGAttachments `xml:"rgAttachments"`
}

type TRGAttachments struct {
	Attachments []TAttachment `xml:"attachment"`
}

type TAttachment struct {
	SFileName TCDATA `xml:"sFileName"`
	SURL      TCDATA `xml:"sURL"`
}

func (this *TBugCaseData) ToBugDataWeb() (result *TBugDataWeb) {
	result = &TBugDataWeb{
		Id:    this.IxBug,
		Title: this.STitle.Text,
	}
	for _, event := range this.Events.Events {
		var newEvent = TBugDataWebEvent{
			Moment:      event.Dt,
			Verb:        event.SVerb.Text,
			Description: event.EvtDescription.Text,
			Person:      event.SPerson.Text,
			Text:        event.S.Text,
			HTML:        event.SHTML.Text,
		}
		result.Events = append(result.Events, newEvent)
	}
	return
}
