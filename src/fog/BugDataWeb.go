package fog

type TBugDataWeb struct {
	Id     string
	Title  string
	Events []TBugDataWebEvent
}

type TBugDataWebAttachment struct {
	FileName string
	KeyURL   string
}

type TBugDataWebEvent struct {
	Moment      string
	Verb        string
	Description string
	Person      string
	Text        string
	HTML        string
	Attachments []TBugDataWebAttachment
}
