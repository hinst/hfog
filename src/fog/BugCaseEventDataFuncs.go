package fog

import "hgo"

func (this *TBugCaseEventData) LoadDBStruct(a TDBFlatStructArray) {
	var fields = a.ReadStrings()
	this.Dt = fields["Dt"]
	this.SVerb.Text = fields["SVerb"]
	this.EvtDescription.Text = fields["EvtDescription"]
	this.SPerson.Text = fields["SPerson"]
	this.S.Text = fields["S"]
	this.SHTML.Text = fields["SHTML"]
	this.LoadDBStructAttachments(fields)
}

func (this *TBugCaseEventData) LoadDBStructAttachments(fields map[string]string) {
	var n = hgo.StrToInt0(fields["attachments"])
	if n > 0 {
		this.RGAttachments.Attachments = make([]TAttachment, n)
		for i := 0; i < n; i++ {
			this.RGAttachments.Attachments[i].SFileName.Text = fields["attachment/"+IntToStr(i)+"/sFileName"]
			this.RGAttachments.Attachments[i].SURL.Text = fields["attachment/"+IntToStr(i)+"sURL"]
		}
	}
}

func (this *TBugCaseEventData) ToDBStruct() TDBFlatStructArray {
	var fields = map[string]string{
		"Dt":             this.Dt,
		"SVerb":          this.SVerb.Text,
		"EvtDescription": this.EvtDescription.Text,
		"SPerson":        this.SPerson.Text,
		"S":              this.S.Text,
		"SHTML":          this.SHTML.Text,
	}
	this.SaveAttachmentsToFields(fields)
	return TDBFlatStructArray{}.SaveStrings(fields)
}

func (this *TBugCaseEventData) SaveAttachmentsToFields(fields map[string]string) {
	var n = len(this.RGAttachments.Attachments)
	fields["attachments"] = IntToStr(n)
	for i := 0; i < n; i++ {
		fields["attachment/"+IntToStr(i)+"/sFileName"] = this.RGAttachments.Attachments[i].SFileName.Text
		fields["attachment/"+IntToStr(i)+"sURL"] = this.RGAttachments.Attachments[i].SURL.Text
	}
}
